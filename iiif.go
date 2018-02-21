package gollico

import (
	"fmt"
	"log"
	"strconv"
)

var formats = []string{
	"jpg",
	"tif",
	"png",
	"gif",
	"jp2",
	"pdf",
	"webp",
}

// Params struct with info to pass to IIIF server
// required fields : Ark, PageID
// other fields have default values if sent as nil
type Params struct {
	Ark      string
	PageID   string
	RegionX  int
	RegionY  int
	Width    int
	Height   int
	Size     int
	Rotation int
	Quality  string
	Format   string
}

// GetImage retrieves an image from the IIIF Server
// requires a Params struct
// required fields : Ark, PageID
// other fields have default values if sent as nil
func (iDoc IIIFDoc) GetImage(p Params) {
	if p.Ark == "" || p.PageID == "" {
		log.Fatalf("missing required param (Ark, PageID")
	}

	var region string
	if p.RegionX > 0 && p.RegionY > 0 && p.Width > 0 && p.Height > 0 {
		region = strconv.Itoa(p.RegionX) + "," + strconv.Itoa(p.RegionY) +
			"," + strconv.Itoa(p.Width) + "," + strconv.Itoa(p.Height)
	} else {
		region = "full"
	}

	var size string
	if p.Size != 0 {
		size = strconv.Itoa(p.Size)
	} else {
		size = "full"
	}

	var rotation string
	if p.Rotation <= 0 || p.Rotation >= 360 {
		rotation = "0"
	} else {
		rotation = strconv.Itoa(p.Rotation)
	}

	// color	The image is returned in full color.
	// gray	The image is returned in grayscale, where each pixel is black, white or any shade of gray in between.
	// bitonal	The image returned is bitonal, where each pixel is either black or white.
	// default
	var quality string
	if p.Quality != "color" && p.Quality != "gray" && p.Quality != "bitonal" {
		quality = "native"
	} else {
		quality = p.Quality
	}

	var format string
	for _, v := range formats {
		if p.Format == v {
			format = p.Format
		} else {
			format = "jpg" // default to jpg, why not?
		}
	}

	url := BaseURLIIIF + p.Ark + "/" + p.PageID + "/" + region + "/" + size + "/" + rotation + "/" + quality + "." + format
	fmt.Println(url)
}
