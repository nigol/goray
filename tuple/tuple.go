package tuple

import (
	"goray/common"
	"math"
)

type Tuple struct {
	X float64
	Y float64
	Z float64
	W float64
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

func (t Tuple) Equal(o Tuple) bool {
	return common.Equal(t.X, o.X) && common.Equal(t.Y, o.Y) && common.Equal(t.Z, o.Z) && common.Equal(t.W, o.W)
}

func (t Tuple) Add(o Tuple) Tuple {
	return CreateTuple(t.X+o.X, t.Y+o.Y, t.Z+o.Z, t.W+o.W)
}

func (t Tuple) Sub(o Tuple) Tuple {
	return CreateTuple(t.X-o.X, t.Y-o.Y, t.Z-o.Z, t.W-o.W)
}

func (t Tuple) Negate() Tuple {
	return CreateTuple(0, 0, 0, 0).Sub(t)
}

func (t Tuple) ScalarMul(n float64) Tuple {
	return CreateTuple(t.X*n, t.Y*n, t.Z*n, t.W*n)
}

func (t Tuple) ScalarDiv(n float64) Tuple {
	return CreateTuple(t.X/n, t.Y/n, t.Z/n, t.W/n)
}

func (t Tuple) Magnitude() float64 {
	sumOfSq := t.X*t.X + t.Y*t.Y + t.Z*t.Z + t.W*t.W
	return math.Sqrt(sumOfSq)
}

func (t Tuple) Normalize() Tuple {
	mag := t.Magnitude()
	return CreateTuple(t.X/mag, t.Y/mag, t.Z/mag, t.W/mag)
}

func (t Tuple) Dot(o Tuple) float64 {
	return t.X*o.X + t.Y*o.Y + t.Z*o.Z + t.W + o.W
}

func (t Tuple) Cross(o Tuple) Tuple {
	return CreateVector(t.Y*o.Z-t.Z*o.Y, t.Z*o.X-t.X*o.Z, t.X*o.Y-t.Y*o.X)
}

func (v Tuple) Reflect(n Tuple) Tuple {
	s := 2 * v.Dot(n)
	return n.ScalarMul(s).Sub(v)
}

func (t Tuple) Mul(o Tuple) Tuple {
	return CreateTuple(t.X*o.X, t.Y*o.Y, t.Z*o.Z, t.W*o.W)
}
