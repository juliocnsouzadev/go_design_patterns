package mealsapi

import (
	"fmt"
	"go_design_patterns/meals"
)

func NewMealHouseSpecial(special string) Meal {
	name := fmt.Sprintf("%s House Special", special)
	switch special {
	case "Breakfast":
		return meals.NewBreakfast(name)
	case "Lunch":
		return meals.NewLunch(name)
	case "Dinner":
		return meals.NewDinner(name)
	default:
		panic("Unknown House special")
	}
}
