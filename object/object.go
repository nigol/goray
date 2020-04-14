package object

import (
	"goray/material"
	"goray/ray"
	"goray/tuple"
)

type Object interface {
	Kind() string
	Intersect(r ray.Ray) Intersections
	NormalAt(p tuple.Tuple) tuple.Tuple
	GetMaterial() material.Material
}
