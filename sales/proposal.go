package sales

import "fmt"

const ProposalState State = "Proposal"

type Proposal struct {
}

func (p *Proposal) Given() State {
	return ProposalState
}

func (p *Proposal) Execute() Action {
	return func() bool {
		fmt.Println(p.Given())
		return true
	}
}

func (p *Proposal) ThenGoTo() State {
	return FinalState
}
