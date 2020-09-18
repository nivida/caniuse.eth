package worker

import (
	"github.com/nivida/eth-rpc-tester/approver"
	"github.com/nivida/eth-rpc-tester/job"
	"github.com/nivida/eth-rpc-tester/provider"
)

type Worker struct {
	jobsPipeline    chan *job.Job
	resultsPipeline chan *job.Job
	provider        *provider.Provider
}

func New(p *provider.Provider, jobsPipeline chan *job.Job, resultsPipeline chan *job.Job) (worker *Worker) {
	w := new(Worker)
	w.jobsPipeline = jobsPipeline
	w.resultsPipeline = resultsPipeline
	w.provider = p

	return w
}

func (w *Worker) Start() {
	for s := range w.jobsPipeline {
		s.Response = w.provider.Send(s.Method, s.Params...)
		result := approver.Check(s.Response, s.Expected)

		if result == false {
			s.Successfull = false

			w.resultsPipeline <- s

			continue
		}

		s.Successfull = true
		w.resultsPipeline <- s
	}
}
