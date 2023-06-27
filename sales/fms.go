package sales

const (
	FinalState State = "Final"
)

var (
	States map[State]Step = map[State]Step{
		LeadGenerationState: &LeadGeneration{},
		QualificationState:  &Qualification{},
		ProposalState:       &Proposal{},
	}
)

type Action func() bool
type State string

type Step interface {
	Given() State
	Execute() Action
	ThenGoTo() State
}

func ExecuteFMS() {
	currentState := LeadGenerationState
	for currentState != FinalState {
		step := States[currentState]
		action := step.Execute()
		if action() {
			currentState = step.ThenGoTo()
		}
	}
}
