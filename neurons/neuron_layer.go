package neurons

type NeuronLayer struct {
	Neurons []Neuron
}

func (n *NeuronLayer) Iter() []*Neuron {
	result := make([]*Neuron, 0)
	for i := range n.Neurons {
		result = append(result, &n.Neurons[i])
	}
	return result
}

func NewNeuronLayer(count int) *NeuronLayer {
	neurons := make([]Neuron, count)
	for i := range neurons {
		neurons[i] = *NewNeuron()
	}
	return &NeuronLayer{neurons}
}
