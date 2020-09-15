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
	p := provider.New(config)

	r := runner.New(p)
	r.Start(2) // Pass the amount of workers it should start

	log.Println("Success", r.SuccessCount)
	log.Println("Failure", r.FailureCount)

	// Delegate result to UI package
}
