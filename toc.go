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

// Toc stores the Table of Content for a document
type Toc struct {
	TocEntries []TocEntry
}

// TocEntry stores an entry in a Table of Content
type TocEntry struct {
	PageOriginal string
	EntryNumber  string
	EntryText    string
	EntryURL     string
}

// TocTEI used to unmarshall TEI Toc
type TocTEI struct {
	TocTEIEntries []TocTEIEntry `xml:"body>div0>div1>table>row"`
}

// TocTEIEntry used to unmarshall individual TEI Entries
type TocTEIEntry struct {
	TeiPO   string `xml:"cell>xref"`
	TeiET   string `xml:"cell>seg"`
	TeiEURL string `xml:"cell>xref from,attr"`
}

// TocHTML used to unmarshall HTML Toc
type TocHTML struct {
	TocHTMLEntries []TocHTMLEntry
}

// TocHTMLEntry used to unmarshall HTML Entries
type TocHTMLEntry struct {
}

// GetToc retrieves the Table of Content for a document
func GetToc(ark string) (Toc, error) {

	// Table of content to return
	toc := Toc{}

	if ark == "" {
		return toc, errors.New("Missing required parameter ark: identifier")
	}

	resp, err := http.Get(BaseURL + "Toc?ark=ark:/12148/" + ark)
	if err != nil {
		fmt.Printf("error is %v", err)
	}

	if resp.StatusCode == http.StatusNotFound {
		return toc, errors.New("document not found, might not be indexed in gallica")
	}

	if resp.StatusCode == http.StatusBadRequest {
		return toc, errors.New("bad request, ark parameter might be missing")
	}

	b, err := ioutil.ReadAll(resp.Body)
	resp.Body.Close()
	if err != nil {
		return toc, err
	}

	// sniff mime type
	ct := http.DetectContentType(b)

	// XML TEI
	if ct[5:8] == "xml" {
		err := parseTEIToc(b, &toc)
		if err != nil {
			return toc, err
		}
		return toc, nil
	}

	// HTML
	err = parseHTMLToc(b, &toc)
	if err != nil {
		return toc, err
	}

	return toc, nil
}

func parseHTMLToc(b []byte, toc *Toc) error {

	return nil
}

func parseTEIToc(b []byte, toc *Toc) error {

	toctei := TocTEI{}

	err := xml.Unmarshal(b, &toctei)
	if err != nil {
		fmt.Printf("error: %v", err)
	}
	fmt.Printf("====\n%v\n", string(b))
	fmt.Printf("TOCTEI : \n%v\n", toctei)

	return nil
}
