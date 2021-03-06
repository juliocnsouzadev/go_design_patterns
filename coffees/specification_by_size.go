package coffees

type SizeSpecification struct {
	Size Size
}

func (sizeSpecification SizeSpecification) IsSatisfied(coffee *coffee) bool {
	return coffee.Size == sizeSpecification.Size
}

func NewSizeSpecification(size Size) SizeSpecification {
	return SizeSpecification{Size: size}
}
