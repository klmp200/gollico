package gollico

import (
	"reflect"
	"testing"
)

func TestGetIIIFDoc(t *testing.T) {
	for _, test := range ManifestStructTest {
		actual, err := GetIIIFDoc(test.ark)
		if test.expectErr {
			// check if err is of error type
			var _ error = err

			// if we expect an error and there isn't one
			if err == nil {
				t.Errorf("expected an error for ark %s, but error is nil", test.ark)
			}
			t.Logf("PASS: got %v", err)
		} else {
			if reflect.DeepEqual(test.expected, actual) {
				t.Logf("PASS: got %v", test.expected)
			} else {
				t.Fatalf("FAIL for %s: expected %v\n\nActual result was %v", test.ark, test.expected, actual)
			}
		}
	}
}
