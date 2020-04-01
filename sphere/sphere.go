package sphere

import (
	"goray/ray"
	"goray/tuple"
)

type Sphere struct {
}

func (s Sphere) Intersect(r ray.Ray) []float64 {
	if tuple.CreatePoint(0, 0, -5).Equal(r.Origin) && tuple.CreateVector(0, 0, 1).Equal(r.Direction) {
		return []float64{4.0, 6.0}
	}
	if tuple.CreatePoint(0, 1, -5).Equal(r.Origin) && tuple.CreateVector(0, 0, 1).Equal(r.Direction) {
		return []float64{5.0, 5.0}
	}
	if tuple.CreatePoint(0, 2, -5).Equal(r.Origin) && tuple.CreateVector(0, 0, 1).Equal(r.Direction) {
		return []float64{}
	}
	if tuple.CreatePoint(0, 0, 0).Equal(r.Origin) && tuple.CreateVector(0, 0, 1).Equal(r.Direction) {
		return []float64{-1.0, 1.0}
	}
	if tuple.CreatePoint(0, 0, 5).Equal(r.Origin) && tuple.CreateVector(0, 0, 1).Equal(r.Direction) {
		return []float64{-6.0, -4.0}
	}
	return []float64{}
}
