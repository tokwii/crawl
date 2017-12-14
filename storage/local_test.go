package storage

import (
	"github.com/stretchr/testify/suite"
	"github.com/tokwii/crawl/common"
	"encoding/xml"
	"testing"
	//"fmt"
)

type LocalStorageTestSuite struct {
	suite.Suite
	ls *LocalStorage
	mockData map[string][]string
}

func (suite *LocalStorageTestSuite) SetupSuite(){
	suite.ls = InitLocalStorage()
	suite.mockData = make(map[string][]string)
	styles := []string {"http://styles.none/stlye.css", "http://styles1.none/stlye.css"}
	js := []string {"http://scripts.none/js.js", "http://scripts.none/js.js"}
	imgs := []string {"http://images.none/img.gif", "http://images.none/img.png"}
	suite.mockData["styles"] = styles
	suite.mockData["scripts"] = js
	suite.mockData["images"] = imgs
	suite.ls.Add("http://janedoe.none", suite.mockData)
	suite.ls.Add("http://johndoe.none", suite.mockData)
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
	xmlByte, _ := xml.MarshalIndent(suite.ls.CreateSiteMap() ,"  ", "    ")

	/*if err != nil {
		fmt.Printf("error: %v\n", err)
	}*/

	sitemap := common.Sitemap{}
	xml.Unmarshal(xmlByte, &sitemap)
	/*if err != nil {
		fmt.Printf("error: %v\n", err)
	}*/
	suite.Equal(len(suite.mockData), len(sitemap.URLS))
	//suite.Equal(suite.mockData["http://janedoe.none"], sitemap.URLS[0].Image)
	//xmlstring := []byte(xml.Header + string(output))
	//os.Stdout.Write(xmlstring)

}

func TestLocalStorage(t *testing.T){
	suite.Run(t, new(LocalStorageTestSuite))
}