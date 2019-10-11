// Copyright (c) 2017-2019 Snowflake Computing Inc. All right reserved.

package gosnowflake

import (
	"database/sql/driver"
	"encoding/hex"
	"fmt"
	"reflect"
	"strconv"
	"strings"
	"time"
)

// goTypeToSnowflake translates Go data type to Snowflake data type.
func goTypeToSnowflake(v driver.Value, tsmode string) string {
	switch v := v.(type) {
	case int64:
		return "FIXED"
	case float64:
		return "REAL"
	case bool:
		return "BOOLEAN"
	case string:
		return "TEXT"
	case []byte:
		if tsmode == "BINARY" {
			return "BINARY" // may be redundant but ensures BINARY type
		}
		if v == nil || len(v) != 1 {
			return "TEXT" // invalid byte array. won't take as BINARY
		}
		_, err := dataTypeMode(v)
		if err != nil {
			return "TEXT" // not supported dataType
		}
		return "CHANGE_TYPE"
	case time.Time:
		return tsmode
	}
	return "TEXT"
}

// snowflakeTypeToGo translates Snowflake data type to Go data type.
func snowflakeTypeToGo(dbtype string, scale int64) reflect.Type {
	switch dbtype {
	case "fixed":
		if scale == 0 {
			return reflect.TypeOf(int64(0))
		}
		return reflect.TypeOf(float64(0))
	case "real":
		return reflect.TypeOf(float64(0))
	case "text", "variant", "object", "array":
		return reflect.TypeOf("")
	case "date", "time", "timestamp_ltz", "timestamp_ntz", "timestamp_tz":
		return reflect.TypeOf(time.Now())
	case "binary":
		return reflect.TypeOf([]byte{})
	case "boolean":
		return reflect.TypeOf(true)
	}
	glog.V(1).Infof("unsupported dbtype is specified. %v", dbtype)
	glog.Flush()
	return reflect.TypeOf("")
}

// valueToString converts arbitrary golang type to a string. This is mainly used in binding data with placeholders
// in queries.
func valueToString(v driver.Value, tsmode string) (*string, error) {
	glog.V(2).Infof("TYPE: %v, %v", reflect.TypeOf(v), reflect.ValueOf(v))
	if v == nil {
		return nil, nil
	}
	v1 := reflect.ValueOf(v)
	switch v1.Kind() {
	case reflect.Bool:
		s := strconv.FormatBool(v1.Bool())
		return &s, nil
	case reflect.Int64:
		s := strconv.FormatInt(v1.Int(), 10)
		return &s, nil
	case reflect.Float64:
		s := strconv.FormatFloat(v1.Float(), 'g', -1, 32)
		return &s, nil
	case reflect.String:
		s := v1.String()
		return &s, nil
	case reflect.Slice, reflect.Map:
		if v1.IsNil() {
			return nil, nil
		}
		if bd, ok := v.([]byte); ok {
			if tsmode == "BINARY" {
				s := hex.EncodeToString(bd)
				return &s, nil
			}
		}
		// TODO: is this good enough?
		s := v1.String()
		return &s, nil
	case reflect.Struct:
		if tm, ok := v.(time.Time); ok {
			switch tsmode {
			case "DATE":
				_, offset := tm.Zone()
				tm = tm.Add(time.Second * time.Duration(offset))
				s := fmt.Sprintf("%d", tm.Unix()*1000)
				return &s, nil
			case "TIME":
				s := fmt.Sprintf("%d",
					(tm.Hour()*3600+tm.Minute()*60+tm.Second())*1e9+tm.Nanosecond())
				return &s, nil
			case "TIMESTAMP_NTZ", "TIMESTAMP_LTZ":
				s := fmt.Sprintf("%d", tm.UnixNano())
				return &s, nil
			case "TIMESTAMP_TZ":
				_, offset := tm.Zone()
				s := fmt.Sprintf("%v %v", tm.UnixNano(), offset/60+1440)
				return &s, nil
			}
		}
	}
	return nil, fmt.Errorf("unsupported type: %v", v1.Kind())
}

