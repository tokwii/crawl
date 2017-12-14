package scheduler

import (
	"sync"
	"github.com/tokwii/crawl/queue"
	"github.com/tokwii/crawl/fetcher"
	"github.com/tokwii/crawl/common"
	"github.com/tokwii/crawl/storage"
	"fmt"
)

// Make this private
type Scheduler struct{
	// TODO Concurrency Issue? User Mutex?
	AggrResult *storage.CrawlerStorage
	taskQueue *queue.TaskQueue
	numWorkers int
}
// Output to a Channel or Sharded Dastructure!

// Global Variable
/*var taskQueue = queue.TaskQueue(10000)

func Schedule(numbThreads int, q *queue.TaskQueue){

}*/

//1 ) Iterate through the Chanel starting of a go rountine <-

func InitSchedule(numWorkers int, taskQCapacity int, seedUrl string)(*Scheduler){
	var s Scheduler
	fmt.Println("Numbere of workers ...")
	fmt.Println(numWorkers)
	s.taskQueue = queue.InitTaskQueue(taskQCapacity)
	s.AggrResult = storage.InitCrawlerStorage()
	s.taskQueue.Push(queue.Task{URL:seedUrl})
	s.numWorkers = numWorkers
	return &s
}

func (s *Scheduler) Schedule(){
	s.initCrawlWorkerPool()
	s.taskQueue.Close()
}

func (s *Scheduler) initCrawlWorkerPool(){
	var wg sync.WaitGroup

	for i := 0; i < s.numWorkers; i++ {
		wg.Add(1)
		go s.crawlWorker(&wg)
	}
	wg.Wait()
}

func (s *Scheduler) crawlWorker(wg *sync.WaitGroup){

	for i := 0; i < s.taskQueue.Len(); i++ {
		task := s.taskQueue.Fetch()
		ok := s.AggrResult.Contains(task.URL)
		if !ok {
			fmt.Println("Currently Crawled Url")
			fmt.Println(task.URL)
			result, err := fetcher.FetchURL(task.URL, false, s.taskQueue)
			if err != nil {
				continue
			}

			siteMetadata := common.FetcherResultMap(result)
			s.AggrResult.Add(task.URL, siteMetadata)
		}
	}
	defer wg.Done()
}
