package main

import (
	"github.com/nivida/eth-rpc-tester/provider"
	"github.com/nivida/eth-rpc-tester/starter"
)

func main () {
	config :=  new(provider.Config);
	config.URL = "http://localhost:7545"
	config.Origin = "http://localhost/"

	starter.Start(config)
}