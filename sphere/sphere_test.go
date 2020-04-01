package sphere

import (
	"goray/ray"
	"goray/tuple"
	"testing"
)

func TestIntersect2Points(t *testing.T) {
	r := ray.Ray{
		tuple.CreatePoint(0, 0, -5),
		tuple.CreateVector(0, 0, 1),
	}
	s := Sphere{}
	xs := s.Intersect(r)
	if xs[0] != 4.0 && xs[1] != 6.0 {
		t.Errorf("Intersection of 2 points is wrong.")
	}
}

func TestIntersectTangent(t *testing.T) {
	r := ray.Ray{
		tuple.CreatePoint(0, 1, -5),
		tuple.CreateVector(0, 0, 1),
	}
	s := Sphere{}
	xs := s.Intersect(r)
	if xs[0] != 5.0 && xs[1] != 5.0 {
		t.Errorf("Intersection at tangent is wrong.")
	}
}

func TestIntersectMiss(t *testing.T) {
	r := ray.Ray{
		tuple.CreatePoint(0, 2, -5),
		tuple.CreateVector(0, 0, 1),
	}
	s := Sphere{}
	xs := s.Intersect(r)
	if len(xs) != 0 {
		t.Errorf("Intersection miss is wrong.")
	}
}

func TestIntersectInside(t *testing.T) {
	r := ray.Ray{
		tuple.CreatePoint(0, 0, 0),
		tuple.CreateVector(0, 0, 1),
	}
	s := Sphere{}
	xs := s.Intersect(r)
	if xs[0] != -1.0 && xs[1] != 1.0 {
		t.Errorf("Intersection inside is wrong.")
	}
}

func TestIntersectBehind(t *testing.T) {
	r := ray.Ray{
		tuple.CreatePoint(0, 0, 5),
		tuple.CreateVector(0, 0, 1),
	}
	s := Sphere{}
	xs := s.Intersect(r)
	if xs[0] != -6.0 && xs[1] != -4.0 {
		t.Errorf("Intersection behind is wrong.")
	}
}
