package meals

import "fmt"

type HouseSpecial string

const (
	Breakfast HouseSpecial = "Breakfast"
	Lunch     HouseSpecial = "Lunch"
	Dinner    HouseSpecial = "Dinner"
)

func NewMealHouseSpecial(special HouseSpecial) *Meal {
	name := fmt.Sprintf("%s House Special", special)
	switch special {
	case Breakfast:

		return NewMealBuilder().
			Named(name).
			WithMainDish().
			Named("Omelette").
			WithDescription("Large Omelette made with 4 eggs and cheese").
			WithSide().
			Named("Toasts").
			WithDescription("4 delicious slices of toasts").
			WithDrink().
			Named("Orange Juice").
			WithSize(Large).
			Build()
	case Lunch:
		return NewMealBuilder().
			Named(name).
			WithMainDish().
			Named("Special Steak BBQ").
			WithDescription("Extra Large Juice Steak covered in BBQ sauce").
			WithSide().
			Named("Garlic Bread").
			WithDescription("4 delicious slices of garlic bread covered in cheese").
			WithDrink().
			Named("Premium Weiss Beer").
			WithSize(Large).
			Build()
	case Dinner:
		return NewMealBuilder().
			Named(name).
			WithMainDish().
			Named("Special Curry Chicken").
			WithDescription("Delicious spicy curry chicken").
			WithSide().
			Named("White rice").
			WithDescription("Medium portion of rice with garlic").
			WithDrink().
			Named("Premium Larger Beer").
			WithSize(Large).
			Build()
	default:
		panic("Unknown House special")
	}
}
