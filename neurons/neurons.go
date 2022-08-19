package neurons

type NeuronInterface interface {
	Iter() []*Neuron
}

func Connect(left, right NeuronInterface) {
	for _, l := range left.Iter() {
		for _, r := range right.Iter() {
			l.ConnectTo(r)
		}
	}
}

func ShowConnections(items ...NeuronInterface) {
	for _, item := range items {
		for _, neuron := range item.Iter() {
			neuron.showConnections()
		}
	}
}
