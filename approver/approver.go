package approver

type Expectation struct {
	Value bool
}

func Check(response interface{}, expectation *Expectation) (bool, error) {
	return true, nil
	// DO ASSERTIONS BASED ON JOB HERE
}
