package main

import (
	"fmt"
	"github.com/tokwii/crawl/fetcher"
)

func main()  {
	result, err := fetcher.FetchURL("http://tomblomfield.com/", false)

	if err != nil {
	}

	fmt.Println(result.Scripts)
}
