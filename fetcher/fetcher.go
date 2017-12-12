package fetcher

import (
	"fmt"
	"errors"
	"github.com/asaskevich/govalidator"
	"net/url"
	"net/http"
	"strings"
	"io"
	"golang.org/x/net/html"
	"github.com/tokwii/crawl/queue"
)

type Fetcher struct {
	BaseUrl string
	Url string
	EnableExternalLinks bool
}

type Result struct{
	Url string
	Styles []string
	Scripts []string
	Images []string
	Links []string
}
// Crawl External Links shoud be a flag be disable

func FetchURL(url string, fetchExternalDomain bool, taskQueue *queue.TaskQueue) (Result, error) {
	// This should be a public method...Args -> url, maps[string][string]string
	valid := govalidator.IsURL(url)

	if !valid {
		return Result{}, fmt.Errorf("%v is Invalid", url)
	}

	response, err := http.Get(url)

	if err != nil {
		return Result{}, fmt.Errorf("%v is Invalid", url)
	}

	bodyByteStream := response.Body
	// Close http Connection after return
	defer bodyByteStream.Close()

	f := &Fetcher{}
	baseUrl, ok := f.getBaseUrl(url)

	if !ok {
		return Result{}, errors.New("Error Retrieving Base Url")
	}

	f.Url = url
	f.BaseUrl = baseUrl
	f.EnableExternalLinks = fetchExternalDomain

	result, err := f.crawl(bodyByteStream, taskQueue)

	if err != nil {
		return Result{}, errors.New("Error Retrieving Crawling")
	}

	return result, nil
}

func (f *Fetcher) buildAbsoluteUrl(uri string) (string){
	return fmt.Sprintf("%s%s", f.BaseUrl, uri)
}

// External links you dont want top scrap

func (f *Fetcher) getBaseUrl(rawUrl string) (string, bool){

	url, err := url.Parse(rawUrl)

	if err != nil {
		return "", false
	}

	if url.Scheme == "" || url.Host == "" {
		return "", false
	}

	return fmt.Sprintf("%s://%s", url.Scheme, url.Host), true
}

func (f *Fetcher) getHrefTag(attributes []html.Attribute) (string, bool){

	for _, attr := range attributes {
		if attr.Key == "href" {
			var url string

			// Skip Discussions and Blogs organised by day
			if strings.HasPrefix(attr.Val, "/day") || strings.HasSuffix(attr.Val, "#disqus_thread"){
				continue
			}

			if !strings.HasPrefix(attr.Val, "http"){
				url = f.buildAbsoluteUrl(attr.Val)
			}else{
				url = attr.Val
			}

			if f.EnableExternalLinks {
				return url, true
			}

			if !f.EnableExternalLinks && strings.HasPrefix(url, f.BaseUrl){
				return url, true
			}
		}
	}

	return "" , false
}

func (f *Fetcher) getTag(attributes []html.Attribute, tagKey string) (string, bool){

	for _, attr := range attributes {
		if attr.Key == tagKey {
			return attr.Val, true
		}
	}
	return "", false
}

func (f *Fetcher) crawl(htmlBody io.Reader, taskQueue *queue.TaskQueue) (Result, error){

	var styles, urls, imgs, js []string

	htmlDoc := html.NewTokenizer(htmlBody)
	// BFS
	for {
		tokenType := htmlDoc.Next()

		if tokenType == html.ErrorToken {
			if htmlDoc.Err().Error() == "EOF"{
				break
			} else{
				return Result{}, htmlDoc.Err()
			}
		}

		if tokenType == html.StartTagToken {

			token := htmlDoc.Token()
			switch {
			case token.Data == "a":
				//Links
				url, ok := f.getHrefTag(token.Attr)

				if ok {
					urls = append(urls, url)
					task := queue.Task{
						URL: url,
					}
					taskQueue.Push(task)
					//taskQueue.Flush()
					// Need to Check wheather it has already be crawled. Will be Checked by Scheduler
					// Recrawl
				}

			case token.Data == "script":
				// Javascript
				script, ok := f.getTag(token.Attr, "src")

				if ok {
					js = append(js, script)
				}

			case token.Data == "img":
				// Images
				img, ok := f.getTag(token.Attr, "src")

				if ok {
					imgs = append(imgs, img)
				}

			case token.Data == "link":
				// Style sheets
				style, ok := f.getTag(token.Attr, "href")

				if ok {
					styles = append(styles, style)
				}

			default:
				continue

			}
		}
	}

	result := Result{
		Url: f.Url,
		Styles: styles,
		Scripts: js,
		Images: imgs,
		Links: urls,
	}
	return result, nil
}
func CreateSiteMap(){

}
