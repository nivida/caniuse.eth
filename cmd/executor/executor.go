package executor

import (
	"log"

	"github.com/nivida/eth-rpc-tester/provider"
)

type Run struct {
	Method string
	Params []interface{}
}

type Executor struct {
	run []Run
}

func New(runs []Run) (executor *Executor) {
	e := new(Executor)
	e.run = runs

	return e
}

func (e *Executor) Start() {
	var config provider.Config;
	config.URL = "http://localhost:7545"
	config.Origin = "http://localhost/"

	p := provider.New(&config)

	log.Println("Started")

	for _, s := range e.run {
		log.Println(p.Send(s.Method, s.Params...))
	}

	response := p.Send("eth_getBlockByNumber", "latest", true)

	log.Println(response)
}

func (e *Executor) Stop() {

}
