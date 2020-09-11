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
	config.URL = "ws://localhost:7545"
	config.Origin = "http://localhost/"

	provider, err := provider.New(&config)

	if err != nil {
		log.Fatal(err)
	}

	println("Started")

	provider.Close()
}

func (r *Runner) Stop() {

}
