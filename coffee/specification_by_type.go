package coffee

type TypeSpecification struct {
	Type Type
}

func (typeSpecification TypeSpecification) IsSatisfied(coffee *Coffee) bool {
	return coffee.Type == typeSpecification.Type
}

func NewTypeSpecification(type_ Type) TypeSpecification {
	return TypeSpecification{Type: type_}
}
