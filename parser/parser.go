package main
/// Pass some page from the Internet

import (
	"fmt"
	"github.com/asaskevich/govalidator"
	"net/http"
	"strings"
	"io"
	//"io/ioutil"
	"golang.org/x/net/html"
	//"bytes"
)

// Change Name to Crawler?
type Parser struct {
	BaseUrl string
	EnableExternalLinks bool
}

type CrawlResult struct{
	URL string
	Styles []string
	Scripts []string
	Images []string
	Links []string
}
// Crawl External Links shoud be a flag be disable

func FetchURL(url string) (io.Reader, error) {
	// This should be a public method...Args -> url, maps[string][string]string
	valid := govalidator.IsURL(url)

	if !valid {
		return nil,fmt.Errorf("%v is Invalid", url)
	}

	resp, err := http.Get(url)

	if err != nil {
		return nil, fmt.Errorf("%v is Invalid", url)
	}

	//body, err := ioutil.ReadAll(resp.Body)

	if err != nil {
		return nil, fmt.Errorf("%v is Invalid", url)
	}

	//fmt.Println(string(body))
	return resp.Body, nil
}

func (p *Parser) buildUrlFromUri(uri string) (string){
	return fmt.Sprintf("%s%s", p.BaseUrl, uri)
}
// External links you dont want top scrap

func (p *Parser) getHrefTag(attributes []html.Attribute) (string, bool){

	for _, attr := range attributes {
		if attr.Key == "href" {
			var url string

			// Skip Discussions and Blogs organised by day
			if strings.HasPrefix(attr.Val, "/day") || strings.HasSuffix(attr.Val, "#disqus_thread"){
				continue
			}

			if !strings.HasPrefix(attr.Val, "http"){
				url = p.buildUrlFromUri(attr.Val)
			}else{
				url = attr.Val
			}

			if p.EnableExternalLinks {
				return url, true
			}

			if !p.EnableExternalLinks && strings.HasPrefix(url, p.BaseUrl){
				return url, true
			}
		}
	}

	return "" , false
}

func (p *Parser) getTag(attributes []html.Attribute, tagKey string) (string, bool){

	for _, attr := range attributes {
		if attr.Key == tagKey {
			return attr.Val, true
		}
	}
	return "", false
}

func (p *Parser) ParseHTML(htmlBody io.Reader) (CrawlResult, error){

	var styles, urls, imgs, js []string

	htmlDoc := html.NewTokenizer(htmlBody)
	// BFS
	for {
		tokenType := htmlDoc.Next()

		// Return Quickly In Case of Error
		if tokenType == html.ErrorToken {
			break
			//return CrawlResult{}, nil
		}

		if tokenType == html.StartTagToken {

			token := htmlDoc.Token()

			//names := htmlDoc.Raw()
			//fmt.Println("Raw Data")
			//fmt.Println(bytes.NewBuffer(names).String())

			switch {
			case token.Data == "a":
				//Links
				url, ok := p.getHrefTag(token.Attr)

				if ok {
					urls = append(urls, url)
					// Need to Check wheather it has already be crawled
					//fmt.Println(url)
					// Recrawl
				}

			case token.Data == "script":
				// Javascript
				script, ok := p.getTag(token.Attr, "src")

				if ok {
					js = append(js, script)
					//fmt.Println(script)
				}

			case token.Data == "img":
				// Images
				img, ok := p.getTag(token.Attr, "src")

				if ok {
					imgs = append(imgs, img)
					//fmt.Println(img)
				}

			case token.Data == "link":
				// Style sheets
				style, ok := p.getTag(token.Attr, "href")

				if ok {
					styles = append(styles, style)
					//fmt.Println(style)
				}

			default:
				continue

			}
		}
	}

	result := CrawlResult{
		URL: "http://tomblomfield.com",
		Styles: styles,
		Scripts: js,
		Images: imgs,
		Links: urls,

	}

	return result, nil
}

func CreateSiteMap(){

}

func main()  {
	body, err := FetchURL("http://tomblomfield.com/")
	if err != nil {

	}

	parser := &Parser{
		BaseUrl: "http://tomblomfield.com",
		EnableExternalLinks: false,
	}

	res, err := parser.ParseHTML(body)
	fmt.Println(res.Scripts)
}