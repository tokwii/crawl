package queue

import (
	"github.com/tokwii/crawl/config"
	"sync"
)

const LOCAL_QUEUE = "local"

const REMOTE_QUEUE = "remote"

type Queue interface {
	Push(string)
	Fetch() string
	Len() int
	Close()
	Flush()
}

type CrawlerQueue struct {
	Mode string
	queue Queue
}

// SingletonStore
var crawlerQueue *CrawlerQueue
var once sync.Once

func InitCrawlerQueue() *CrawlerQueue{
	once.Do(func(){
		crawlerQueue = build()
	})
	return crawlerQueue
}

func (q *CrawlerQueue) Push(task string){
	q.queue.Push(task)
}

func (q *CrawlerQueue) Fetch() (string) {
	return q.queue.Fetch()
}

func (q *CrawlerQueue) Len() (int) {
	return q.queue.Len()
}
func (q *CrawlerQueue) Close() {
	q.queue.Close()
}

func (q *CrawlerQueue) Flush() {
	q.queue.Flush()
}

func build() *CrawlerQueue{
	cq := CrawlerQueue{}

	queueMode := config.Conf.Queue.Mode
	switch queueMode {
	case LOCAL_QUEUE:
		ls := initLocalQueue(config.Conf.Queue.Local.Capacity)
		cq.Mode = queueMode
		cq.queue = ls
	case REMOTE_QUEUE:
		//Remote
	default:
		// Defaults to Local Storage (For Test Too)
		ls := initLocalQueue(100000)
		cq.Mode = queueMode
		cq.queue = ls
	}
	return  &cq
}