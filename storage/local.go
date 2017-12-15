package storage

import (
	"sync"
	"github.com/tokwii/crawl/common"
)

type LocalStorage struct {
	store map[string]map[string][]string

}

// Singleton Local Store
// Mutex for Thread Safe Ops on Map
var localStore *LocalStorage
var lsOnce sync.Once
var mutex = &sync.Mutex{}

func InitLocalStorage() *LocalStorage{
	lsOnce.Do(func(){
		localStore = &LocalStorage{
			store: make(map[string]map[string][]string),
		}
	})
	return localStore
}

func (s *LocalStorage) Contains(key string) (bool) {
	mutex.Lock()
	_, ok := s.store[key]
	mutex.Unlock()
	return ok
}

func (s *LocalStorage) Add (key string, value map[string][]string){
	mutex.Lock()
	s.store[key] = value
	mutex.Unlock()
}

func (s *LocalStorage) Get(key string) (map[string][]string, bool) {
	value, ok := s.store[key]
	if ok {
		return value, true
	}
	return nil, false
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
