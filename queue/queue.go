package queue

import (
	"sync"
)

type TaskQueue struct {
	queue chan Task
}

// Singleton Object
var queue *TaskQueue
var once sync.Once

// Initialize Singleton Task Queue. Thread Safe. HA for Scheduler
func InitTaskQueue(length int) *TaskQueue{
	once.Do(func(){
		queue = &TaskQueue{
			queue: make(chan Task, length),
		}
	})
	return queue
}

func (q *TaskQueue) Push(task Task){
	q.queue <- task
}

func (q *TaskQueue) Fetch() Task{
	task := <- q.queue
	return task
}

func (q *TaskQueue) Len() int{
	return len(q.queue)
}

func (q *TaskQueue) Flush(){
	for i := 0; i < q.Len(); i++ {
		<-q.queue
	}
}

func (q *TaskQueue) Close(){
	close(q.queue)
}