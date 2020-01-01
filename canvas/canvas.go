package canvas

import (
	"goray/common"
)

type Color struct {
	Red   float64
	Green float64
	Blue  float64
}

func (c Color) Equal(o Color) bool {
	return common.Equal(c.Red, o.Red) && common.Equal(c.Green, o.Green) && common.Equal(c.Blue, o.Blue)
}

func (c Color) Add(o Color) Color {
	return Color{c.Red + o.Red, c.Green + o.Green, c.Blue + o.Blue}
}

func (c Color) Sub(o Color) Color {
	return Color{c.Red - o.Red, c.Green - o.Green, c.Blue - o.Blue}
}

func (c Color) ScalarMul(n float64) Color {
	return Color{c.Red * n, c.Green * n, c.Blue * n}
}

func (c Color) Product(o Color) Color {
	return Color{c.Red * o.Red, c.Green * o.Green, c.Blue * o.Blue}
}

type Canvas struct {
	Width  int
	Height int
	pixels []Color
}

func CreateCanvas(w int, h int) Canvas {
	pixs := make([]Color, w*h)
	result := Canvas{
		w,
		h,
		pixs,
	}
	return result
}
