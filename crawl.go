package main

import (
	"fmt"
	"github.com/tokwii/crawl/fetcher"
	"github.com/tokwii/crawl/queue"
)

func main()  {
	taskQueue := queue.InitTaskQueue(10)
	result, err := fetcher.FetchURL("http://tomblomfield.com/", false, taskQueue)
	taskQueue.Flush()

	if err != nil {
	}

	fmt.Println(result.Scripts)
}
