package neurons

import (
	"fmt"
	"github.com/google/uuid"
)

type Neuron struct {
	Id      uuid.UUID
	In, Out []*Neuron
}

func NewNeuron() *Neuron {
	return &Neuron{Id: uuid.New()}
}

func (n *Neuron) Iter() []*Neuron {
	return []*Neuron{n}
}

func (n *Neuron) ConnectTo(other *Neuron) {
	n.Out = append(n.Out, other)
	other.In = append(other.In, n)
}

func (n *Neuron) showConnections() {
	for _, out := range n.Out {
		for _, in := range out.In {
			fmt.Printf("%s -> %s\n", in.Id, out.Id)
		}
	}
}
