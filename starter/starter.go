package starter

import (
	"log"

	"github.com/nivida/eth-rpc-tester/worker"
	"github.com/nivida/eth-rpc-tester/provider"
	"github.com/nivida/eth-rpc-tester/approver"
)

// Starts the worker pipeline and will later use the Loader to get all test cases
func Start(config *provider.Config) {
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

	approver := approver.New(results)

	// Start 2 workers
	for w := 1; w <= 2; w++ {
		worker := worker.New(approver, config, jobs, results)
		go worker.Start()
	}

	// Pass one task after another into the jobs channel for our workers
	for _, v := range tasks {
		jobs <- v
	}

	close(jobs)

	approver.check(tasks)
	
	// TODO: Display report in CLI and if possible with ease on a static HTML page
	// log.Println(approver.SuccessCount)
	// log.Println(approver.FailureCount)
	// log.Println(approver.FailedCases)
	// log.Println(approver.SuccessfullCases)
}
