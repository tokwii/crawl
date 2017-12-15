package storage

import (
	"github.com/tokwii/crawl/config"
	"github.com/tokwii/crawl/common"
	"sync"
)

const LOCAL_STORAGE = "local"

const REMOTE_STORAGE = "remote"

type Storage interface {
	Contains(string) bool
	Add(string, map[string][]string)
	CreateSiteMap() common.Sitemap
	Get(string) (map[string][]string, bool)
}

type CrawlerStorage struct {
	Mode string
	store Storage
}

// SingletonStore
var crawlerStore *CrawlerStorage
var once sync.Once

func InitCrawlerStorage() *CrawlerStorage{
	once.Do(func(){
		crawlerStore = build()
	})
	return crawlerStore
}

func (s *CrawlerStorage) Contains(key string) (bool) {
	ok := s.store.Contains(key)
	return ok
}

func (s *CrawlerStorage) Add (key string, value map[string][]string){
	s.store.Add(key, value)
}

func (s *CrawlerStorage) Get(key string) (map[string][]string, bool){
	res, ok := s.store.Get(key)
	return res, ok
}

func (s *CrawlerStorage) CreateSiteMap () (common.Sitemap){
	return s.store.CreateSiteMap()
}

func build() *CrawlerStorage{

	cs := CrawlerStorage{}
	storageMode := config.Conf.Storage.Mode

	switch storageMode {
	case LOCAL_STORAGE:
		ls := InitLocalStorage()
		cs.Mode = storageMode
		cs.store = ls
	case REMOTE_STORAGE:
		//Remote
	default:
		// Defaults to Local Storage
		ls := InitLocalStorage()
		cs.Mode = storageMode
		cs.store = ls
	}
	return  &cs
}
