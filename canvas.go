package main

type Color struct {
	red   float64
	green float64
	blue  float64
}

func (c Color) Add(o Color) Color {
	return Color{c.red + o.red, c.green + o.green, c.blue + o.blue}
}
