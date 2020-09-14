package main

import (
	"github.com/nivida/eth-rpc-tester/worker"
	"github.com/nivida/eth-rpc-tester/provider"
)

func main () {
	config :=  new(provider.Config);
	config.URL = "http://localhost:7545"
	config.Origin = "http://localhost/"


	var runs = make([]worker.Run, 2)
	runs[0] = worker.Run{Method: "eth_getBlockByNumber", Params: []interface{} {"latest", true}}
	runs[1] = worker.Run{Method: "eth_getBlockByNumber", Params: []interface{} {"latest", true}}

	worker := worker.New(runs, config)

	worker.Start()
}