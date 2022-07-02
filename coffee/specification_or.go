package coffee

type OrSpecification struct {
	Specifications []Specification
}

func (orSpecification OrSpecification) IsSatisfied(coffee *coffee) bool {
	for _, specification := range orSpecification.Specifications {
		if specification.IsSatisfied(coffee) {
			return true
		}
	}
	return false
}

func NewOrSpecification(specifications ...Specification) OrSpecification {
	return OrSpecification{Specifications: specifications}
}