// extractTimestamp extracts the internal timestamp data to epoch time in seconds and milliseconds
func extractTimestamp(srcValue *string) (sec int64, nsec int64, err error) {
	glog.V(2).Infof("SRC: %v", srcValue)
	var i int
	for i = 0; i < len(*srcValue); i++ {
		if (*srcValue)[i] == '.' {
			sec, err = strconv.ParseInt((*srcValue)[0:i], 10, 64)
			if err != nil {
				return 0, 0, err
			}
			break
		}
	}
	if i == len(*srcValue) {
		// no fraction
		sec, err = strconv.ParseInt(*srcValue, 10, 64)
		if err != nil {
			return 0, 0, err
		}
		nsec = 0
	} else {
		s := (*srcValue)[i+1:]
		nsec, err = strconv.ParseInt(s+strings.Repeat("0", 9-len(s)), 10, 64)
		if err != nil {
			return 0, 0, err
		}
	}
	glog.V(2).Infof("sec: %v, nsec: %v", sec, nsec)
	return sec, nsec, nil
}

// stringToValue converts a pointer of string data to an arbitrary golang variable. This is mainly used in fetching
// data.
func stringToValue(dest *driver.Value, srcColumnMeta execResponseRowType, srcValue *string) error {
	if srcValue == nil {
		glog.V(3).Infof("snowflake data type: %v, raw value: nil", srcColumnMeta.Type)
		*dest = nil
		return nil
	}
	glog.V(3).Infof("snowflake data type: %v, raw value: %v", srcColumnMeta.Type, *srcValue)
	switch srcColumnMeta.Type {
	case "text", "fixed", "real", "variant", "object":
		*dest = *srcValue
		return nil
	case "date":
		v, err := strconv.ParseInt(*srcValue, 10, 64)
		if err != nil {
			return err
		}
		*dest = time.Unix(v*86400, 0).UTC()
		return nil
	case "time":
		sec, nsec, err := extractTimestamp(srcValue)
		if err != nil {
			return err
		}
		t0 := time.Time{}
		*dest = t0.Add(time.Duration(sec*1e9 + nsec))
		return nil
	case "timestamp_ntz":
		sec, nsec, err := extractTimestamp(srcValue)
		if err != nil {
			return err
		}
		*dest = time.Unix(sec, nsec).UTC()
		return nil
	case "timestamp_ltz":
		sec, nsec, err := extractTimestamp(srcValue)
		if err != nil {
			return err
		}
		*dest = time.Unix(sec, nsec)
		return nil
	case "timestamp_tz":
		glog.V(2).Infof("tz: %v", *srcValue)

		tm := strings.Split(*srcValue, " ")
		if len(tm) != 2 {
			return &SnowflakeError{
				Number:   ErrInvalidTimestampTz,
				SQLState: SQLStateInvalidDataTimeFormat,
				Message:  fmt.Sprintf("invalid TIMESTAMP_TZ data. The value doesn't consist of two numeric values separated by a space: %v", *srcValue),
			}
		}
		sec, nsec, err := extractTimestamp(&tm[0])
		if err != nil {
			return err
		}
		offset, err := strconv.ParseInt(tm[1], 10, 64)
		if err != nil {
			return &SnowflakeError{
				Number:   ErrInvalidTimestampTz,
				SQLState: SQLStateInvalidDataTimeFormat,
				Message:  fmt.Sprintf("invalid TIMESTAMP_TZ data. The offset value is not integer: %v", tm[1]),
			}
		}
		loc := Location(int(offset) - 1440)
		tt := time.Unix(sec, nsec)
		*dest = tt.In(loc)
		return nil
	case "binary":
		b, err := hex.DecodeString(*srcValue)
		if err != nil {
			return &SnowflakeError{
				Number:   ErrInvalidBinaryHexForm,
				SQLState: SQLStateNumericValueOutOfRange,
				Message:  err.Error(),
			}
		}
		*dest = b
		return nil
	}
	*dest = *srcValue
	return nil
}
