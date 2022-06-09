package coffee

type SizeSpecification struct {
	Size Size
}

func (sizeSpecification SizeSpecification) IsSatisfied(coffee *Coffee) bool {
	return coffee.Size == sizeSpecification.Size
}

func NewSizeSpecification(size Size) SizeSpecification {
	return SizeSpecification{Size: size}
}
