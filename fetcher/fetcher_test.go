package fetcher

import (
	"github.com/stretchr/testify/suite"
	"github.com/tokwii/crawl/storage"
	"github.com/tokwii/crawl/queue"
	"gopkg.in/h2non/gock.v1"
	"testing"
	"fmt"
)

type FetcherTestSuite struct {
	suite.Suite
	fetcher Fetcher
	htmlBody string
	taskQueue *queue.CrawlerQueue
	crawStore *storage.CrawlerStorage
}


func (suite *FetcherTestSuite) SetupSuite(){
	suite.taskQueue = queue.InitCrawlerQueue()
	suite.crawStore = storage.InitCrawlerStorage()
	suite.fetcher = Fetcher{}
	suite.fetcher.BaseUrl, _ = suite.fetcher.getBaseUrl("http://johndoe.com/article")
	suite.fetcher.EnableExternalLinks = false
	suite.htmlBody = `
	<html xmlns="http://www.w3.org/1999/xhtml" xml:lang="en" lang="en">
		<head>
			<link rel="stylesheet" type="text/css" href="http://johndoe.none/assets/css/prettify.css">
			<link rel="stylesheet" type="text/css" href="http://janendoe.none/assets/css/uglify.css">
			<link rel="stylesheet" type="text/css">
			<script type="text/javascript" src="http://johndoe.none/assets/js/awesome.js"></script>
		</head>
		<body>
			<div>
				<img src="http://johndoe.none/assets/images/zoro.png">
			</div>
			<div>
				<a href="http://janedoe.none/blog"></a>
				<a></a>
				<div>
					<a href="/books"></a>
					<a href="/books/favourite"></a>
				</div>
				<div>
					<a href="http://randomavatar.none/travel"></a>
					<a href="http://johndoe.none/sports"></a>
				</div>
			</div>
			<script type="text/javascript" src="http://akamai.net/johndeo/assets/js/unify.js"></script>
		</body>
	</html>
	`
}

func (suite *FetcherTestSuite) TestSuiteTearDown(){
	suite.taskQueue.Flush()
	suite.taskQueue.Close()
}

func (suite *FetcherTestSuite) TestGetBaseURL(){
	// TODO Encoded URLs
	rawUrls := make(map[string]string)
	rawUrls["http://google.com"] = "http://google.com/search?q=Monzo"
	rawUrls["https://gobyexample.com"] = "https://gobyexample.com/maps"
	rawUrls["emptyUrl"] = ""

	for  root, rawUrl := range rawUrls {
		res, ok := suite.fetcher.getBaseUrl(rawUrl)

		if  root == "emptyUrl" {
			suite.Equal(false, ok)
		}else{
			suite.Equal(root, res)
		}

	}
}

func (suite *FetcherTestSuite) TestBuildAbsoluteUrl(){
	relUrl := "/books"
	absUrl := suite.fetcher.buildAbsoluteUrl(relUrl)
	expectedUrl := fmt.Sprintf("%s%s", suite.fetcher.BaseUrl, relUrl)
	suite.Equal(expectedUrl, absUrl)
}

func (suite *FetcherTestSuite) TestGetTags(){

}

func (suite *FetcherTestSuite) TestCrawlerDisableExternalDomains(){
	//Cant Crawl Malformed HTML
	defer gock.Off()

	gock.New("http://johndoe.none").
		Reply(200).
		BodyString(suite.htmlBody)

	result, _ := FetchURL("http://johndoe.none", false, suite.taskQueue, suite.crawStore)

	links := []string {"http://johndoe.none/books", "http://johndoe.none/books/favourite", "http://johndoe.none/sports"}
	scripts := []string {"http://johndoe.none/assets/js/awesome.js", "http://akamai.net/johndeo/assets/js/unify.js"}
	styles := []string {"http://johndoe.none/assets/css/prettify.css","http://janendoe.none/assets/css/uglify.css"}

	suite.Contains(result.Images, "http://johndoe.none/assets/images/zoro.png")

	/*for _, link := range links {
		suite.Contains(result.Links, link)
	}*/

	for _, script := range scripts {
		suite.Contains(result.Scripts, script)
	}

	for _, style := range styles {
		suite.Contains(result.Styles, style)
	}
	suite.Equal(len(links), suite.taskQueue.Len())
	suite.taskQueue.Flush()
}


func (suite *FetcherTestSuite) TestCrawlerEnableExternalDomains(){
	//Cant Crawl Malformed HTML
	defer gock.Off()

	gock.New("http://johndoe.none").
		Reply(200).
		BodyString(suite.htmlBody)

	suite.taskQueue.Flush()

	FetchURL("http://johndoe.none", true, suite.taskQueue, suite.crawStore)

	links := []string {"http://janedoe.none/blog"}
	var resQ []string

	for  i := 0; i < suite.taskQueue.Len(); i++{
		resQ = append(resQ, suite.taskQueue.Fetch())
	}

	for _, link := range links {
		suite.Contains(resQ, link)
	}
	suite.taskQueue.Flush()
}

func (suite *FetcherTestSuite) TestCrawlerDomainsRedirection(){

	defer gock.Off()

	gock.New("http://johndoe.none").
		Get("/redirect").
		Reply(302).
		SetHeader("Location", "http://johndoe.none/resource")

	gock.New("http://johndoe.none").
		Get("/resource").
		Reply(200).
		BodyString(suite.htmlBody)

	result, _ := FetchURL("http://johndoe.none/redirect", false, suite.taskQueue, suite.crawStore)

	suite.Equal("http://johndoe.none/resource", result.Url)
	suite.Contains(result.Images, "http://johndoe.none/assets/images/zoro.png")
}

func TestFecter(t *testing.T){
	suite.Run(t, new(FetcherTestSuite))
}