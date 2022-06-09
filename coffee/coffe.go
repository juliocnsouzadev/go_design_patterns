package coffee

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

type Coffee struct {
	Name string
	Type Type
	Size Size
}

func NewCoffee(name string, coffeeType Type, size Size) *Coffee {
	return &Coffee{
		Name: name,
		Type: coffeeType,
		Size: size,
	}
}