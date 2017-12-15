package queue

import (
	"github.com/stretchr/testify/suite"
	"testing"
)

type LocalQueueTestSuite struct {
	suite.Suite
	localQ *localQueue
	capacity int
	task string
}

func (suite *LocalQueueTestSuite) SetupSuite(){
	suite.capacity = 1000
	suite.localQ = initLocalQueue(suite.capacity)
	suite.task = "http://google.com"
}

func (suite *LocalQueueTestSuite) TestSingletonLocalQueue(){
	localQ2 := initLocalQueue(10)
	localQ2.Push(suite.task)
	suite.Equal(suite.localQ, localQ2)
	t2 := suite.localQ.Fetch()
	suite.Equal(suite.task, t2)
}

func (suite*LocalQueueTestSuite) TestPushItem(){
	suite.localQ.Push(suite.task)
	suite.Equal(1, suite.localQ.Len())
	t := suite.localQ.Fetch()
	suite.Equal(t, suite.task)
	suite.Equal(0, suite.localQ.Len())
}

func (suite *LocalQueueTestSuite) TestFetchItem(){
	suite.localQ.Push(suite.task)
	t := suite.localQ.Fetch()
	suite.Equal(suite.task, t)
	suite.Equal(0, suite.localQ.Len())
}

func (suite *LocalQueueTestSuite) TestSuiteTearDown(){
	for suite.localQ.Len() > 0 {
		suite.localQ.Fetch()
	}
	suite.localQ.Close()
}

func TestTaskQueue(t *testing.T){
	suite.Run(t, new(LocalQueueTestSuite))
}
