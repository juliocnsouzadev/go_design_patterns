package main

import (
	"fmt"
	"go_design_patterns/coffees"
	"go_design_patterns/cs"
	"go_design_patterns/meals"
	"go_design_patterns/neurons"
	"go_design_patterns/transactions"
	"go_design_patterns/wallet"
	"log"
)

func main() {
	fmt.Println("** Design Patterns in Go **")

	fmt.Println("\n========> Decorator Pattern <========")
	transaction := transactions.BankTransaction
	transaction = transactions.SmsDecorator(transaction)
	transaction = transactions.EmailDecorator(transaction)
	transaction(transactions.Transaction{Desc: "salary", Value: 1000, In: true})
	transaction(transactions.Transaction{Desc: "rent", Value: 500})

	fmt.Println("\n========> Specification Pattern (Open and Closed Principle) <========")
	cappuccinoSm := coffees.NewCoffee(coffees.Cappuccino, coffees.Small)
	cappuccinoMed := coffees.NewCoffee(coffees.Cappuccino, coffees.Medium)
	cappuccinoLar := coffees.NewCoffee(coffees.Cappuccino, coffees.Large)

	frenchPressSm := coffees.NewCoffee(coffees.FrenchPress, coffees.Small)
	frenchPressMed := coffees.NewCoffee(coffees.FrenchPress, coffees.Medium)
	frenchPressLar := coffees.NewCoffee(coffees.FrenchPress, coffees.Large)

	blackSm := coffees.NewCoffee(coffees.Black, coffees.Small)
	blackMed := coffees.NewCoffee(coffees.Black, coffees.Medium)
	blackLar := coffees.NewCoffee(coffees.Black, coffees.Large)

	coffeeList := coffees.AsList(
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

	coffeeFilter := coffees.CoffeeFilter{}

	bitterCoffees := coffeeFilter.Filter(coffeeList, coffees.NewTypeSpecification(coffees.Bitter))

	fmt.Println("\nBitter Coffees:")
	for _, coffee := range bitterCoffees {
		fmt.Println(coffee.Name)
	}

	sweetAndSmallCoffees := coffeeFilter.Filter(
		coffeeList,
		coffees.NewAndSpecification(
			coffees.NewTypeSpecification(coffees.Sweet), coffees.NewSizeSpecification(coffees.Small),
		),
	)
	fmt.Println("\nSweet and Small Coffees:")
	for _, coffee := range sweetAndSmallCoffees {
		fmt.Println(coffee.Name)
	}

	fmt.Println("\n========> Builder Pattern <========")
	someMeal := meals.NewMealBuilder().
		Named("Super Mega Blaster Combo").
		WithMainDish().
		Named("Special Steak BBQ").
		WithDescription("Extra Large Juice Steak covered in BBQ sauce").
		WithSide().
		Named("Garlic Bread").
		WithDescription("4 delicious slices of garlic bread covered in cheese").
		WithDrink().
		Named("Premium Weiss Beer").
		WithSize(meals.Large).
		Build()
	fmt.Println("Built Meal:\n", someMeal.String())

	fmt.Println("\n========> Factory Pattern <========")
	breakfast := meals.NewMealHouseSpecial(meals.Breakfast)
	lunch := meals.NewMealHouseSpecial(meals.Lunch)
	dinner := meals.NewMealHouseSpecial(meals.Dinner)

	for _, houseSpecial := range []*meals.Meal{breakfast, lunch, dinner} {
		fmt.Println("House Special Meal:\n", houseSpecial.String())
		goLarge := meals.NewMealGoLarge(houseSpecial)
		fmt.Println("Go Large!:\n", goLarge.String())
	}

	fmt.Println("\n========> Singleton Pattern <========")

	db := transactions.GetSingletonDatabase()
	incomes := db.GetIncomesSum()
	expenses := db.GetExpensesSum()
	fmt.Println("Incomes:", incomes, "Expenses:", expenses, "Balance:", incomes+expenses)

	fmt.Println("\n========> Composite Pattern <========")
	neuron1, neuron2 := neurons.NewNeuron(), neurons.NewNeuron()
	layer1, layer2 := neurons.NewNeuronLayer(3), neurons.NewNeuronLayer(4)

	neurons.Connect(neuron1, neuron2)
	neurons.Connect(neuron1, layer1)
	neurons.Connect(layer2, neuron1)
	neurons.Connect(layer1, layer2)

	neurons.ShowConnections(neuron1, neuron2, layer1, layer2)

	fmt.Println("\n========> Façade Pattern <========")

	walletFacade := wallet.NewWalletFacade("abc", 1234)
	fmt.Println()

	err := walletFacade.AddMoneyToWallet("abc", 1234, 10)
	if err != nil {
		log.Fatalf("Error: %s\n", err.Error())
	}

	fmt.Println()
	err = walletFacade.DeductMoneyFromWallet("abc", 1234, 5)
	if err != nil {
		log.Fatalf("Error: %s\n", err.Error())
	}

	fmt.Println("\n========> Flyweight Pattern <========")
	cs.CounterStrike(10)

}
