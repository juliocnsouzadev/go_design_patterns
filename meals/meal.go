package meals

import "strings"

type Meal struct {
	name     string
	mainDish *MainDish
	side     *Side
	drink    *Drink
}

func (m *Meal) String() string {
	var sb strings.Builder
	sb.WriteString(m.name)
	sb.WriteString(" is composed of:")
	sb.WriteString("\n- main dish:\n")
	sb.WriteString(m.mainDish.String())
	sb.WriteString("\n- side:\n")
	sb.WriteString(m.side.String())
	sb.WriteString("\n- drink:\n")
	sb.WriteString(m.drink.String())
	return sb.String()
}

type Builder struct {
	meal *Meal
}

func NewMealBuilder() *Builder {
	return &Builder{&Meal{
		"",
		&MainDish{},
		&Side{},
		&Drink{},
	}}
}

func (b *Builder) Named(name string) *Builder {
	b.meal.name = name
	return b
}

func (b *Builder) WithMainDish() *MainDishBuilder {
	return &MainDishBuilder{b}
}

func (b *Builder) WithSide() *SideBuilder {
	return &SideBuilder{b}
}

func (b *Builder) WithDrink() *DrinkBuilder {
	return &DrinkBuilder{b}
}

func (b *Builder) Build() *Meal {
	return b.meal
}
