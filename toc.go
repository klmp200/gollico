// Package gollico provides functions for the APIs provided by the Bibliothèque Nationale de France
// on top of its Gallica digital library
package gollico

import (
	"encoding/json"
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
	Text string
	Href []href
}

type href struct {
	PageNumber string
	URL        string
}

// GetToc retrieves the Table of Content for a document
// returns it as json
func GetToc(ark string) ([]byte, error) {

	// Table of content to return
	toc := Toc{}
	var result []byte

	if ark == "" {
		return result, errors.New("Missing required parameter ark: identifier")
	}

	resp, err := http.Get(BaseURL + "Toc?ark=ark:/12148/" + ark)
	if err != nil {
		return result, err
	}

	if resp.StatusCode == http.StatusNotFound {
		return result, errors.New("document not found, might not be indexed in Gallica")
	}

	if resp.StatusCode != http.StatusOK {
		return result, errors.New("Status not OK")
	}

	b, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return result, err
	}
	resp.Body.Close()

	// create an xml tree
	doc := etree.NewDocument()
	// needed for non-UTF8 encoding of xml returned
	doc.ReadSettings.CharsetReader = func(label string, input io.Reader) (io.Reader, error) {
		return input, nil
	}

	// xml Decoder.Strict = false, needed for input containing common
	// mistakes, specifically unknown or malformed
	//character entities (sequences beginning with &)
	doc.ReadSettings.Permissive = true
	if err := doc.ReadFromBytes(b); err != nil {
		return result, err
	}
	root := doc.Root()

	// TEI or HTML
	switch root.Tag {
	case "TEI.2":
		err := toc.extractTEI(ark, doc, root)
		if err != nil {
			return result, err
		}
	case "html":
		err := toc.extractHTML(ark, doc, root)
		if err != nil {
			return result, err
		}
	default:
		return result, errors.New("Format returned unknown (neither TEI nor HTML)")
	}

	// marshalling the struct into json
	result, err = json.Marshal(toc)
	if err != nil {
		return result, err
	}

	return result, nil
}

// extractTEI parses the TEI response into a Toc struct
func (toc *Toc) extractTEI(ark string, doc *etree.Document, root *etree.Element) error {
	for _, row := range root.FindElements("//row") {
		tocEntry := TocEntry{}
		refs := []href{}
		// regex to extract 59 from "FOREIGN(9754046/000059.jp2)"
		// i.e. the page reference to generate the right URL for each ToC Entry
		var pRef = regexp.MustCompile(`/([0-9]+)\.`)

		for _, cell := range row.FindElements("cell/*") {
			switch tag := cell.Tag; tag {
			case "seg":
				tocEntry.Text = cell.Text()
			case "xref":

				fromAttr := cell.SelectAttrValue("from", "")

				// regex to extract 59 from "FOREIGN(9754046/000059.jp2)"
				res := pRef.FindStringSubmatch(fromAttr)
				// our regex didn't catch a group
				if len(res) < 2 {
					continue
				}
				pNum := strings.TrimLeft(res[1], "0")
				URL := "http://gallica.bnf.fr/ark:/12148/" + ark + "/f" + pNum

				thisRef := href{
					PageNumber: cell.Text(),
					URL:        URL,
				}
				refs = append(refs, thisRef)
			}
		}

		tocEntry.Href = refs
		toc.TocEntries = append(toc.TocEntries, tocEntry)
	}
	if len(toc.TocEntries) == 0 {
		return errors.New("There were no entries in this table of contents")
	}

	return nil
}

// TODO: extractHTML
func (toc *Toc) extractHTML(ark string, doc *etree.Document, root *etree.Element) error {
	for _, row := range root.FindElements("//div[@class='Texte']") {
		tocEntry := TocEntry{}
		tocEntry.Text = row.Text()
		refs := []href{}

		// regex: 152 from "javascript:allerA('0083037', '152')"
		var pRef = regexp.MustCompile(`'(\d+)'\)`)
		a := row.SelectElements("a")
		for _, entry := range a {
			if entry != nil {

				attr := entry.SelectAttrValue("href", "")
				res := pRef.FindStringSubmatch(attr)
				// our regex didn't catch a group
				if len(res) < 2 {
					continue
				}
				URL := "http://gallica.bnf.fr/ark:/12148/" + ark + "/f" + res[1]

				thisRef := href{
					PageNumber: entry.Text(),
					URL:        URL,
				}
				refs = append(refs, thisRef)
			}
		}
		tocEntry.Href = refs
		toc.TocEntries = append(toc.TocEntries, tocEntry)
	}
	return nil
}
