package confeito

import (
	"reflect"
	"regexp"
	"testing"
)

// Assert is an assertion wrapper.
type Assert struct {
	tb       testing.TB
	expected []interface{}
}

// NewAssert returns an new assert with the testing context and the expected values.
func NewAssert(tb testing.TB, expected ...interface{}) *Assert {
	return &Assert{
		tb:       tb,
		expected: expected,
	}
}

// Equal checks that the given actual values equals the expected values.
func (assert *Assert) Equal(actual ...interface{}) {
	assert.tb.Helper()
	if len(assert.expected) != len(actual) {
		assert.tb.Fatalf("expected %d values, but got %d values", len(assert.expected), len(actual))
	} else {
		for i, expected := range assert.expected {
			if !reflect.DeepEqual(expected, actual[i]) {
				assert.tb.Errorf("at #%d value, expected %#v (%T), but got %#v (%T)", i, expected, expected, actual[i], actual[i])
			}
		}
	}
}

// EqualWithoutError checks that the given actual values equals the expected values without any error.
func (assert *Assert) EqualWithoutError(actual_err ...interface{}) {
	assert.tb.Helper()
	if len(actual_err) < 2 {
		assert.tb.Fatalf("actual_err must be at least two: (actual..., err)")
	}
	err := actual_err[len(actual_err)-1]
	if err != nil {
		assert.tb.Fatalf("unexpected error: %s", err)
	}
	assert.Equal(actual_err[0 : len(actual_err)-1]...)
}

// ExpectError checks that the error is returned expectedly.
func (assert *Assert) ExpectError(_err ...interface{}) {
	assert.tb.Helper()
	if len(_err) < 1 {
		assert.tb.Fatalf("_err must be at least one: (_..., err)")
	}
	err, ok := _err[len(_err)-1].(error)
	if !ok && _err[len(_err)-1] != nil {
		assert.tb.Fatalf("the last element of _err must be error, but got %T", err)
	}
	if err == nil {
		assert.tb.Errorf("expected an error, but got no error")
		return
	}
	if len(assert.expected) > 0 {
		if len(assert.expected) != 1 {
			assert.tb.Errorf("the number of error pattern must be at most one")
		}
		pattern, ok := assert.expected[0].(string)
		if !ok {
			assert.tb.Errorf("error pattern must be nil or string")
		}
		if matched, e := regexp.MatchString(pattern, err.Error()); e != nil {
			assert.tb.Errorf("malformed expected error pattern: %s", e)
		} else if !matched {
			assert.tb.Errorf("expected error pattern %q, but got error %q", pattern, err)
		}
	}
}

// SuccessNew check that New function successes without any error.
func (assert *Assert) SucceedNew(o interface{}, err error) interface{} {
	assert.tb.Helper()
	if err != nil {
		assert.tb.Errorf("unexpected error in New function: %s", err)
	}
	return o
}

// SuccessWithoutError check that the function successes without any error.
func (assert *Assert) SucceedWithoutError(err error) {
	assert.tb.Helper()
	if err != nil {
		assert.tb.Errorf("unexpected error: %s", err)
	}
}
