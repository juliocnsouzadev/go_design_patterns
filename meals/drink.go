package meals

import (
	"strings"
)

type Size int

const (
	Small  Size = 300
	Medium Size = 500
	Large  Size = 700
)

type Drink struct {
	Name string
	Size Size
}

func (d *Drink) String() string {
	var sb strings.Builder
	sb.WriteString(d.Name)
	switch d.Size {
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
	b.meal.Drink.Name = name
	return b
}

func (b *DrinkBuilder) WithSize(size Size) *DrinkBuilder {
	b.meal.Drink.Size = size
	return b
}
