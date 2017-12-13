// Package gollico provides functions for the APIs provided by the Bibliothèque Nationale de France
// on top of its Gallica sigital library
package gollico

import (
	"encoding/xml"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
)

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

type DocType struct {
	Lang        string `xml:"lang,attr"`
	TypeDisplay string `xml:",chardata"`
}

type Right struct {
	Lang          string `xml:"lang,attr"`
	RightsDisplay string `xml:",chardata"`
}

type Sound struct {
	PageNum string `xml:"num,attr"`
	Title   string `xml:"media>title"`
	FileURL string `xml:"media>file"`
}

func GetOAIRecord(ark string) (Record, error) {

	r := Record{}
	if ark == "" {
		return r, errors.New("Missing required parameter ark: identifier")
	}

	resp, err := http.Get(BaseURL + "OAIRecord?ark=" + ark)
	if err != nil {
		fmt.Printf("error is %v", err)
	}

	b, err := ioutil.ReadAll(resp.Body)
	resp.Body.Close()

	err = xml.Unmarshal(b, &r)
	if err != nil {
		fmt.Printf("error: %v", err)
	}

	return r, nil

}
