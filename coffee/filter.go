package coffee

type CoffeeFilter struct{}

func (coffeeFilter *CoffeeFilter) Filter(coffees []*coffee, specification ...Specification) []*coffee {
	filteredCoffees := []*coffee{}
	for _, coffee := range coffees {
		for _, specification := range specification {
			if specification.IsSatisfied(coffee) {
				filteredCoffees = append(filteredCoffees, coffee)
			}
		}
	}
	return filteredCoffees
}
