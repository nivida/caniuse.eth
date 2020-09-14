package runner

import (
	"log"

	"github.com/nivida/eth-rpc-tester/provider"
)

type Runner struct {
	
}

func New() (runner *Runner) {
	return new(Runner)
}

func (r *Runner) Start() {
	var config provider.Config;
	config.URL = "http://localhost:7545"
	config.Origin = "http://localhost/"

	p := provider.New(&config)

	log.Println("Started")

	response := p.Send("eth_getBlockByNumber", "latest", true)

	log.Println(response)
}

func (r *Runner) Stop() {

}
