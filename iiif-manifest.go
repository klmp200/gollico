// Package gollico provIDes functions for the APIs provIDed by the Bibliothèque Nationale de France
// on top of its Gallica digital library
// see http://api.bnf.fr/api-iiif-de-recuperation-des-images-de-gallica
package gollico

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
)

// IIIFDoc is a complete document returned by the IIIF API
type IIIFDoc struct {
	ID           string
	Label        string
	Attribution  string
	License      string
	Related      string
	Description  string
	IIIFMetadata IIIFMetadata
	Images       []Image
	Thumbnail    string
}

// IIIFMetadata is the embedded Metadata struct of a document
type IIIFMetadata struct {
	Repository     string
	Provider       string
	Disseminator   string
	SourceImages   string
	MetadataSource string
	Shelfmark      string
	Title          string
	Date           string
	Language       []string
	Format         []string
	Creator        string
	Relation       string
	Type           map[string][]string
}

// Image is the embedded image in a document
type Image struct {
	ID        string
	Label     string
	Height    float64
	Width     float64
	Format    string
	Thumbnail string
}

// GetIIIFDoc retrieves a document struct from the API
func GetIIIFDoc(ark string) (IIIFDoc, error) {
	doc := IIIFDoc{}

	if ark == "" {
		return doc, errors.New("Missing required parameter ark: IDentifier")
	}

	url := BaseURLIIIF + ark + "/manifest.json"
	resp, err := http.Get(url)
	if err != nil {
		fmt.Println(err)
		return doc, err
	}

	if resp.StatusCode != http.StatusOK {
		fmt.Println(err)
		return doc, errors.New("Status not OK")
	}

	b, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println(err)
		return doc, err
	}
	resp.Body.Close()

	// we unmarshal into a map[string]interface{}
	// rather than directly into the IIIFDoc struct
	// because we want to work on some of the fields / values beforehand
	var f interface{}
	err = json.Unmarshal(b, &f)
	if err != nil {
		fmt.Println(err)
		return doc, err
	}
	m := f.(map[string]interface{})

	// First level fields
	doc.Description = m["description"].(string)
	doc.ID = m["@id"].(string)
	doc.Label = m["label"].(string)
	doc.Attribution = m["attribution"].(string)
	doc.License = m["license"].(string)
	doc.Related = m["related"].(string)
	docThumb := m["thumbnail"].(map[string]interface{})
	doc.Thumbnail = docThumb["@id"].(string)

	// metadata
	// TODO: separate func
	for _, v := range m["metadata"].([]interface{}) {
		entry := v.(map[string]interface{})

		switch entry["label"] {
		case "Repository":
			doc.IIIFMetadata.Repository = entry["value"].(string)
		case "Provider":
			doc.IIIFMetadata.Provider = entry["value"].(string)
		case "Disseminator":
			doc.IIIFMetadata.Disseminator = entry["value"].(string)
		case "Source Images":
			doc.IIIFMetadata.SourceImages = entry["value"].(string)
		case "Metadata Source":
			doc.IIIFMetadata.MetadataSource = entry["value"].(string)
		case "Shelfmark":
			doc.IIIFMetadata.Shelfmark = entry["value"].(string)
		case "Title":
			doc.IIIFMetadata.Title = entry["value"].(string)
		case "Date":
			doc.IIIFMetadata.Date = entry["value"].(string)
		case "Creator":
			doc.IIIFMetadata.Creator = entry["value"].(string)
		case "Relation":
			doc.IIIFMetadata.Relation = entry["value"].(string)
		case "Language":
			// TODO: language and Format below are the same -> abstract
			languages := []string{}
			lang := entry["value"].([]interface{})
			for _, l := range lang {
				for _, ll := range l.(map[string]interface{}) {
					languages = append(languages, ll.(string))
				}
			}
			doc.IIIFMetadata.Language = languages
		case "Format":
			formats := []string{}
			format := entry["value"].([]interface{})
			for _, f := range format {
				for _, ff := range f.(map[string]interface{}) {
					formats = append(formats, ff.(string))
				}
			}
			doc.IIIFMetadata.Format = formats
		case "Type":
			typesMap := map[string][]string{}
			tDict := entry["value"].([]interface{})
			for _, t := range tDict {
				//		fmt.Println(t)
				tt := t.(map[string]interface{})
				typesMap[tt["@language"].(string)] = append(typesMap[tt["@language"].(string)], tt["@value"].(string))
			}
			doc.IIIFMetadata.Type = typesMap
		default:
			fmt.Printf("unknown entry: %v\n\n", entry["label"])
		}

	}

	// images
	// TODO: separate func
	doc.Images = []Image{}
	for _, seq := range m["sequences"].([]interface{}) {
		v := seq.(map[string]interface{})
		canvases := v["canvases"].([]interface{})

		for _, c := range canvases {
			// each c is an image
			img := Image{}
			cc := c.(map[string]interface{})
			img.Label = cc["label"].(string)
			img.ID = cc["@id"].(string)
			img.Height = cc["height"].(float64)
			img.Width = cc["width"].(float64)
			thumbnail := cc["thumbnail"].(map[string]interface{})
			img.Thumbnail = thumbnail["@id"].(string)
			doc.Images = append(doc.Images, img)
		}
	}

	// TODO: remove scaffold
	/*	s := reflect.ValueOf(&doc).Elem()
		typeOfT := s.Type()

		for i := 0; i < s.NumField(); i++ {
			f := s.Field(i)
			fmt.Printf("%d: %s %s = %v\n", i,
				typeOfT.Field(i).Name, f.Type(), f.Interface())
		}
	*/
	return doc, nil
}

// GetIIIFDocMetadata returns the Metadata for a document
func GetIIIFDocMetadata(ark string) ([]byte, error) {

	if ark == "" {
		return nil, errors.New("Missing required parameter ark: Identifier")
	}

	iDoc, err := GetIIIFDoc(ark)
	if err != nil {
		return nil, err
	}

	data, err := json.Marshal(iDoc.IIIFMetadata)
	if err != nil {
		return nil, err
	}
	return data, nil

}

// GetThumbnails retrieves the thumbnails for all the images in a document
func (iDoc IIIFDoc) GetThumbnails() []string {
	var result []string
	for _, img := range iDoc.Images {
		result = append(result, img.Thumbnail)
	}
	return result
}

// GetCoverThumbnail retrieves the thumbnail for the cover of a document
func (iDoc IIIFDoc) GetCoverThumbnail() string {
	return iDoc.Thumbnail
}

// ImagesCount returns the number of images in a document
func (iDoc IIIFDoc) ImagesCount() int {
	return len(iDoc.Images)
}

// GetImagesData retrieves the data for all the images in a document
func (iDoc IIIFDoc) GetImagesData() ([]byte, error) {

	data, err := json.Marshal(iDoc.Images)
	if err != nil {
		return nil, err
	}
	return data, nil
}
