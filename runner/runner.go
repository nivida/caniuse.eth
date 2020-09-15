package runner

import (
	"log"

	"github.com/nivida/eth-rpc-tester/worker"
	"github.com/nivida/eth-rpc-tester/provider"
	"github.com/nivida/eth-rpc-tester/approver"
)

type Runner struct {
	SuccessCount int
	FailureCount int
	FailedCases []*worker.Job
	SuccessfullCases []*worker.Job
	results: chan<- interface{}
}

func New() (runner *Runner) {
	return new(Runner)
}

// Starts the worker pipeline and will later use the Loader to get all test cases
func (r *Runner) Start(config *provider.Config) {
	// TODO: Implement buffered loader to push one after an other into the jobs pipeline (all at once takes a bit more memory on the same time)
	/*
	* loader := loader.New("testcases.json")
	* taskChunk = <-loader.GetTasks()
	*
	*/
	var tasks = make([]worker.Job, 4)
	tasks[0] = worker.Job{Method: "eth_getBlockByNumber", Params: []interface{} {"latest", true}}
	tasks[1] = worker.Job{Method: "eth_getBlockByNumber", Params: []interface{} {"latest", true}}
	tasks[2] = worker.Job{Method: "eth_getBlockByNumber", Params: []interface{} {"latest", true}}
	tasks[3] = worker.Job{Method: "eth_getBlockByNumber", Params: []interface{} {"latest", true}}

	jobs := make(chan worker.Job, 2)
	results := make(chan interface{}, 2)

	// Start 2 workers
	for w := 1; w <= 2; w++ {
		worker := worker.New(config, jobs, r.results)
		go worker.Start()
	}

	// Pass one task after another into the jobs channel for our workers
	for _, v := range tasks {
		jobs <- v
	}

	close(jobs)

	for i := 1; i <= len(tasks); i++ {
		result := approver.check(<-r.results)
		if result.Successfull == true {
			r.SuccessCount++
			append(r.SuccessfullCases, result)

			continue
		}

		r.FailureCount++
		append(r.FailedCases, result)
	} 

	// log.Println(approver.SuccessCount)
	// log.Println(approver.FailureCount)
	// log.Println(approver.FailedCases)
	// log.Println(approver.SuccessfullCases)
}
