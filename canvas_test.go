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
	comp := c3.red == 1.6 && c3.green == 0.7 && c3.blue == 1.0
	if !comp {
		t.Errorf("Color add is not working.")
	}
}
