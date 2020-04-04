package object

import (
	"goray/ray"
)

type Object interface {
	Kind() string
	Intersect(r ray.Ray) Intersections
}
