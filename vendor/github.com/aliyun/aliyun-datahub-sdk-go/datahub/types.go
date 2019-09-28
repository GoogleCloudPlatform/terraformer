package datahub

import (
	"encoding/json"
	"errors"
	"fmt"
	"strconv"
)

type DataType interface {
	fmt.Stringer
}

// Bigint
type Bigint int64

func (bi Bigint) String() string {
	return strconv.FormatInt(int64(bi), 10)
}

// String
type String string

func (str String) String() string {
	return string(str)
}

// Boolean
type Boolean bool

func (bl Boolean) String() string {
	return strconv.FormatBool(bool(bl))
}

// Double
type Double float64

func (d Double) String() string {
	return strconv.FormatFloat(float64(d), 'f', -1, 64)
}

// Timestamp
type Timestamp uint64

func (t Timestamp) String() string {
	return strconv.FormatUint(uint64(t), 10)
}

// FieldType
type FieldType string

func (ft FieldType) String() string {
	return string(ft)
}

const (
	// BIGINT 8-bit long signed integer, not include (-9223372036854775808)
	// -9223372036854775807 ~ 9223372036854775807
	BIGINT FieldType = "BIGINT"

	// only support utf-8
	// 1Mb max size
	STRING FieldType = "STRING"

	// BOOLEAN
	// True/False，true/false, 0/1
	BOOLEAN FieldType = "BOOLEAN"

	// DOUBLE 8-bit double
	// -1.0 * 10^308 ~ 1.0 * 10^308
	DOUBLE FieldType = "DOUBLE"

	// TIMESTAMP
	// unit: us
	TIMESTAMP FieldType = "TIMESTAMP"
)

// ValidateFieldType validate field type
func ValidateFieldType(ft FieldType) bool {
	switch ft {
	case BIGINT, STRING, BOOLEAN, DOUBLE, TIMESTAMP:
		return true
	default:
		return false
	}
}

// ValidateFieldValue validate field value
func ValidateFieldValue(ft FieldType, val interface{}) (DataType, error) {
	switch ft {
	case BIGINT:
		var realval Bigint
		switch v := val.(type) {
		case Bigint:
			realval = v
		case int:
			realval = Bigint(v)
		case int8:
			realval = Bigint(v)
		case int16:
			realval = Bigint(v)
		case int32:
			realval = Bigint(v)
		case int64:
			realval = Bigint(v)
		case uint:
			realval = Bigint(v)
		case uint8:
			realval = Bigint(v)
		case uint16:
			realval = Bigint(v)
		case uint32:
			realval = Bigint(v)
		case uint64:
			if v > 9223372036854775807 {
				return nil, errors.New("BIGINT type field must be in [-9223372036854775807,9223372036854775807]")
			}
			realval = Bigint(v)
		case json.Number:
			nval, err := v.Int64()
			if err != nil {
				return nil, err
			}
			realval = Bigint(nval)
		default:
			return nil, errors.New(fmt.Sprintf("value type[%T] not match field type[BIGINT]", val))
		}
		if int64(realval) < -9223372036854775807 || int64(realval) > 9223372036854775807 {
			return nil, errors.New("BIGINT type field must be in [-9223372036854775807,9223372036854775807]")
		}
		return realval, nil
	case STRING:
		var realval String
		switch v := val.(type) {
		case String:
			realval = v
		case string:
			realval = String(v)
		default:
			return nil, errors.New(fmt.Sprintf("value type[%T] not match field type[STRING]", val))
		}
		if len(string(realval)) > 1*1024*1024 {
			return nil, errors.New("STRING type value length must less than 1*1024*1024")
		}
		return realval, nil
	case BOOLEAN:
		switch v := val.(type) {
		case Boolean:
			return v, nil
		case bool:
			return Boolean(v), nil
		default:
			return nil, errors.New(fmt.Sprintf("value type[%T] not match field type[BOOLEAN]", val))
		}
	case DOUBLE:
		switch v := val.(type) {
		case Double:
			return v, nil
		case float64:
			return Double(v), nil
		case json.Number:
			nval, err := v.Float64()
			if err != nil {
				return nil, err
			}
			return Double(nval), nil
		default:
			return nil, errors.New(fmt.Sprintf("value type[%T] not match field type[DOUBLE]", val))
		}
	case TIMESTAMP:
		var realval Timestamp
		switch v := val.(type) {
		case Timestamp:
			realval = v
		case uint:
			realval = Timestamp(v)
		case uint8:
			realval = Timestamp(v)
		case uint16:
			realval = Timestamp(v)
		case uint32:
			realval = Timestamp(v)
		case uint64:
			realval = Timestamp(v)
		case int:
			if v < 0 {
				return nil, errors.New("TIMESTAMP type field must be in positive")
			}
			realval = Timestamp(v)
		case int8:
			if v < 0 {
				return nil, errors.New("TIMESTAMP type field must be in positive")
			}
			realval = Timestamp(v)
		case int16:
			if v < 0 {
				return nil, errors.New("TIMESTAMP type field must be in positive")
			}
			realval = Timestamp(v)
		case int32:
			if v < 0 {
				return nil, errors.New("TIMESTAMP type field must be in positive")
			}
			realval = Timestamp(v)
		case int64:
			if v < 0 {
				return nil, errors.New("TIMESTAMP type field must be in positive")
			}
			realval = Timestamp(v)
		case json.Number:
			nval, err := v.Int64()
			if err != nil {
				return nil, err
			}
			if nval < 0 {
				return nil, errors.New("TIMESTAMP type field must be in positive")
			}
			realval = Timestamp(nval)
		default:
			return nil, errors.New(fmt.Sprintf("value type[%T] not match field type[TIMESTAMP]", val))
		}
		return realval, nil
	default:
		return nil, errors.New(fmt.Sprintf("field type[%T] is not illegal", ft))
	}
}

