package runner

import (
	"log"

	"github.com/nivida/eth-rpc-tester/provider"
	"github.com/nivida/eth-rpc-tester/worker"
)

type Runner struct {
	SuccessCount     int
	FailureCount     int
	FailedCases      []*worker.Job
	SuccessfullCases []*worker.Job
	Results          chan *worker.Job
	Jobs             chan *worker.Job
	Provider         *provider.Provider
	tasks            *[]worker.Job
}

func New(p *provider.Provider) (runner *Runner) {
	r := new(Runner)
	r.Provider = p
	r.Jobs = make(chan *worker.Job)
	r.Results = make(chan *worker.Job)

	return r
}

func (r *Runner) Start() {
	log.Println("RUNNER STARTED")
	r.startWorkers(2)
	log.Println("WORKERS STARTED")
	r.passJobs()
	log.Println("JOBS PASSED")
	r.processTestResults()
	log.Println("TEST RESULTS PROCESSED")
}

func (r *Runner) passJobs() {
	r.tasks = r.getJobs()
	// Pass one task after another into the jobs channel for our workers
	for _, v := range *r.tasks {
		log.Println("JOB PASSED")
		r.Jobs <- &v
	}

	close(r.Jobs)
}

func (r *Runner) startWorkers(amount int) {
	for w := 1; w <= amount; w++ {
		worker := worker.New(r.Provider, r.Jobs, r.Results)
		go worker.Start()
	}
}

func (r *Runner) processTestResults() {
	for i := 1; i <= len(*r.tasks); i++ {
		result := <-r.Results
		if result.Successfull == true {
			r.SuccessCount++
			r.SuccessfullCases = append(r.SuccessfullCases, result)
			log.Println("SUCCESS")

			continue
		}

		log.Println("ERROR")
		r.FailureCount++
		r.FailedCases = append(r.FailedCases, result)
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

	return &tasks
}
