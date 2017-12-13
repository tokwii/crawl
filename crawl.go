package main

import (
	//"fmt"
	//"github.com/tokwii/crawl/fetcher"
	//"github.com/tokwii/crawl/queue"
	//"github.com/tokwii/crawl/scheduler"
	"github.com/tokwii/crawl/config"
	//"time"
	"fmt"
)

func main()  {
	// If the Buffer is Smaller than the links found. It will be blocked ... To try to read from the Buffer
	// Faster than we write... The Read on Buffer blocks when it is full
	// Recieve on the buffer is is blocked when it is emtpy
	// We need some sort of Pool of fetchers to read continoulys for the worker Q until empty
	// Crazy !! Create a go routine for each taskQ
	//taskQueue := queue.InitTaskQueue(1000)
	//taskQueue.Push()
	/*result, err := fetcher.FetchURL("http://tomblomfield.com", false, taskQueue)

	taskQueue.Flush()
	taskQueue.Close()

	if err != nil {
	}

	fmt.Println(result.Links)*/


	//numWorkers int, taskQCapacity int, seedUrl string
	/*startTime := time.Now()
	s := scheduler.InitSchedule(10, 1000000, "http://cnn.com")
	s.Schedule()
	//fmt.Println(s.GetAggregateResults())
	endTime := time.Now()
	diff := endTime.Sub(startTime)
	fmt.Println("total time taken ", diff.Seconds(), "seconds")*/

	// Config ..... Config ....
	var c config.Config
	err := c.Load("config/settings.toml")
	if err != nil {
		fmt.Println(err)
	}
	//fmt.Println(c.Queue.Mode)
	fmt.Println(c.Queue.Mode)
	fmt.Println(c.Queue.Remote)

}
