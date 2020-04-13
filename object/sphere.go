package object

import (
	"goray/material"
	"goray/matrix"
	"goray/ray"
	"goray/transformation"
	"goray/tuple"
	"math"
)

type Sphere struct {
	kind      string
	Transform matrix.Matrix
	Material  material.Material
}

func CreateSphere() Sphere {
	return Sphere{
		"sphere",
		transformation.Identity(),
		material.CreateMaterial(),
	}
}

func (s Sphere) Kind() string {
	return s.kind
}

func (s Sphere) Intersect(r ray.Ray) Intersections {
	// transform ray by the inverse of sphere's transformation matrix
	rt := r.Transform(s.Transform.Inverse4x4())
	// get vector from sphere center (at the world origin 0, 0, 0) to the ray origin
	sphereToRay := rt.Origin.Sub(tuple.CreatePoint(0, 0, 0))
	a := rt.Direction.Dot(rt.Direction)
	b := 2 * rt.Direction.Dot(sphereToRay)
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

func (s Sphere) NormalAt(p tuple.Tuple) tuple.Tuple {
	objectPoint := s.Transform.Inverse4x4().MulTuple(p)
	objectNormal := objectPoint.Sub(tuple.CreatePoint(0, 0, 0))
	transform3x3 := s.Transform.Submatrix(3, 3)
	worldNormal := transform3x3.Inverse3x3().Transpose().MulTuple(objectNormal)
	return worldNormal.Normalize()
}
