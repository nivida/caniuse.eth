package runner

import (
	"github.com/nivida/eth-rpc-tester/job"
	"github.com/nivida/eth-rpc-tester/loader"
	"github.com/nivida/eth-rpc-tester/provider"
	"github.com/nivida/eth-rpc-tester/worker"
)

// Runner struct
type Runner struct {
	SuccessCount     int
	FailureCount     int
	FailedCases      []*job.Job
	SuccessfullCases []*job.Job
	Results          chan *job.Job
	Jobs             chan *job.Job
	Provider         *provider.Provider
	Loader           *loader.Loader
	tasks            []job.Job
}

// New initiates the Runner
func New(p *provider.Provider, l *loader.Loader) (runner *Runner) {
	r := new(Runner)
	r.Provider = p
	r.Loader = l

	return r
}

// Start the whole test run
func (r *Runner) Start(amount int) {
	r.tasks = r.Loader.GetTasks()
	r.Jobs = make(chan *job.Job, len(r.tasks))
	r.Results = make(chan *job.Job, len(r.tasks))

	r.startWorkers(amount)
	r.passJobs()
	r.processTestResults()
}

// Pass one task after another into the jobs channel for our workers
func (r *Runner) passJobs() {
	/**
	* Learning comment:
	* range can't be used cause the for-each loop is reusing the memory location which means
	* the used value for our workers does update and it shouldn't.
	*
	* Another fix for this would be to pass it by value instead of reference. Let me see.
	**/
	for i := 0; i < len(r.tasks); i++ {
		r.Jobs <- &r.tasks[i]
	}

	close(r.Jobs)
}

// Start the correct amount of workers
func (r *Runner) startWorkers(amount int) {
	for w := 1; w <= amount; w++ {
		worker := worker.New(r.Provider, r.Jobs, r.Results)
		go worker.Start()
	}
}

// Process test results
func (r *Runner) processTestResults() {
	for i := 1; i <= len(r.tasks); i++ {
		result := <-r.Results
		if result.Successfull == true {
			r.SuccessCount++
			r.SuccessfullCases = append(r.SuccessfullCases, result)

			continue
		}

		r.FailureCount++
		r.FailedCases = append(r.FailedCases, result)
	}
}
