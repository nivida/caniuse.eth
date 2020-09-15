package worker

import (
	"github.com/nivida/eth-rpc-tester/approver"
	"github.com/nivida/eth-rpc-tester/provider"
)

type Job struct {
	Method      string
	Params      []interface{}
	Expectation approver.Expectation
	Response    interface{}
	Successfull bool
}

type Worker struct {
	jobs     chan *Job
	results  chan *Job
	provider *provider.Provider
}

func New(p *provider.Provider, jobs chan *Job, results chan *Job) (worker *Worker) {
	w := new(Worker)
	w.jobs = jobs
	w.results = results
	w.provider = p

	return w
}

func (w *Worker) Start() {
	for s := range w.jobs {
		s.Response = w.provider.Send(s.Method, s.Params...)
		result := approver.Check(s.Response, &s.Expectation)

		if result == false {
			s.Successfull = false

			w.results <- s

			return
		}

		s.Successfull = true
		w.results <- s
	}
}
