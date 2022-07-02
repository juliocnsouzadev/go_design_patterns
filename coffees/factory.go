package coffees

import "fmt"

type HouseCoffee string

const (
	Cappuccino  HouseCoffee = "Cappuccino"
	FrenchPress HouseCoffee = "French Press"
	Black       HouseCoffee = "Black"
)

func NewCoffee(houseType HouseCoffee, size Size) *coffee {
	name := fmt.Sprintf("%s %s", houseType, size)
	switch houseType {
	case Cappuccino:
		return newCoffee(name, Sweet, size)
	case FrenchPress:
		return newCoffee(name, Light, size)
	case Black:
		return newCoffee(name, Bitter, size)
	default:
		panic("Unknown House coffees")
	}
}

func AsList(coffees ...*coffee) []*coffee {
	return coffees
}
