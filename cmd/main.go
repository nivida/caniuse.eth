package main

import (
	"github.com/nivida/eth-rpc-tester/cmd/runner"
)

func main () {
	runner := runner.New()

	runner.Start()
}