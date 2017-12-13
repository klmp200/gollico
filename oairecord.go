// Package gollico provides functions for the APIs provided by the Bibliothèque Nationale de France
// on top of its Gallica digital library
package gollico

import (
	"encoding/xml"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
)

// Record is the Full record returned by the API
type Record struct {
	Identifier string   `xml:"notice>record>header>identifier"`
	Dewey      string   `xml:"dewey"`
	SDewey     string   `xml:"sdewey"`
	Typedoc    string   `xml:"typedoc"`
	Nqamoyen   string   `xml:"nqamoyen"`
	DCRecord   DCRecord `xml:"notice>record>metadata>dc"`
	Sounds     []Sound  `xml:"sounds>page"`
	VideoFile  string   `xml:"video>file"`
}

// DCRecord is the OAI-DC part of the record
type DCRecord struct {
	Title       string    `xml:"title"`
	Creator     string    `xml:"creator"`
	Contributor string    `xml:"contributor"`
	Description string    `xml:"description"`
	Subject     string    `xml:"subject"`
	Publisher   string    `xml:"publisher"`
	Date        string    `xml:"date"` // can we have a date.Date
	Format      string    `xml:"format"`
	Language    string    `xml:"language"`
	Relation    string    `xml:"relation"`
	DocTypes    []DocType `xml:"type"`
	Source      string    `xml:"source"`
	Rights      []Right   `xml:"rights"`
}

// DocType is the type of document concerned, as a text string in a given language
type DocType struct {
	Lang        string `xml:"lang,attr"`
	TypeDisplay string `xml:",chardata"`
}

// Right is the legal rights attached to the document, as a text string in a given language
type Right struct {
	Lang          string `xml:"lang,attr"`
	RightsDisplay string `xml:",chardata"`
}

// Sound gathers the fields specific to sound documents
type Sound struct {
	PageNum string `xml:"num,attr"`
	Title   string `xml:"media>title"`
	FileURL string `xml:"media>file"`
}

// GetOAIRecord retrieves a Bibliographic record from the library, using it's ID (ark number)
func GetOAIRecord(ark string) (Record, error) {

	r := Record{}
	if ark == "" {
		return r, errors.New("Missing required parameter ark: identifier")
	}

	resp, err := http.Get(BaseURL + "OAIRecord?ark=" + ark)
	if err != nil {
		fmt.Printf("error is %v", err)
	}

	if resp.StatusCode == http.StatusNotFound {
		return r, errors.New("document not found, might not be indexed in gallica")
	}

	if resp.StatusCode == http.StatusBadRequest {
		return r, errors.New("bad request, ark parameter might be missing")
	}

	b, err := ioutil.ReadAll(resp.Body)
	resp.Body.Close()

	err = xml.Unmarshal(b, &r)
	if err != nil {
		fmt.Printf("error: %v", err)
	}

	return r, nil

}
