package assert

import (
	"fmt"
	"path"
	"reflect"
	"regexp"
	"runtime"
	"testing"
)

// IsEqual returns whether val1 is equal to val2 taking into account Pointers, Interfaces and their underlying types
func IsEqual(val1, val2 interface{}) bool {
	v1 := reflect.ValueOf(val1)
	v2 := reflect.ValueOf(val2)

	if v1.Kind() == reflect.Ptr {
		v1 = v1.Elem()
	}

	if v2.Kind() == reflect.Ptr {
		v2 = v2.Elem()
	}

	if !v1.IsValid() && !v2.IsValid() {
		return true
	}

	switch v1.Kind() {
	case reflect.Chan, reflect.Func, reflect.Interface, reflect.Map, reflect.Ptr, reflect.Slice:
		if v1.IsNil() {
			v1 = reflect.ValueOf(nil)
		}
	}

	switch v2.Kind() {
	case reflect.Chan, reflect.Func, reflect.Interface, reflect.Map, reflect.Ptr, reflect.Slice:
		if v2.IsNil() {
			v2 = reflect.ValueOf(nil)
		}
	}

	v1Underlying := reflect.Zero(reflect.TypeOf(v1)).Interface()
	v2Underlying := reflect.Zero(reflect.TypeOf(v2)).Interface()

	if v1 == v1Underlying {
		if v2 == v2Underlying {
			goto CASE4
		} else {
			goto CASE3
		}
	} else {
		if v2 == v2Underlying {
			goto CASE2
		} else {
			goto CASE1
		}
	}

CASE1:
	return reflect.DeepEqual(v1.Interface(), v2.Interface())
CASE2:
	return reflect.DeepEqual(v1.Interface(), v2)
CASE3:
	return reflect.DeepEqual(v1, v2.Interface())
CASE4:
	return reflect.DeepEqual(v1, v2)
}

// NotMatchRegex validates that value matches the regex, either string or *regex
// and throws an error with line number
func NotMatchRegex(t *testing.T, value string, regex interface{}) {
	NotMatchRegexSkip(t, 2, value, regex)
}

// NotMatchRegexSkip validates that value matches the regex, either string or *regex
// and throws an error with line number
// but the skip variable tells NotMatchRegexSkip how far back on the stack to report the error.
// This is a building block to creating your own more complex validation functions.
func NotMatchRegexSkip(t *testing.T, skip int, value string, regex interface{}) {

	if r, ok, err := regexMatches(regex, value); ok || err != nil {
		_, file, line, _ := runtime.Caller(skip)

		if err != nil {
			fmt.Printf("%s:%d %v error compiling regex %v\n", path.Base(file), line, value, r.String())
		} else {
			fmt.Printf("%s:%d %v matches regex %v\n", path.Base(file), line, value, r.String())
		}

		t.FailNow()
	}
}

// MatchRegex validates that value matches the regex, either string or *regex
// and throws an error with line number
func MatchRegex(t *testing.T, value string, regex interface{}) {
	MatchRegexSkip(t, 2, value, regex)
}

// MatchRegexSkip validates that value matches the regex, either string or *regex
// and throws an error with line number
// but the skip variable tells MatchRegexSkip how far back on the stack to report the error.
// This is a building block to creating your own more complex validation functions.
func MatchRegexSkip(t *testing.T, skip int, value string, regex interface{}) {

	if r, ok, err := regexMatches(regex, value); !ok {
		_, file, line, _ := runtime.Caller(skip)

		if err != nil {
			fmt.Printf("%s:%d %v error compiling regex %v\n", path.Base(file), line, value, r.String())
		} else {
			fmt.Printf("%s:%d %v does not match regex %v\n", path.Base(file), line, value, r.String())
		}

		t.FailNow()
	}
}

func regexMatches(regex interface{}, value string) (*regexp.Regexp, bool, error) {

	var err error

	r, ok := regex.(*regexp.Regexp)

	// must be a string
	if !ok {
		if r, err = regexp.Compile(regex.(string)); err != nil {
			return r, false, err
		}
	}

	return r, r.MatchString(value), err
}

// Equal validates that val1 is equal to val2 and throws an error with line number
func Equal(t *testing.T, val1, val2 interface{}) {
	EqualSkip(t, 2, val1, val2)
}

// EqualSkip validates that val1 is equal to val2 and throws an error with line number
// but the skip variable tells EqualSkip how far back on the stack to report the error.
// This is a building block to creating your own more complex validation functions.
func EqualSkip(t *testing.T, skip int, val1, val2 interface{}) {

	if !IsEqual(val1, val2) {
		_, file, line, _ := runtime.Caller(skip)
		fmt.Printf("%s:%d %v does not equal %v\n", path.Base(file), line, val1, val2)
		t.FailNow()
	}
}

// NotEqual validates that val1 is not equal val2 and throws an error with line number
func NotEqual(t *testing.T, val1, val2 interface{}) {
	NotEqualSkip(t, 2, val1, val2)
}

// NotEqualSkip validates that val1 is not equal to val2 and throws an error with line number
// but the skip variable tells NotEqualSkip how far back on the stack to report the error.
// This is a building block to creating your own more complex validation functions.
func NotEqualSkip(t *testing.T, skip int, val1, val2 interface{}) {

	if IsEqual(val1, val2) {
		_, file, line, _ := runtime.Caller(skip)
		fmt.Printf("%s:%d %v should not be equal %v\n", path.Base(file), line, val1, val2)
		t.FailNow()
	}
}

// PanicMatches validates that the panic output of running fn matches the supplied string
func PanicMatches(t *testing.T, fn func(), matches string) {
	PanicMatchesSkip(t, 2, fn, matches)
}

// PanicMatchesSkip validates that the panic output of running fn matches the supplied string
// but the skip variable tells PanicMatchesSkip how far back on the stack to report the error.
// This is a building block to creating your own more complex validation functions.
func PanicMatchesSkip(t *testing.T, skip int, fn func(), matches string) {

	_, file, line, _ := runtime.Caller(skip)

	defer func() {
		if r := recover(); r != nil {
			err := fmt.Sprintf("%s", r)

			if err != matches {
				fmt.Printf("%s:%d Panic...  expected [%s] received [%s]", path.Base(file), line, matches, err)
				t.FailNow()
			}
		} else {
			fmt.Printf("%s:%d Panic Expected, none found...  expected [%s]", path.Base(file), line, matches)
			t.FailNow()
		}
	}()

	fn()
}
