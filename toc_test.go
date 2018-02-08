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
				PageNumber: "1",
				Text:       "I. - Le Festin",
				URL:        "http://gallica.bnf.fr/ark:/12148/bpt6k61076295/f8",
			},
			TocEntry{
				PageNumber: "22",
				Text:       "II. - A Sicca",
				URL:        "http://gallica.bnf.fr/ark:/12148/bpt6k61076295/f29",
			},
			TocEntry{
				PageNumber: "47",
				Text:       "III. - Salammbô",
				URL:        "http://gallica.bnf.fr/ark:/12148/bpt6k61076295/f54",
			},
			TocEntry{
				PageNumber: "56",
				Text:       "IV. - Sous les murs de Carthage",
				URL:        "http://gallica.bnf.fr/ark:/12148/bpt6k61076295/f63",
			},
			TocEntry{
				PageNumber: "77",
				Text:       "V. - Tanit",
				URL:        "http://gallica.bnf.fr/ark:/12148/bpt6k61076295/f84",
			},
			TocEntry{
				PageNumber: "95",
				Text:       "VI - Hannon",
				URL:        "http://gallica.bnf.fr/ark:/12148/bpt6k61076295/f102",
			},
			TocEntry{
				PageNumber: "118",
				Text:       "VII. - Hamilcar Barca",
				URL:        "http://gallica.bnf.fr/ark:/12148/bpt6k61076295/f125",
			},
			TocEntry{
				PageNumber: "160",
				Text:       "VIII. - La Bataille du Macar",
				URL:        "http://gallica.bnf.fr/ark:/12148/bpt6k61076295/f167",
			},
			TocEntry{
				PageNumber: "181",
				Text:       "IX. - En Campagne",
				URL:        "http://gallica.bnf.fr/ark:/12148/bpt6k61076295/f188",
			},
			TocEntry{
				PageNumber: "198",
				Text:       "X. - Le Serpent",
				URL:        "http://gallica.bnf.fr/ark:/12148/bpt6k61076295/f205",
			},
			TocEntry{
				PageNumber: "213",
				Text:       "XI. - Sous la Tente",
				URL:        "http://gallica.bnf.fr/ark:/12148/bpt6k61076295/f220",
			},
			TocEntry{
				PageNumber: "235",
				Text:       "XII. - L'Aqueduc",
				URL:        "http://gallica.bnf.fr/ark:/12148/bpt6k61076295/f242",
			},
			TocEntry{
				PageNumber: "258",
				Text:       "XIII. - Moloch",
				URL:        "http://gallica.bnf.fr/ark:/12148/bpt6k61076295/f265",
			},
			TocEntry{
				PageNumber: "299",
				Text:       "XIV. - Le Défilé de la Hache",
				URL:        "http://gallica.bnf.fr/ark:/12148/bpt6k61076295/f306",
			},
			TocEntry{
				PageNumber: "342",
				Text:       "XV. - Mâtho",
				URL:        "http://gallica.bnf.fr/ark:/12148/bpt6k61076295/f349",
			},
			TocEntry{
				PageNumber: "353",
				Text:       "APPENDICE",
				URL:        "http://gallica.bnf.fr/ark:/12148/bpt6k61076295/f360",
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
			Ark:      "bpt6k61076295", // Should be tei toc
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
