package meals

import "strings"

type Side struct {
	name        string
	description string
}

func (s *Side) String() string {
	var sb strings.Builder
	sb.WriteString(s.name)
	sb.WriteString(" (")
	sb.WriteString(s.description)
	sb.WriteString(")")
	return sb.String()
}

type SideBuilder struct {
	*Builder
}

func (b *SideBuilder) Named(name string) *SideBuilder {
	b.meal.side.name = name
	return b
}

func (b *SideBuilder) WithDescription(description string) *SideBuilder {
	b.meal.side.description = description
	return b
}
