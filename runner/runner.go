package runner

import (
	"github.com/nivida/eth-rpc-tester/loader"
	"github.com/nivida/eth-rpc-tester/provider"
	"github.com/nivida/eth-rpc-tester/worker"
)

// Runner struct
type Runner struct {
	SuccessCount     int
	FailureCount     int
	FailedCases      []*worker.Job
	SuccessfullCases []*worker.Job
	Results          chan *worker.Job
	Jobs             chan *worker.Job
	Provider         *provider.Provider
	Loader           *loader.Loader
	tasks            *[]worker.Job
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
	r.Jobs = make(chan *worker.Job, len(*r.tasks))
	r.Results = make(chan *worker.Job, len(*r.tasks))

	r.startWorkers(amount)
	r.passJobs()
	r.processTestResults()
}

// Pass one task after another into the jobs channel for our workers
func (r *Runner) passJobs() {
	for _, v := range *r.tasks {
		r.Jobs <- &v
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
	for i := 1; i <= len(*r.tasks); i++ {
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
