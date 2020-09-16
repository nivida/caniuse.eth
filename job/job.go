package job

type Expectation struct {
	Value bool
}

type Job struct {
	Method      string
	Params      []interface{}
	Expectation Expectation
	Response    interface{}
	Successfull bool
}

func New(method string, params []interface{}, expectation Expectation) (job *Job) {
	job = new(Job)
	job.Method = method
	job.Params = params
	job.Expectation = expectation

	return job
}
