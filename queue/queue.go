package queue

import (
	"sync"
)

type taskQueue struct {
	queue chan Task
}

var queue *taskQueue
var once sync.Once

// Initialize Singleton Task Queue. Thread Safe. HA for Scheduler
func TaskQueue(length int) *taskQueue{

	once.Do(func(){
		queue = &taskQueue{
			queue: make(chan Task, length),
		}
	})
	return queue
}

func (q *taskQueue) Push(task Task){
	q.queue <- task
}

func (q *taskQueue) Fetch() Task{
	task := <- q.queue
	return task
}

func (q *taskQueue) Len() int{
	return len(q.queue)
}

func (q *taskQueue) Close(){
	close(q.queue)
}