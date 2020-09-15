package worker

import (
	"log"

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
	log.Println("WORKER CREATED")
	w := new(Worker)
	w.jobs = jobs
	w.results = results
	w.provider = p

	return w
}

func (w *Worker) Start() {
	for s := range w.jobs {
		s.Response = w.provider.Send(s.Method, s.Params...)
		log.Println(s.Response)
		//result, err = approver.Check(s.Response, s.Expectation)

		/*if err != nil {
			s.Successfull = false

			w.results <- s

			return
		}*/

		s.Successfull = true
		w.results <- s
	}
}
