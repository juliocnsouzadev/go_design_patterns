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
	cappuccinoSm := coffee.NewCoffee(coffee.Cappuccino, coffee.Small)
	cappuccinoMed := coffee.NewCoffee(coffee.Cappuccino, coffee.Medium)
	cappuccinoLar := coffee.NewCoffee(coffee.Cappuccino, coffee.Large)

	frenchPressSm := coffee.NewCoffee(coffee.FrenchPress, coffee.Small)
	frenchPressMed := coffee.NewCoffee(coffee.FrenchPress, coffee.Medium)
	frenchPressLar := coffee.NewCoffee(coffee.FrenchPress, coffee.Large)

	blackSm := coffee.NewCoffee(coffee.Black, coffee.Small)
	blackMed := coffee.NewCoffee(coffee.Black, coffee.Medium)
	blackLar := coffee.NewCoffee(coffee.Black, coffee.Large)

	coffees := coffee.AsList(
		cappuccinoSm,
		cappuccinoMed,
		cappuccinoLar,
		frenchPressSm,
		frenchPressMed,
		frenchPressLar,
		blackSm,
		blackMed,
		blackLar,
	)

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
	someMeal := meal.NewMealBuilder().
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
	fmt.Println("Built Meal:\n", someMeal.String())

	fmt.Println("\n========> Factory Pattern <========")
	breakfast := meal.NewMealHouseSpecial(meal.Breakfast)
	lunch := meal.NewMealHouseSpecial(meal.Lunch)
	dinner := meal.NewMealHouseSpecial(meal.Dinner)

	for _, houseSpecial := range []*meal.Meal{breakfast, lunch, dinner} {
		fmt.Println("Default Meal:\n", houseSpecial.String())
	}
}
