package main

import (
	"testing"
)

func TestColor(t *testing.T) {
	c := Color{
		-0.5,
		0.4,
		1.7,
	}
	comp := c.red == -0.5 && c.green == 0.4 && c.blue == 1.7
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
		t.Errorf("Color sub is not working. %f, %f, %f", c3.red, c3.green, c3.blue)
	}
}

func TestScalarMulColor(t *testing.T) {
	c1 := Color{0.2, 0.3, 0.4}
	c3 := c1.ScalarMul(2)
	c4 := Color{0.4, 0.6, 0.8}
	if !c3.Equal(c4) {
		t.Errorf("Color scalar mul is not working. %f, %f, %f", c3.red, c3.green, c3.blue)
	}
}
