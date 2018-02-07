// Package gollico provides functions for the APIs provided by the Bibliothèque Nationale de France
// on top of its Gallica digital library
package gollico

import (
	"reflect"
	"testing"
)

type getTocTest struct {
	Ark      string
	Expected Toc
	Err      error
}

func TestGetToc(t *testing.T) {
	testToc1 := Toc{
		[]TocEntry{
			TocEntry{
				PageNumber: "2",
				Text:       "EntryText",
				URL:        "entryurl",
			},
		},
	}
	var GetTocTests = []getTocTest{
		/*{
			Ark:      "bpt6k83037p​", // Should be html toc
			Expected: testToc1,
			Err:      nil,
		},*/
		{
			Ark:      "bpt6k97540464", // Should be tei toc
			Expected: testToc1,
			Err:      nil,
		},
	}

	for _, test := range GetTocTests {
		actual, _ := GetToc(test.Ark)
		if reflect.DeepEqual(test.Expected, actual) {
			t.Logf("PASS: got %v", test.Expected)
		} else {
			t.Fatalf("FAIL for %s: expected %v, actual result was %v", test.Ark, test.Expected, actual)
		}
	}
}
