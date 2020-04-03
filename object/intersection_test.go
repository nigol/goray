package object

import (
	"testing"
)

func TestAllPositiveHit(t *testing.T) {
	s := Sphere{}
	i1 := Intersection{1, s}
	i2 := Intersection{2, s}
	xs := Intersections{[]Intersection{i1, i2}}
	i := xs.Hit()
	if i != i1 {
		t.Errorf("Wrong hit for all positive.")
	}
}

func TestSomeNegativeHit(t *testing.T) {
	s := Sphere{}
	i1 := Intersection{-1, s}
	i2 := Intersection{1, s}
	xs := Intersections{[]Intersection{i1, i2}}
	i := xs.Hit()
	if i != i2 {
		t.Errorf("Wrong hit for some negative.")
	}
}

func TestAllNegativeHit(t *testing.T) {
	s := Sphere{}
	i1 := Intersection{-2, s}
	i2 := Intersection{-1, s}
	xs := Intersections{[]Intersection{i1, i2}}
	i := xs.Hit()
	if i.Object != nil {
		t.Errorf("Wrong hit for all negative.")
	}
}

func TestLowestNonnegativeHit(t *testing.T) {
	s := Sphere{}
	i1 := Intersection{5, s}
	i2 := Intersection{7, s}
	i3 := Intersection{-3, s}
	i4 := Intersection{2, s}
	xs := Intersections{[]Intersection{i1, i2, i3, i4}}
	i := xs.Hit()
	if i != i4 {
		t.Errorf("Wrong hit for lowest non negative.")
	}
}
