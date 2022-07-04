package meals

import (
	"strings"
)

type MainDish struct {
	Name        string
	Description string
}

func (m *MainDish) String() string {
	var sb strings.Builder
	sb.WriteString(m.Name)
	sb.WriteString(" (")
	sb.WriteString(m.Description)
	sb.WriteString(")")
	return sb.String()
}

type MainDishBuilder struct {
	*Builder
}

func (b *MainDishBuilder) Named(name string) *MainDishBuilder {
	b.meal.MainDish.Name = name
	return b
}

func (b *MainDishBuilder) WithDescription(description string) *MainDishBuilder {
	b.meal.MainDish.Description = description
	return b
}
