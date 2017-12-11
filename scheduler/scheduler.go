package scheduler

import (
	"github.com/tokwii/crawl/queue"
)

// tasks being processed
//  Views of Tasks being processe map[string]Task
//
//
//
// Plugin Different Outputs
//

type Scheduler struct{
	Jobs map[string]queue.Task
	Results map[string]string
}

func Schedule(){

}



