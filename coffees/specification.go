package coffees

type Specification interface {
	IsSatisfied(coffee *coffee) bool
}
