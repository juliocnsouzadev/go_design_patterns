package meals

import "strings"

type MainDish struct {
	name        string
	description string
}

func (m *MainDish) String() string {
	var sb strings.Builder
	sb.WriteString(m.name)
	sb.WriteString(" (")
	sb.WriteString(m.description)
	sb.WriteString(")")
	return sb.String()
}

type MainDishBuilder struct {
	*Builder
}

func (b *MainDishBuilder) Named(name string) *MainDishBuilder {
	b.meal.mainDish.name = name
	return b
}

func (b *MainDishBuilder) WithDescription(description string) *MainDishBuilder {
	b.meal.mainDish.description = description
	return b
}
