package canvas

import (
	"strings"
	"testing"
)

func TestColor(t *testing.T) {
	c := Color{
		-0.5,
		0.4,
		1.7,
	}
	comp := c.Red == -0.5 && c.Green == 0.4 && c.Blue == 1.7
	if !comp {
		t.Errorf("Color is incorrect.")
	}
}

func TestAddColor(t *testing.T) {
	c1 := Color{0.9, 0.6, 0.75}
	c2 := Color{0.7, 0.1, 0.25}
	c3 := c1.Add(c2)
	c4 := Color{1.6, 0.7, 1.0}
	if !c3.Equal(c4) {
		t.Errorf("Color add is not working.")
	}
}

func TestSubColor(t *testing.T) {
	c1 := Color{0.9, 0.6, 0.75}
	c2 := Color{0.7, 0.1, 0.25}
	c3 := c1.Sub(c2)
	c4 := Color{0.2, 0.5, 0.5}
	if !c3.Equal(c4) {
		t.Errorf("Color sub is not working. %f, %f, %f", c3.Red, c3.Green, c3.Blue)
	}
}

func TestScalarMulColor(t *testing.T) {
	c1 := Color{0.2, 0.3, 0.4}
	c3 := c1.ScalarMul(2)
	c4 := Color{0.4, 0.6, 0.8}
	if !c3.Equal(c4) {
		t.Errorf("Color scalar mul is not working. %f, %f, %f", c3.Red, c3.Green, c3.Blue)
	}
}

func TestProductColor(t *testing.T) {
	c1 := Color{1.0, 0.2, 0.4}
	c2 := Color{0.9, 1.0, 0.1}
	c3 := c1.Product(c2)
	c4 := Color{0.9, 0.2, 0.04}
	if !c3.Equal(c4) {
		t.Errorf("Color product is not working. %f, %f, %f", c3.Red, c3.Green, c3.Blue)
	}
}

func TestCreateCanvas(t *testing.T) {
	c := CreateCanvas(10, 20)
	comp := c.Width == 10 && c.Height == 20
	if !comp {
		t.Errorf("Canvas dimensions are wrong.")
	}
	b := Color{0.0, 0.0, 0.0}
	comp = true
	for _, color := range c.pixels {
		comp = comp && b.Equal(color)
	}
	if !comp {
		t.Errorf("All colors should be black.")
	}
}

func TestWritePixel(t *testing.T) {
	c := CreateCanvas(10, 20)
	red := Color{1.0, 0.0, 0.0}
	c.WritePixel(2, 3, red)
	if !red.Equal(c.PixelAt(2, 3)) {
		t.Errorf("Color should be red.")
	}
}

func TestWritePixelOutOfBounds(t *testing.T) {
	c := CreateCanvas(10, 20)
	red := Color{1.0, 0.0, 0.0}
	c.WritePixel(-2, 30, red)
	if !red.Equal(c.PixelAt(0, 0)) {
		t.Errorf("Color should be red.")
	}
}

func TestPpm(t *testing.T) {
	c := CreateCanvas(5, 3)
	c1 := Color{1.5, 0.0, 0.0}
	c2 := Color{0.0, 0.5, 0.0}
	c3 := Color{-0.5, 0.0, 1.0}
	c.WritePixel(0, 0, c1)
	c.WritePixel(2, 1, c2)
	c.WritePixel(4, 2, c3)
	slc := strings.SplitN(c.Ppm(), "\n", -1)
	comp := slc[0] == "P3" && slc[1] == "5 3" && slc[2] == "255"
	if !comp {
		t.Errorf("PPM header is not valid. %s", c.Ppm())
	}
	comp = slc[3] == "255 0 0 " &&
		slc[5] == "0 0 0 0 128 0 0 0 0 0 0 0 0 0 0 " &&
		slc[6] == "0 0 0 0 0 0 0 0 0 0 0 255 "
	if !comp {
		t.Errorf("PPM body is not valid. %s", c.Ppm())
	}
}
