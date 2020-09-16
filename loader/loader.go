package loader

import (
	"github.com/nivida/eth-rpc-tester/job"
)

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
func (l *Loader) GetTasks() *[]job.Job {
	var tasks = make([]job.Job, 4)
	tasks[0] = *job.New("eth_getBlockByNumber", []interface{}{"latest", true}, job.Expectation{Value: true})
	tasks[1] = *job.New("eth_getBlockByNumber", []interface{}{"latest", true}, job.Expectation{Value: true})
	tasks[2] = *job.New("eth_getBlockByNumber", []interface{}{"latest", true}, job.Expectation{Value: true})
	tasks[3] = *job.New("eth_getBlockByNumber", []interface{}{"latest", true}, job.Expectation{Value: true})

	return &tasks
}
