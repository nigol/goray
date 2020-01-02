package canvas

import (
	"goray/common"
	"strconv"
	"strings"
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

func (can Canvas) WritePixel(x int, y int, c Color) {
	can.pixels[y*can.Width+x] = c
}

func (can Canvas) PixelAt(x int, y int) Color {
	return can.pixels[y*can.Width+x]
}

func (can Canvas) Ppm() string {
	var sb strings.Builder
	ppmHeader(can, &sb)
	ppmBody(can, &sb)
	return sb.String()
}

func ppmHeader(can Canvas, sb *strings.Builder) {
	sb.WriteString("P3\n")
	sb.WriteString(strconv.Itoa(can.Width))
	sb.WriteString(" ")
	sb.WriteString(strconv.Itoa(can.Height))
	sb.WriteString("\n")
	sb.WriteString("255\n")
}

func ppmBody(can Canvas, sb *strings.Builder) {
	for i, color := range can.pixels {
		sb.WriteString(normalizeColor(color.Red))
		sb.WriteString(" ")
		sb.WriteString(normalizeColor(color.Green))
		sb.WriteString(" ")
		sb.WriteString(normalizeColor(color.Blue))
		sb.WriteString(" ")
		if i%5 == 0 {
			sb.WriteString("\n")
		}
	}
	sb.WriteString("\n")
}

func normalizeColor(c float64) string {
	norm := int(c * 256)
	if norm > 255 {
		norm = 255
	}
	if norm < 0 {
		norm = 0
	}
	return strconv.Itoa(norm)
}
