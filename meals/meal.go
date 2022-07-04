package meals

import (
	"encoding/json"
	"strings"
)

type Meal struct {
	Name     string
	MainDish *MainDish
	Side     *Side
	Drink    *Drink
}

func (m *Meal) deepCopy() *Meal {
	byt, _ := json.Marshal(m)
	meal := Meal{}
	json.Unmarshal(byt, &meal)
	return &meal
}

func (m *Meal) String() string {
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

func (b *Builder) Build() *Meal {
	return b.meal
}
