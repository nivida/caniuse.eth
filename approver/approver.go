package approver

import (
	"github.com/nivida/eth-rpc-tester/worker"
)

type Approver struct {
	SuccessCount int
	FailureCount int
	FailedCases []*worker.Job
	SuccessfullCases []*worker.Job
	results: chan<- interface{}
}

func New(results chan<- interface{}) (approver *Approver) {
	return new(Approver)
}

func (a *Approver) check(jobs *[]worker.Job) {
	for i := 1; i <= len(jobs); i++ {
		// Check <-results === jobs[i].ExpectedValue
	} 
}
