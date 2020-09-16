package approver

import (
	"github.com/nivida/eth-rpc-tester/job"
)

// Check runs the actual test case with the given expectation of a Job
func Check(response interface{}, expectation *job.Expectation) bool {
	return true
	// DO ASSERTIONS BASED ON JOB HERE
}
