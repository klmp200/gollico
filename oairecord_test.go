package gollico

import (
	"reflect"
	"testing"
)

type getOAIRecordTest struct {
	Ark      string
	Expected Record
}

func TestGetOAIRecord(t *testing.T) {

	testRecordSound := Record{
		Identifier: "oai:bnf.fr:gallica/ark:/12148/bpt6k1279113",
		Typedoc:    "sonore",
		Nqamoyen:   "0.0",
		DCRecord: DCRecord{
			Title:       "[Archives de la parole]. , Discours d'inauguration des Archives de la parole / M. Brunot, aut., participant",
			Creator:     "Brunot, Ferdinand (1860-1938). Auteur du texte",
			Contributor: "Brunot, Ferdinand (1860-1938). Participant",
			Publisher:   "Université de Paris, Archives de la parole (Paris)",
			Date:        "1911",
			Description: "Enregistrement : (France) Paris, Université de Paris, La Sorbonne, 03-06-1911",
			Subject:     "enregistrement parlé",
			Format:      "multipart/mixed",
			Language:    "français",
			Relation:    "Notice du catalogue : http://catalogue.bnf.fr/ark:/12148/cb385163820",
			DocTypes: []DocType{
				{
					Lang:        "eng",
					TypeDisplay: "sound",
				},
				{
					Lang:        "fre",
					TypeDisplay: "document sonore",
				},
			},
			Source: "Bibliothèque nationale de France, département Audiovisuel, AP-3",
			Rights: []Right{
				{
					Lang:          "fre",
					RightsDisplay: "domaine public",
				},
				{
					Lang:          "eng",
					RightsDisplay: "public domain",
				},
			},
		},
		Sounds: []Sound{
			{
				PageNum: "0",
				Title:   "Plage 1",
				FileURL: "http://gallica.bnf.fr/ark:/12148/bpt6k1279113/f1.audio",
			},
			{
				PageNum: "1",
				Title:   "Plage 1",
				FileURL: "http://gallica.bnf.fr/ark:/12148/bpt6k1279113/f2.audio",
			},
		},
	}

	var GetOAIRRecordTests = []getOAIRecordTest{
		{
			Ark:      "bpt6k1279113",
			Expected: testRecordSound,
		},
	}
	
	for _, test := range GetOAIRRecordTests {
		actual, _ := GetOAIRecord(test.Ark)
		if reflect.DeepEqual(test.Expected, actual) {
			t.Logf("PASS: got %v", test.Expected)
		} else {
			t.Fatalf("FAIL for %s: expected %v, actual result was %v", test.Ark, test.Expected, actual)
		}
	}
}
