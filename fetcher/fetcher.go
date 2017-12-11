package fetcher

import(
	"net/url"
	"github.com/tokwii/crawl/queue"
)

// Fetcher
// Append links that were encounter
type Fetcher struct {
	task queue.Task
	result map[string]string
}

func (f *Fetcher) fetch


func (f *Fetcher) Run(t *queue.Task) map[string]string{
	q =: queue.TaskQueue(10)
}

