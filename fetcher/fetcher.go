package fetcher

import (
	"fmt"
	"github.com/asaskevich/govalidator"
	"net/url"
	"net/http"
	"strings"
	"io"
	"golang.org/x/net/html"
	"github.com/tokwii/crawl/queue"
	"github.com/tokwii/crawl/storage"
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
}

func FetchURL(rawUrl string, fetchExternalDomain bool, taskQueue *queue.CrawlerQueue, crawlerStore *storage.CrawlerStorage) (Result, error) {

	valid := govalidator.IsURL(rawUrl)

	if !valid {
		return Result{}, fmt.Errorf("%v is Invalid", rawUrl)
	}

	response, err := http.Get(rawUrl)

	if err != nil {
		return Result{}, fmt.Errorf("%v is Invalid", rawUrl)
	}

	bodyByteStream := response.Body
	defer bodyByteStream.Close()

	// Follow Redirects HTTP 301/302
	requestUrl := response.Request.URL.String()

	// Removes Poorly formatted url(s) -> Hash
	parsedUrl, _ := url.Parse(requestUrl)
	cleanReqUrl := fmt.Sprintf("%s://%s%s", parsedUrl.Scheme, parsedUrl.Host, parsedUrl.Path)

	ok := crawlerStore.Contains(cleanReqUrl)

	// Different Urls(Aliases) that redirect to the same url
	if ok {
		return Result{}, fmt.Errorf("Alias for %v already crawled", rawUrl)
	}

	f := Fetcher{}

	baseUrl, ok := f.getBaseUrl(requestUrl)

	if !ok {
		return Result{}, fmt.Errorf("Error Retrieving Base Url for %v ", rawUrl)
	}

	f.Url = requestUrl
	f.BaseUrl = baseUrl
	f.EnableExternalLinks = fetchExternalDomain

	result, err := f.crawl(bodyByteStream, taskQueue)

	if err != nil {
		return Result{}, fmt.Errorf("Error Parsing page for %v ", rawUrl)
	}

	return result, nil
}

func (f *Fetcher) buildAbsoluteUrl(uri string) (string){
	return fmt.Sprintf("%s%s", f.BaseUrl, uri)
}

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

func (f *Fetcher) crawl(htmlBody io.Reader, taskQueue *queue.CrawlerQueue) (Result, error){

	var styles, imgs, js []string

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
					taskQueue.Push(url)
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
	}

	return result, nil
}