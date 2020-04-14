package ray

import (
	"goray/matrix"
	"goray/tuple"
)

type Ray struct {
	Origin    tuple.Tuple
	Direction tuple.Tuple
}

func (r Ray) Transform(m matrix.Matrix) Ray {
	ra := Ray{
		m.MulTuple(r.Origin),
		m.MulTuple(r.Direction),
	}
	return ra
}

func (r Ray) Position(t float64) tuple.Tuple {
	return r.Origin.Add(r.Direction.ScalarMul(t))
}
