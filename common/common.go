package common

import (
	"github.com/tokwii/crawl/fetcher"
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
	Styles []Style `xml:"style"`
	Scripts []Script `xml:"script"`
	Image []Image `xml:"image"`
}

type Sitemap struct {
	XMLName xml.Name `xml:"urlset"`
	URLS     []URL    `xml:"url"`
}

func FetcherResultMap(fetcherResult fetcher.Result) (map[string][]string){
	siteMetadata := make(map[string][]string)
	siteMetadata["links"] = fetcherResult.Links
	siteMetadata["images"] = fetcherResult.Images
	siteMetadata["styles"] = fetcherResult.Styles
	siteMetadata["scripts"] = fetcherResult.Images
	return siteMetadata
}