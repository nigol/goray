package object

import (
	"goray/ray"
	"goray/tuple"
	"math"
)

type Sphere struct {
}

func (s Sphere) Intersect(r ray.Ray) Intersections {
	// get vector from sphere center (at the wolrd origin 0, 0, 0) to the ray origin
	sphereToRay := r.Origin.Sub(tuple.CreatePoint(0, 0, 0))
	a := r.Direction.Dot(r.Direction)
	b := 2 * r.Direction.Dot(sphereToRay)
	c := sphereToRay.Dot(sphereToRay) - 1
	discriminant := b*b - 4*a*c
	if discriminant < 0 {
		return Intersections{}
	} else {
		return Intersections{
			[]Intersection{
				{
					(-b - math.Sqrt(discriminant)) / (2 * a),
					s,
				},
				{
					(-b + math.Sqrt(discriminant)) / (2 * a),
					s,
				},
			},
		}
	}
}
