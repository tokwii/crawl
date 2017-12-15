package common

import (
	"encoding/xml"
)
// Add Last modified field
type Script struct{
	Loc string `xml:"script:loc,omitempty"`
}

type Image struct{
	Loc string `xml:"image:loc,omitempty"`
}

type Style struct{
	Loc string `xml:"style:loc,omitempty"`
}

type URL struct {
	Loc string `xml:"loc"`
	Styles []Style `xml:"style:style"`
	Scripts []Script `xml:"script:script"`
	Images []Image `xml:"image:image"`
}

type Sitemap struct {
	XMLName xml.Name `xml:"urlset"`
	URLS     []URL    `xml:"url"`
}
