package storage

import (
	"github.com/tokwii/crawl/config"
	"github.com/tokwii/crawl/common"
	"sync"
	"fmt"
	"os"
)

const LOCAL_STORAGE = "local"

const REMOTE_STORAGE = "remote"

type Storage interface {
	Contains(string) bool
	Add(string, map[string][]string)
	CreateSiteMap() common.Sitemap
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

func (s *CrawlerStorage) CreateSiteMap () (common.Sitemap){
	return s.store.CreateSiteMap()
}

func build() *CrawlerStorage{

	var c config.Config
	err := c.Load("config/settings.toml")
	if err != nil {
  		fmt.Errorf("error: %v\n", err)
		os.Exit(-1)
	}

	cs := CrawlerStorage{}

	storageMode := c.Storage.Mode

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
