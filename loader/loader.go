package loader

import "github.com/nivida/eth-rpc-tester/worker"

// TODO: Move Job struct type to loader

// Loader struct
type Loader struct {
	path string
}

// New initiates a new Loader
func New(p string) (loader *Loader) {
	loader = new(Loader)
	loader.path = p

	return loader
}

// GetTasks returns all the defined tests from a JSON file
func (l *Loader) GetTasks() *[]worker.Job {
	var tasks = make([]worker.Job, 4)
	tasks[0] = worker.Job{Method: "eth_getBlockByNumber", Params: []interface{}{"latest", true}}
	tasks[1] = worker.Job{Method: "eth_getBlockByNumber", Params: []interface{}{"latest", true}}
	tasks[2] = worker.Job{Method: "eth_getBlockByNumber", Params: []interface{}{"latest", true}}
	tasks[3] = worker.Job{Method: "eth_getBlockByNumber", Params: []interface{}{"latest", true}}

	return &tasks
}
