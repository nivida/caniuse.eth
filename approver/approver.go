package approver

type Expectation struct {
	Value bool
}

func Check(response interface{}, expectation *Expectation) bool {
	return true
	// DO ASSERTIONS BASED ON JOB HERE
}
