package fastly

import (
	"fmt"
	"net/http"
	"reflect"
	"time"

	"github.com/mitchellh/mapstructure"
)

// mapToHTTPHeaderHookFunc returns a function that converts maps into an
// http.Header value.
func mapToHTTPHeaderHookFunc() mapstructure.DecodeHookFunc {
	return func(
		f reflect.Type,
		t reflect.Type,
		data interface{}) (interface{}, error) {
		if f.Kind() != reflect.Map {
			return data, nil
		}
		if t != reflect.TypeOf(new(http.Header)) {
			return data, nil
		}

		typed, ok := data.(map[string]interface{})
		if !ok {
			return nil, fmt.Errorf("cannot convert %T to http.Header", data)
		}

		n := map[string][]string{}
		for k, v := range typed {
			switch v.(type) {
			case string:
				n[k] = []string{v.(string)}
			case []string:
				n[k] = v.([]string)
			case int, int8, int16, int32, int64:
				n[k] = []string{fmt.Sprintf("%d", v.(int))}
			case float32, float64:
				n[k] = []string{fmt.Sprintf("%f", v.(float64))}
			default:
				return nil, fmt.Errorf("cannot convert %T to http.Header", v)
			}
		}

		return n, nil
	}
}

// stringToTimeHookFunc returns a function that converts strings to a time.Time
// value.
func stringToTimeHookFunc() mapstructure.DecodeHookFunc {
	return func(
		f reflect.Type,
		t reflect.Type,
		data interface{}) (interface{}, error) {
		if f.Kind() != reflect.String {
			return data, nil
		}
		if t != reflect.TypeOf(time.Now()) {
			return data, nil
		}

		// Convert it by parsing
		v, err := time.Parse(time.RFC3339, data.(string))
		if err != nil {
			// DictionaryInfo#get uses it's own special time format for now.
			return time.Parse("2006-01-02 15:04:05", data.(string))
		}
		return v, err
	}
}
