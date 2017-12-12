package common

import (
	"github.com/stretchr/testify/suite"
	"testing"
)

type CommonTestSuite struct {
	suite.Suite
}

func (suite *CommonTestSuite) SetupSuite(){
}


func (suite *CommonTestSuite) TestSuiteTearDown(){

}

func TestCommon(t *testing.T){
	suite.Run(t, new(CommonTestSuite))
}

