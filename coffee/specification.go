package coffee

type Specification interface {
	IsSatisfied(coffee *Coffee) bool
}
