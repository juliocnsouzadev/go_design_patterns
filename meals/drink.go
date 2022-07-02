package meals

import "strings"

type Size int

const (
	Small  Size = 300
	Medium Size = 500
	Large  Size = 700
)

type Drink struct {
	name string
	size Size
}

func (d *Drink) String() string {
	var sb strings.Builder
	sb.WriteString(d.name)
	switch d.size {
	case Small:
		sb.WriteString(" (small)")
	case Medium:
		sb.WriteString(" (medium)")
	case Large:
		sb.WriteString(" (large)")
	}
	return sb.String()
}

type DrinkBuilder struct {
	*Builder
}

func (b *DrinkBuilder) Named(name string) *DrinkBuilder {
	b.meal.drink.name = name
	return b
}

func (b *DrinkBuilder) WithSize(size Size) *DrinkBuilder {
	b.meal.drink.size = size
	return b
}
