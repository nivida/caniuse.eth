package approver

import (
	"log"

	"github.com/ybbus/jsonrpc"
)

// Check runs the actual test case with the given expectation of a Job
func Check(response *jsonrpc.RPCResponse, expected string) bool {
	returnValue, err := response.GetString()

	if err != nil {
		log.Fatal(err)
	}

	return returnValue == expected
}

/*
// State change executes for example a simple eth_sendTransaction etc.
func stateChange() {

}
*/
