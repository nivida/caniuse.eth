package main

import (
	"github.com/nivida/eth-rpc-tester/provider"
	"github.com/nivida/eth-rpc-tester/runner"
)

func main () {
	config :=  new(provider.Config);
	config.URL = "http://localhost:7545"
	config.Origin = "http://localhost/"
	
	r := new(Runner)
	r.Start(config)
}
