package worker

import (
	"log"

	"github.com/nivida/eth-rpc-tester/provider"
)

type Run struct {
	Method string
	Params []interface{}
}

type Worker struct {
	run []Run
	provider *provider.Provider
}

func New(runs []Run, config *provider.Config) (worker *Worker) {
	w := new(Worker)
	w.run = runs
	w.provider = provider.New(config)

	return w
}

func (w *Worker) Start(c chan []interface{}) {
	responses := make([]interface{}, len(w.run))

	for i, s := range w.run {
		responses[i] = w.provider.Send(s.Method, s.Params...)
	}
	
	log.Println(responses[0])

	c <- responses
}