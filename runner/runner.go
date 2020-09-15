package runner

import (
	"github.com/nivida/eth-rpc-tester/provider"
	"github.com/nivida/eth-rpc-tester/worker"
)

type Runner struct {
	SuccessCount     int
	FailureCount     int
	FailedCases      []*worker.Job
	SuccessfullCases []*worker.Job
	Results          chan<- interface{}
	Jobs             <-chan worker.Job
	Provider         *provider.Provider
}

func New(provider *provider.Provider) (runner *Runner) {
	r := new(Runner)
	r.Provider = provider

	return r
}

func (r *Runner) Start() {
	r.startWorkers(2, provider)
	r.passJobs()
	r.processTestResults()
}

func (r *Runner) passJobs() {
	tasks := r.getJobs()

	// Pass one task after another into the jobs channel for our workers
	for _, v := range tasks {
		r.jobs <- v
	}

	close(jobs)
}

func (r *Runner) startWorkers(amount int) {
	for w := 1; w <= amount; w++ {
		worker := worker.New(r.Provider, jobs, r.results)
		go worker.Start()
	}
}

func (r *Runner) processTestResults() {
	for i := 1; i <= len(tasks); i++ {
		result := <-r.results
		if result.Successfull == true {
			r.SuccessCount++
			append(r.SuccessfullCases, result)

			continue
		}

		r.FailureCount++
		append(r.FailedCases, result)
	}
}

// TODO: Implement buffered loader to push one after an other into the jobs pipeline (all at once takes a bit more memory on the same time)
/*
* loader := loader.New("testcases.json")
* taskChunk = <-loader.GetTasks()
*
 */
func (r *Runner) getJobs() *[]worker.Job {
	var tasks = make([]worker.Job, 4)
	tasks[0] = worker.Job{Method: "eth_getBlockByNumber", Params: []interface{}{"latest", true}}
	tasks[1] = worker.Job{Method: "eth_getBlockByNumber", Params: []interface{}{"latest", true}}
	tasks[2] = worker.Job{Method: "eth_getBlockByNumber", Params: []interface{}{"latest", true}}
	tasks[3] = worker.Job{Method: "eth_getBlockByNumber", Params: []interface{}{"latest", true}}

	return tasks
}
