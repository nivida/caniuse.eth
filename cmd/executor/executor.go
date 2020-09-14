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
	run *[]Run
}

func New(methods []Run) (executor *Executor) {
	return new(Executor)
}

func (r *Executor) Start() {
	var config provider.Config;
	config.URL = "http://localhost:7545"
	config.Origin = "http://localhost/"

	p := provider.New(&config)

	log.Println("Started")

	// TODO: Use Array of "Runs"
	response := p.Send("eth_getBlockByNumber", "latest", true)

	log.Println(response)
}

func (r *Executor) Stop() {

}
