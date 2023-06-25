package meals

func NewBreakfast(name string) *MealImpl {
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
		WithSize(Medium).
		Build()
}

func NewLunch(name string) *MealImpl {
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
		WithSize(Medium).
		Build()
}
func NewDinner(name string) *MealImpl {
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
		WithSize(Medium).
		Build()
}
