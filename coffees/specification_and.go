package coffees

type AndSpecification struct {
	Specifications []Specification
}

func (andSpecification AndSpecification) IsSatisfied(coffee *coffee) bool {
	for _, specification := range andSpecification.Specifications {
		if !specification.IsSatisfied(coffee) {
			return false
		}
	}
	return true
}

func NewAndSpecification(specifications ...Specification) AndSpecification {
	return AndSpecification{Specifications: specifications}
}
