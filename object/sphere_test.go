package object

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
	if xs.Xs[0].Object != s && xs.Xs[1].Object != s && xs.Xs[0].T != 4.0 && xs.Xs[1].T != 6.0 {
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
	if xs.Xs[0].Object != s && xs.Xs[1].Object != s && xs.Xs[0].T != 5.0 && xs.Xs[1].T != 5.0 {
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
	if len(xs.Xs) != 0 {
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
	if xs.Xs[0].Object != s && xs.Xs[1].Object != s && xs.Xs[0].T != -1.0 && xs.Xs[1].T != 1.0 {
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
	if xs.Xs[0].Object != s && xs.Xs[1].Object != s && xs.Xs[0].T != -6.0 && xs.Xs[1].T != -4.0 {
		t.Errorf("Intersection behind is wrong.")
	}
}