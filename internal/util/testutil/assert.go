package testutil

import (
	"reflect"
	"runtime"
	"testing"
)

// Collection of simple functions to make testing values a bit more succinct

// AssertNoErr will fail immediately if the err is not null
// usefult for halting testing when an error exists
func AssertNoErr(t *testing.T, err error, message string) {
	if err != nil {
		_, file, line, _ := runtime.Caller(1)
		t.Logf("%s:%d %s\n", file, line, message)
		t.Log(err.Error())
		t.FailNow()
	}
}

// AssertNotNil will fail immediately if the preovided value is nil
// this will halt test execution to prevent any further operations
func AssertNotNil(t *testing.T, v interface{}, message string) {
	if v == nil || (reflect.ValueOf(v).Kind() == reflect.Ptr && reflect.ValueOf(v).IsNil()) {
		_, file, line, _ := runtime.Caller(1)
		t.Logf("%s:%d %s\n", file, line, message)
		t.Log("value must not be nil")
		t.FailNow()
	}
}

// CheckEq will check if the given value is equal to the given expectation
func CheckEq(t *testing.T, exp, got interface{}, message string) {
	if (exp == nil && got != nil) || !reflect.DeepEqual(exp, got) {
		_, file, line, _ := runtime.Caller(1)
		t.Logf("%s:%d %s\n", file, line, message)
		t.Logf("- %+v\n", exp)
		t.Logf("+ %+v\n", got)
		t.Fail()
	}
}

// Check will check if the given value is true
func Check(t *testing.T, got bool, message string) {
	if !got {
		_, file, line, _ := runtime.Caller(1)
		t.Logf("%s:%d %s\n", file, line, message)
		t.Fail()
	}
}
