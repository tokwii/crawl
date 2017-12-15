package queue

import (
"sync"
)

type localQueue struct {
	queue chan string
}

// Singleton Object
var queue *localQueue
var lsOnce sync.Once

func initLocalQueue(capacity int) *localQueue{
	lsOnce.Do(func(){
		queue = &localQueue{
			queue: make(chan string, capacity),
		}
	})
	return queue
}

func (q *localQueue) Push(task string){
	q.queue <- task
}

func (q *localQueue) Fetch() string{
	task := <- q.queue
	return task
}

func (q *localQueue) Len() int{
	return len(q.queue)
}

func (q *localQueue) Flush(){
	for i := 0; i < q.Len(); i++ {
		<-q.queue
	}
}

func (q *localQueue) Close(){
	close(q.queue)
}
