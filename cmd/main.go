package main

import (
	"github.com/nivida/eth-rpc-tester/cmd/executor"
)

func main () {
	var runs = make([]executor.Run, 2)
	runs[0] = executor.Run{Method: "eth_getBlockByNumber", Params: []interface{} {"latest", true}}
	runs[1] = executor.Run{Method: "eth_getBlockByNumber", Params: []interface{} {"latest", true}}

	executor := executor.New(runs)

	executor.Start()
}