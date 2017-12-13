package storage

import (
	"github.com/stretchr/testify/suite"
	"encoding/xml"
	"testing"
	"os"
	"fmt"
)

type LocalStorageTestSuite struct {
	suite.Suite
	ls *LocalStorage
}

func (suite *LocalStorageTestSuite) SetupSuite(){
	suite.ls = InitLocalStorage()
	mockData := make(map[string][]string)
	styles := []string {"http://styles.none/stlye.css", "http://styles1.none/stlye.css"}
	js := []string {"http://scripts.none/js.js", "http://scripts.none/js.js"}
	imgs := []string {"http://images.none/img.gif", "http://images.none/img.png"}
	mockData["styles"] = styles
	mockData["scripts"] = js
	mockData["images"] = imgs
	suite.ls.Add("http://janedoe.none", mockData)
	suite.ls.Add("http://johndoe.none", mockData)
}

func (suite *LocalStorageTestSuite) TestSingletonLocalStorage(){
	localStore := InitLocalStorage()
	suite.Equal(localStore, suite.ls)
	suite.Equal(localStore.Contains("http://janedoe.none"), suite.ls.Contains("http://janedoe.none"))
}

func (suite *LocalStorageTestSuite) TestWhetherItemExistsItem(){
	suite.Equal(true, suite.ls.Contains("http://janedoe.none"))
	suite.Equal(false, suite.ls.Contains("http://jamesdoe.none"))
}

func (suite *LocalStorageTestSuite) TestAddItem(){
	mockData := make(map[string][]string)
	imgs := []string {"http://images.none/img.gif", "http://images.none/img.png"}
	mockData["styles"] = make([]string, 0)
	mockData["scripts"] = make([]string, 0)
	mockData["images"] = imgs
	suite.ls.Add("http://cliqr.none", mockData)
	suite.Equal(true, suite.ls.Contains("http://cliqr.none"))
	suite.Empty(suite.ls.store["http://cliqr.none"]["scripts"])
	suite.Equal(suite.ls.store["http://cliqr.none"]["images"], imgs)
}

func (suite *LocalStorageTestSuite) TestCreateSiteMap(){

	output, err := xml.MarshalIndent(suite.ls.CreateSiteMap() ,"  ", "    ")
	if err != nil {
		fmt.Printf("error: %v\n", err)
	}
	os.Stdout.Write(output)

}

func TestLocalStorage(t *testing.T){
	suite.Run(t, new(LocalStorageTestSuite))
}