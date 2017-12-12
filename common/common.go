package common

import (
	"github.com/tokwii/crawl/fetcher"
)

func FetcherResultMap(fetcherResult fetcher.Result) (map[string][]string){
	siteMetadata := make(map[string][]string)
	siteMetadata["links"] = fetcherResult.Links
	siteMetadata["images"] = fetcherResult.Images
	siteMetadata["styles"] = fetcherResult.Styles
	siteMetadata["scripts"] = fetcherResult.Images
	return siteMetadata
}