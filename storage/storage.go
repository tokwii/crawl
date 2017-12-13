package storage

/*import (
	"github.com/tokwii/crawl/config"
	"sync"
)

//
// Serialize
// Contains()
// SerialSiteMap() <- Priority!
// Maps Thread Safe?
// Use Thread Safe Ops
//
const LOCAL_STORAGE = "local"

const REMOTE_STORAGE = "remote"

type Storage interface {
	Contains() bool
	Add(string, map[string][]string)
	CreateSiteMap()
}

type CrawlerStore struct {
	Mode string
	store Storage
}

// SingletonStore
var store *CrawlerStore
var once sync.Once

func New() *CrawlerStore{
	once.Do(func(){
		queue = &TaskQueue{
			queue: make(chan Task, capacity),
		}
	})
	return queue
}

func build(){

	storageMode := config.Config.Storage.Mode
	switch storageMode{
	case LOCAL_STORAGE:

	case REMOTE_STORAGE:

	default:

	}
}
*/
