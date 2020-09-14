package starter

import (
	"log"

	"github.com/nivida/eth-rpc-tester/worker"
	"github.com/nivida/eth-rpc-tester/provider"
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

	// Start 2 workers
	for w := 1; w <= 2; w++ {
		worker := worker.New(config, jobs, results)
		go worker.Start()
	}

	// Pass one task after another into the jobs channel for our workers
	for _, v := range tasks {
		jobs <- v
	}

	close(jobs)

	// Print response as all of the tasks are finished
	for i := 1; i <= len(tasks); i++ {
		log.Println(<-results)
	} 
}