package main

type Color struct {
	red   float64
	green float64
	blue  float64
}

func (c Color) Equal(o Color) bool {
	return equal(c.red, o.red) && equal(c.green, o.green) && equal(c.blue, o.blue)
}

func (c Color) Add(o Color) Color {
	return Color{c.red + o.red, c.green + o.green, c.blue + o.blue}
}

func (c Color) Sub(o Color) Color {
	return Color{c.red - o.red, c.green - o.green, c.blue - o.blue}
}

func (c Color) ScalarMul(n float64) Color {
	return Color{c.red * n, c.green * n, c.blue * n}
}

func (c Color) Product(o Color) Color {
	return Color{c.red * o.red, c.green * o.green, c.blue * o.blue}
}
