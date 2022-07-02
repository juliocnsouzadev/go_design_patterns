package coffees

type Type string

type Size string

const (
	Bitter Type = "bitter"
	Light  Type = "light"
	Sweet  Type = "sweet"
	Small  Size = "small"
	Medium Size = "medium"
	Large  Size = "large"
)

type coffee struct {
	Name string
	Type Type
	Size Size
}

func newCoffee(name string, coffeeType Type, size Size) *coffee {
	return &coffee{
		Name: name,
		Type: coffeeType,
		Size: size,
	}
}