// CastValueFromString cast value from string
func CastValueFromString(str string, ft FieldType) (DataType, error) {
	switch ft {
	case BIGINT:
		v, err := strconv.ParseInt(str, 10, 64)
		if err == nil {
			return Bigint(v), nil
		}
		return nil, err
	case STRING:
		return String(str), nil
	case BOOLEAN:
		v, err := strconv.ParseBool(str)
		if err == nil {
			return Boolean(v), nil
		}
		return nil, err
	case DOUBLE:
		v, err := strconv.ParseFloat(str, 64)
		if err == nil {
			return Double(v), nil
		}
		return nil, err
	case TIMESTAMP:
		v, err := strconv.ParseUint(str, 10, 64)
		if err == nil {
			return Timestamp(v), nil
		}
		return nil, err
	default:
		return nil, errors.New(fmt.Sprintf("not support field type %p", string(ft)))
	}
}

// RecordType
type RecordType string

func (rt RecordType) String() string {
	return string(rt)
}

const (
	// BLOB record
	BLOB RecordType = "BLOB"

	// TUPLE record
	TUPLE RecordType = "TUPLE"
)

// ValidateRecordType validate record type
func ValidateRecordType(rt RecordType) bool {
	switch rt {
	case BLOB, TUPLE:
		return true
	default:
		return false
	}
}

// ShardState
type ShardState string

func (state ShardState) String() string {
	return string(state)
}

const (
	// OPENING shard is creating or fail over, not available
	OPENING ShardState = "OPENING"

	// ACTIVE is available
	ACTIVE ShardState = "ACTIVE"

	// CLOSED read-only
	CLOSED ShardState = "CLOSED"

	// CLOSING shard is closing, not available
	CLOSING ShardState = "CLOSING"
)

// CursorType
type CursorType string

func (ct CursorType) String() string {
	return string(ct)
}

const (
	// OLDEST
	OLDEST CursorType = "OLDEST"

	// LATEST
	LATEST CursorType = "LATEST"

	// SYSTEM_TIME point to first record after system_time
	SYSTEM_TIME CursorType = "SYSTEM_TIME"
)

// ValidateCursorType validate field type
func ValidateCursorType(ct CursorType) bool {
	switch ct {
	case OLDEST, LATEST, SYSTEM_TIME:
		return true
	default:
		return false
	}
}

// SubscriptionType
type SubscriptionType int

const (
	// SUBTYPE_USER
	SUBTYPE_USER SubscriptionType = iota

	// SUBTYPE_SYSTEM
	SUBTYPE_SYSTEM

	// SUBTYPE_TT
	SUBTYPE_TT
)

func (subType SubscriptionType) Value() int {
	return int(subType)
}

// SubscriptionState
type SubscriptionState int

const (
	// SUB_OFFLINE
	SUB_OFFLINE SubscriptionState = iota

	// SUB_ONLINE
	SUB_ONLINE
)

func (subState SubscriptionState) Value() int {
	return int(subState)
}
