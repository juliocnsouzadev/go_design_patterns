package main

import (
	"fmt"
	"go_design_patterns/coffee"
	"go_design_patterns/meal"
	"go_design_patterns/transactions"
)

func main() {
	fmt.Println("** Design Patterns in Go **")

	fmt.Println("\n========> Decorator Pattern <========")
	transaction := transactions.BankTransction
	transaction = transactions.SmsDecorator(transaction)
	transaction = transactions.EmailDecorator(transaction)
	transaction("salary", 1000, true)
	transaction("rent", 500, false)

	fmt.Println("\n========> Specification Pattern (Open and Closed Principle) <========")
	cappuccinoSm := coffee.NewCoffee("Cappuccino Small", coffee.Sweet, coffee.Small)
	cappuccinoMed := coffee.NewCoffee("Cappuccino Medium", coffee.Sweet, coffee.Medium)
	cappuccinoLar := coffee.NewCoffee("Cappuccino Large", coffee.Sweet, coffee.Large)

	frenchPressSm := coffee.NewCoffee("French Press Small", coffee.Light, coffee.Small)
	frenchPressMed := coffee.NewCoffee("French Press Medium", coffee.Light, coffee.Medium)
	frenchPressLar := coffee.NewCoffee("French Press Large", coffee.Light, coffee.Large)

	blackSm := coffee.NewCoffee("Black Small", coffee.Bitter, coffee.Small)
	blackMed := coffee.NewCoffee("Black Press Medium", coffee.Bitter, coffee.Medium)
	blackLar := coffee.NewCoffee("Black Press Large", coffee.Bitter, coffee.Large)

	coffees := []*coffee.Coffee{cappuccinoSm, cappuccinoMed, cappuccinoLar, frenchPressSm, frenchPressMed, frenchPressLar, blackSm, blackMed, blackLar}

	coffeeFilter := coffee.CoffeeFilter{}

	bitterCoffees := coffeeFilter.Filter(coffees, coffee.NewTypeSpecification(coffee.Bitter))

	fmt.Println("\nBitter Coffees:")
	for _, coffee := range bitterCoffees {
		fmt.Println(coffee.Name)
	}

	sweetAndSmallCoffees := coffeeFilter.Filter(
		coffees,
		coffee.NewAndSpecification(
			coffee.NewTypeSpecification(coffee.Sweet), coffee.NewSizeSpecification(coffee.Small),
		),
	)
	fmt.Println("\nSweet and Small Coffees:")
	for _, coffee := range sweetAndSmallCoffees {
		fmt.Println(coffee.Name)
	}

	fmt.Println("\n========> Builder Pattern <========")
	meal := meal.NewMealBuilder().
		Named("Super Mega Blaster Combo").
		WithMainDish().
		Named("Special Steak BBQ").
		WithDescription("Extra Large Juice Steak covered in BBQ sauce").
		WithSide().
		Named("Garlic Bread").
		WithDescription("4 delicious slices of garlic bread covered in cheese").
		WithDrink().
		Named("Premium Weiss Beer").
		WithSize(meal.Large).
		Build()
	fmt.Println("Built Meal:\n", meal.String())
}
