package meals

import (
	"strings"
)

type MealImpl struct {
	Name     string
	MainDish *MainDish
	Side     *Side
	Drink    *Drink
}

func (m *MealImpl) GoLarge() {
	m.Drink.Size = Large
}

func (m *MealImpl) String() string {
	var sb strings.Builder
	sb.WriteString(m.Name)
	sb.WriteString(" is composed of:")
	sb.WriteString("\n- main dish:\n")
	sb.WriteString(m.MainDish.String())
	sb.WriteString("\n- Side:\n")
	sb.WriteString(m.Side.String())
	sb.WriteString("\n- Drink:\n")
	sb.WriteString(m.Drink.String())
	return sb.String()
}

type Builder struct {
	meal *MealImpl
}

func NewMealBuilder() *Builder {
	return &Builder{&MealImpl{
		"",
		&MainDish{},
		&Side{},
		&Drink{},
	}}
}

func (b *Builder) Named(name string) *Builder {
	b.meal.Name = name
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

func (b *Builder) Build() *MealImpl {
	return b.meal
}
