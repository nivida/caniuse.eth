package main

import (
	"flag"
	"log"

	"github.com/nivida/eth-rpc-tester/loader"
	"github.com/nivida/eth-rpc-tester/provider"
	"github.com/nivida/eth-rpc-tester/runner"
)

func main() {
	workers := flag.Int("w", 4, "Amount of workers to start")
	flag.Parse()

	config := new(provider.Config)
	config.URL = "http://localhost:7545"
	config.Origin = "http://localhost/"
	p := provider.New(config)
	l := loader.New("pathToMyJson.json")

	r := runner.New(p, l)
	r.Start(*workers)

	log.Println("Success", r.SuccessCount)
	log.Println("Failure", r.FailureCount)

	// Delegate result to UI package
}
