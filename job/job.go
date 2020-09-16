package job

type Job struct {
	Method      string        `json:"method"`
	Params      []interface{} `json:"params"`
	Expected    interface{}   `json:"expected"`
	Response    interface{}
	Successfull bool
}

type Jobs struct {
	Jobs []Job `json:"tests"`
}
