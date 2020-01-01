package common

import (
	"math"
)

func Equal(a float64, b float64) bool {
	e := 0.00001
	return math.Abs(a-b) < e
}
