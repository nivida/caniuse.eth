package worker

import (
	"github.com/nivida/eth-rpc-tester/provider"
	"github.com/nivida/eth-rpc-tester/approver"
)

type Job struct {
	Method string
	Params []interface{}
	ExpectedValue interface{}
}

type Worker struct {
	jobs <-chan Job
	results chan<- interface{}
	provider *provider.Provider
	approver *approver.Approver
}

func New(approver *Approver, config *provider.Config, jobs <-chan Job, results chan<- interface{}) (worker *Worker) {
	w := new(Worker)
	w.jobs = jobs
	w.results = results
	w.provider = provider.New(config)
	w.approver = approver

	return w
}

func (w *Worker) Start() {
	for s := range w.jobs {
		w.results <- w.provider.Send(s.Method, s.Params...)
	}
}
