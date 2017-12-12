package queue

import (
	"github.com/stretchr/testify/suite"
	"testing"
)

type TaskQueueTestSuite struct {
	suite.Suite
	taskQ *TaskQueue
	capacity int
	task Task
}

func (suite *TaskQueueTestSuite) SetupSuite(){
	suite.capacity = 1000
	suite.taskQ = InitTaskQueue(suite.capacity)
	suite.task = Task{
		URL: "https://google.com",
	}
}

func (suite *TaskQueueTestSuite) TestSingletonTaskQueue(){
	taskQ2 := InitTaskQueue(10)
	taskQ2.Push(suite.task)
	suite.Equal(taskQ2, suite.taskQ)
	t2 := suite.taskQ.Fetch()
	suite.Equal(suite.task.URL, t2.URL)
}

func (suite *TaskQueueTestSuite) TestPushItem(){
	suite.taskQ.Push(suite.task)
	suite.Equal(1, suite.taskQ.Len())
	t := suite.taskQ.Fetch()
	suite.Equal(suite.task.URL, t.URL)
	suite.Equal(0, suite.taskQ.Len())
}

func (suite *TaskQueueTestSuite) TestFetchItem(){
	suite.taskQ.Push(suite.task)
	t := suite.taskQ.Fetch()
	suite.Equal(suite.task.URL, t.URL)
	suite.Equal(0, suite.taskQ.Len())
}

// TODO run Scheduler in HA
func (suite *TaskQueueTestSuite) TestConcurrentTaskQueueCreation(){

}

func (suite *TaskQueueTestSuite) TestSuiteTearDown(){
	for suite.taskQ.Len() > 0 {
		suite.taskQ.Fetch()
	}
	suite.taskQ.Close()
}

func TestTaskQueue(t *testing.T){
	suite.Run(t, new(TaskQueueTestSuite))
}
