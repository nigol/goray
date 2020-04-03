package object

import (
	"goray/ray"
)

type Object interface {
	Intersect(r ray.Ray) Intersections
}
