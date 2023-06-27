package sales

import "fmt"

const QualificationState State = "Qualification"

type Qualification struct {
}

func (q *Qualification) Given() State {
	return QualificationState
}

func (q *Qualification) Execute() Action {
	return func() bool {
		fmt.Println(q.Given())
		return true
	}
}

func (q *Qualification) ThenGoTo() State {
	return ProposalState
}
