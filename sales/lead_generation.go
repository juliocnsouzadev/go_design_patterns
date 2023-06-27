package sales

import "fmt"

const LeadGenerationState State = "Lead Generation"

type LeadGeneration struct {
}

func (l *LeadGeneration) Given() State {
	return LeadGenerationState
}

func (l *LeadGeneration) Execute() Action {
	return func() bool {
		fmt.Println(l.Given())
		return true
	}
}

func (l *LeadGeneration) ThenGoTo() State {
	return QualificationState
}
