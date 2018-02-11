// Package gollico provides functions for the APIs provided by the Bibliothèque Nationale de France
// on top of its Gallica digital library
package gollico

import (
	"errors"
	"reflect"
	"testing"
)

type getTocTest struct {
	Ark      string
	Expected []byte
	Err      error
}

var tocTEI = `{"TocEntries":[{"Text":"I. - Le Festin","Href":[{"PageNumber":"1","URL":"http://gallica.bnf.fr/ark:/12148/bpt6k61076295/f8"}]},{"Text":"II. - A Sicca","Href":[{"PageNumber":"22","URL":"http://gallica.bnf.fr/ark:/12148/bpt6k61076295/f29"}]},{"Text":"III. - Salammbô","Href":[{"PageNumber":"47","URL":"http://gallica.bnf.fr/ark:/12148/bpt6k61076295/f54"}]},{"Text":"IV. - Sous les murs de Carthage","Href":[{"PageNumber":"56","URL":"http://gallica.bnf.fr/ark:/12148/bpt6k61076295/f63"}]},{"Text":"V. - Tanit","Href":[{"PageNumber":"77","URL":"http://gallica.bnf.fr/ark:/12148/bpt6k61076295/f84"}]},{"Text":"VI - Hannon","Href":[{"PageNumber":"95","URL":"http://gallica.bnf.fr/ark:/12148/bpt6k61076295/f102"}]},{"Text":"VII. - Hamilcar Barca","Href":[{"PageNumber":"118","URL":"http://gallica.bnf.fr/ark:/12148/bpt6k61076295/f125"}]},{"Text":"VIII. - La Bataille du Macar","Href":[{"PageNumber":"160","URL":"http://gallica.bnf.fr/ark:/12148/bpt6k61076295/f167"}]},{"Text":"IX. - En Campagne","Href":[{"PageNumber":"181","URL":"http://gallica.bnf.fr/ark:/12148/bpt6k61076295/f188"}]},{"Text":"X. - Le Serpent","Href":[{"PageNumber":"198","URL":"http://gallica.bnf.fr/ark:/12148/bpt6k61076295/f205"}]},{"Text":"XI. - Sous la Tente","Href":[{"PageNumber":"213","URL":"http://gallica.bnf.fr/ark:/12148/bpt6k61076295/f220"}]},{"Text":"XII. - L'Aqueduc","Href":[{"PageNumber":"235","URL":"http://gallica.bnf.fr/ark:/12148/bpt6k61076295/f242"}]},{"Text":"XIII. - Moloch","Href":[{"PageNumber":"258","URL":"http://gallica.bnf.fr/ark:/12148/bpt6k61076295/f265"}]},{"Text":"XIV. - Le Défilé de la Hache","Href":[{"PageNumber":"299","URL":"http://gallica.bnf.fr/ark:/12148/bpt6k61076295/f306"}]},{"Text":"XV. - Mâtho","Href":[{"PageNumber":"342","URL":"http://gallica.bnf.fr/ark:/12148/bpt6k61076295/f349"}]},{"Text":"APPENDICE","Href":[{"PageNumber":"353","URL":"http://gallica.bnf.fr/ark:/12148/bpt6k61076295/f360"}]}]}`

func TestGetToc(t *testing.T) {
	errEmptyArk := errors.New("Missing required parameter ark: identifier")
	//	errWrongArk := errors.New("document not found, might not be indexed in Gallica")

	var GetTocTests = []getTocTest{
		{
			Ark:      "",
			Expected: nil,
			Err:      errEmptyArk,
		},
		{
			Ark:      "bpt6k61076295", // Should be tei toc
			Expected: []byte(tocTEI),
			Err:      nil,
		},
	}

	for _, test := range GetTocTests {
		actual, _ := GetToc(test.Ark)
		if reflect.DeepEqual(test.Expected, actual) {
			t.Logf("PASS: got %v", test.Expected)
		} else {
			t.Fatalf("FAIL for %s: expected %v, actual result was %v", test.Ark, string(test.Expected), string(actual))
		}
	}
}
