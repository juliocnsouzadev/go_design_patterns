package coffee

type Specification interface {
	IsSatisfied(coffee *coffee) bool
}
