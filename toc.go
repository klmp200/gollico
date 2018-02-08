// Package gollico provides functions for the APIs provided by the Bibliothèque Nationale de France
// on top of its Gallica digital library
package gollico

import (
	"errors"
	"io"
	"io/ioutil"
	"net/http"
	"regexp"
	"strings"

	"github.com/beevik/etree"
)

// Toc stores the Table of Content for a document
type Toc struct {
	TocEntries []TocEntry
}

// TocEntry stores an entry in a Table of Content
type TocEntry struct {
	PageNumber string
	Text       string
	URL        string
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
		return toc, err
	}

	if resp.StatusCode == http.StatusNotFound {
		return toc, errors.New("document not found, might not be indexed in Gallica")
	}

	if resp.StatusCode == http.StatusBadRequest {
		return toc, errors.New("bad request, ark parameter might be missing")
	}

	if resp.StatusCode != http.StatusOK {
		return toc, errors.New("Status not OK")
	}

	b, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return toc, err
	}
	resp.Body.Close()

	// create an xml tree
	doc := etree.NewDocument()
	// needed for non-UTF8 encoding of xml returned
	doc.ReadSettings.CharsetReader = func(label string, input io.Reader) (io.Reader, error) {
		return input, nil
	}

	if err := doc.ReadFromBytes(b); err != nil {
		return toc, err
	}
	root := doc.Root()

	// TEI or HTML
	switch root.Tag {
	case "TEI.2":
		err := toc.extractTEI(ark, doc, root)
		if err != nil {
			return toc, err
		}
	case "html":
		err := toc.extractHTML(ark, doc, root)
		if err != nil {
			return toc, err
		}
	default:
		return toc, errors.New("Format returned unknown (neither TEI nor HTML)")
	}

	return toc, nil
}

func (toc *Toc) extractTEI(ark string, doc *etree.Document, root *etree.Element) error {
	for _, row := range root.FindElements("//row") {
		tocEntry := TocEntry{}

		// we're going to use this regex to extract
		// the page reference to generate the right URL for each ToC Entry
		var pRef = regexp.MustCompile(`/([0-9]+)\.`)

		for _, cell := range row.FindElements("cell/*") {
			switch tag := cell.Tag; tag {
			case "seg":
				tocEntry.Text = cell.Text()
			case "xref":
				tocEntry.PageNumber = cell.Text()

				fromAttr := cell.SelectAttrValue("from", "")

				// regex to extract 59 from "FOREIGN(9754046/000059.jp2)"
				res := pRef.FindStringSubmatch(fromAttr)
				// our regex didn't catch a group
				if len(res) < 2 {
					continue
				}
				pNum := strings.TrimLeft(res[1], "0")
				pURL := "http://gallica.bnf.fr/ark:/12148/" + ark + "/f" + pNum
				tocEntry.URL = pURL
			}
		}
		toc.TocEntries = append(toc.TocEntries, tocEntry)
		// TODO: check order of entries in slice
	}
	if len(toc.TocEntries) == 0 {
		return errors.New("There was no entries in this table of contents")
	}

	return nil
}

func (toc *Toc) extractHTML(ark string, doc *etree.Document, root *etree.Element) error {

	return nil
}
