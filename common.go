package main

import (
	"math"
)

func equal(a float64, b float64) bool {
	e := 0.00001
	return math.Abs(a-b) < e
}
