package storage

import (
	"sync"
	"github.com/tokwii/crawl/common"
)

type LocalStorage struct {
	// Sync Map??
	store map[string]map[string][]string

}
// Singleton Local Store

var localStore *LocalStorage
var lsOnce sync.Once

func InitLocalStorage() *LocalStorage{
	lsOnce.Do(func(){
		localStore = &LocalStorage{
			store: make(map[string]map[string][]string),
		}
	})
	return localStore
}

func (s *LocalStorage) Contains(key string) (bool) {
	_, ok := s.store[key]
	return ok
}

func (s *LocalStorage) Add (key string, value map[string][]string){
	s.store[key] = value
}

func (s *LocalStorage) CreateSiteMap () (common.Sitemap){
	var sitemap common.Sitemap
	var urls []common.URL
	for url, assets := range s.store {
		 u := common.URL{}
		 u.Loc = url
		 var js []common.Script
		 var styles []common.Style
		 var imgs []common.Image

		for aType, aValues  := range assets {
			switch aType {
			case "scripts":
				for _, val := range aValues{
					sc := common.Script{
						Loc: val,
					}
					js = append(js, sc)
				}
			case "styles":
				for _, val := range aValues{
					st := common.Style{
						Loc: val,
					}
					styles = append(styles, st)
				}
			case "images":
				for _, val := range aValues{
					im := common.Image{
						Loc: val,
					}
					imgs = append(imgs, im)
				}
			default:
				continue
			}
		}
		u.Image = imgs
		u.Scripts = js
		u.Styles = styles
		urls = append(urls, u)
	}
	sitemap.URLS = urls
	return sitemap
}
