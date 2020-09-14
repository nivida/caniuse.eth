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

func (w *Worker) Start() {
	for _, s := range w.run {
		log.Println(w.provider.Send(s.Method, s.Params...))
	}
}

func (w *Worker) Stop() {

}
