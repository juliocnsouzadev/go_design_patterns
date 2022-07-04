package meals

import (
	"strings"
)

type Side struct {
	Name        string
	Description string
}

func (s *Side) String() string {
	var sb strings.Builder
	sb.WriteString(s.Name)
	sb.WriteString(" (")
	sb.WriteString(s.Description)
	sb.WriteString(")")
	return sb.String()
}

type SideBuilder struct {
	*Builder
}

func (b *SideBuilder) Named(name string) *SideBuilder {
	b.meal.Side.Name = name
	return b
}

func (b *SideBuilder) WithDescription(description string) *SideBuilder {
	b.meal.Side.Description = description
	return b
}
