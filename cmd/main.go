package main

import (
	"log"

	"github.com/nivida/eth-rpc-tester/provider"
	"github.com/nivida/eth-rpc-tester/runner"
)

func main() {
	config := new(provider.Config)
	config.URL = "http://localhost:7545"
	config.Origin = "http://localhost/"

	r := runner.New()
	r.Start(config)

	log.Println(r.SuccessCount)
	log.Println(r.FailureCount)
	log.Println(r.FailedCases)
	log.Println(r.SuccessfullCases)

	// Delegate result to UI package
}
