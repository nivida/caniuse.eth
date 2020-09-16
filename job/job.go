package job

import "github.com/ybbus/jsonrpc"

type Job struct {
	Comment     string        `json:"_comment"`
	Method      string        `json:"method"`
	Params      []interface{} `json:"params"`
	Expected    string        `json:"expected"`
	Response    *jsonrpc.RPCResponse
	Successfull bool
}

type Jobs struct {
	Jobs []Job `json:"tests"`
}
