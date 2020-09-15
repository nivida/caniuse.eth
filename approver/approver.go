package approver

type Expectation struct {
	Value interface{}
}

func check(response interface{}, expectation *Expectation) (bool, error) {
	// DO ASSERTIONS BASED ON JOB HERE
}
