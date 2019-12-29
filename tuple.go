package main

import (
	"math"
)

type Tuple struct {
	x float64
	y float64
	z float64
	w float64
}

func CreateTuple(x float64, y float64, z float64, w float64) Tuple {
	result := Tuple{
		x,
		y,
		z,
		w,
	}
	return result
}

func CreatePoint(x float64, y float64, z float64) Tuple {
	return CreateTuple(x, y, z, 1.0)
}

func CreateVector(x float64, y float64, z float64) Tuple {
	return CreateTuple(x, y, z, 0.0)
}

func equal(a float64, b float64) bool {
	e := 0.00001
	return math.Abs(a-b) < e
}

func (t Tuple) Equal(o Tuple) bool {
	return equal(t.x, o.x) && equal(t.y, o.y) && equal(t.z, o.z) && equal(t.w, o.w)
}

func (t Tuple) Add(o Tuple) Tuple {
	return CreateTuple(t.x+o.x, t.y+o.y, t.z+o.z, t.w+o.w)
}

func (t Tuple) Sub(o Tuple) Tuple {
	return CreateTuple(t.x-o.x, t.y-o.y, t.z-o.z, t.w-o.w)
}

func (t Tuple) Negate() Tuple {
	return CreateTuple(0, 0, 0, 0).Sub(t)
}

func (t Tuple) ScalarMul(n float64) Tuple {
	return CreateTuple(t.x*n, t.y*n, t.z*n, t.w*n)
}

func (t Tuple) ScalarDiv(n float64) Tuple {
	return CreateTuple(t.x/n, t.y/n, t.z/n, t.w/n)
}

func (t Tuple) Magnitude() float64 {
	sumOfSq := t.x*t.x + t.y*t.y + t.z*t.z + t.w*t.w
	return math.Sqrt(sumOfSq)
}
