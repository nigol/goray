package object

import (
	"goray/material"
	"goray/ray"
	"goray/transformation"
	"goray/tuple"
	"math"
	"testing"
)

func TestIntersect2Points(t *testing.T) {
	r := ray.Ray{
		tuple.CreatePoint(0, 0, -5),
		tuple.CreateVector(0, 0, 1),
	}
	s := CreateSphere()
	xs := s.Intersect(r)
	if xs.Xs[0].Object.Kind() != s.Kind() &&
		xs.Xs[1].Object.Kind() != s.Kind() &&
		xs.Xs[0].T != 4.0 && xs.Xs[1].T != 6.0 {
		t.Errorf("Intersection of 2 points is wrong.")
	}
}

func TestIntersectTangent(t *testing.T) {
	r := ray.Ray{
		tuple.CreatePoint(0, 1, -5),
		tuple.CreateVector(0, 0, 1),
	}
	s := CreateSphere()
	xs := s.Intersect(r)
	if xs.Xs[0].Object.Kind() != s.Kind() &&
		xs.Xs[1].Object.Kind() != s.Kind() &&
		xs.Xs[0].T != 5.0 && xs.Xs[1].T != 5.0 {
		t.Errorf("Intersection at tangent is wrong.")
	}
}

func TestIntersectMiss(t *testing.T) {
	r := ray.Ray{
		tuple.CreatePoint(0, 2, -5),
		tuple.CreateVector(0, 0, 1),
	}
	s := CreateSphere()
	xs := s.Intersect(r)
	if xs.Defined() {
		t.Errorf("Intersection miss is wrong.")
	}
}

func TestIntersectInside(t *testing.T) {
	r := ray.Ray{
		tuple.CreatePoint(0, 0, 0),
		tuple.CreateVector(0, 0, 1),
	}
	s := CreateSphere()
	xs := s.Intersect(r)
	if xs.Xs[0].Object.Kind() != s.Kind() &&
		xs.Xs[1].Object.Kind() != s.Kind() && xs.Xs[0].T != -1.0 && xs.Xs[1].T != 1.0 {
		t.Errorf("Intersection inside is wrong.")
	}
}

func TestIntersectBehind(t *testing.T) {
	r := ray.Ray{
		tuple.CreatePoint(0, 0, 5),
		tuple.CreateVector(0, 0, 1),
	}
	s := CreateSphere()
	xs := s.Intersect(r)
	if xs.Xs[0].Object.Kind() != s.Kind() &&
		xs.Xs[1].Object.Kind() != s.Kind() && xs.Xs[0].T != -6.0 && xs.Xs[1].T != -4.0 {
		t.Errorf("Intersection behind is wrong.")
	}
}

func TestDefaultTransformation(t *testing.T) {
	s := CreateSphere()
	if !s.Transform.Equal(transformation.Identity()) {
		t.Errorf("Sphere initial transformation is incorrect.")
	}
}

func TestChangeTransformation(t *testing.T) {
	s := CreateSphere()
	s.Transform = transformation.Translation(2, 3, 4)
	if !s.Transform.Equal(transformation.Translation(2, 3, 4)) {
		t.Errorf("Sphere set transformation doesn't work.")
	}
}

func TestIntersectScaledSphereRay(t *testing.T) {
	s := CreateSphere()
	r := ray.Ray{
		tuple.CreatePoint(0, 0, -5),
		tuple.CreateVector(0, 0, 1),
	}
	s.Transform = transformation.Scaling(2, 2, 2)
	xs := s.Intersect(r)
	if len(xs.Xs) != 2 || xs.Xs[0].T != 3 || xs.Xs[1].T != 7 {
		t.Errorf("Scaled intersection is wrong.")
	}
}

func TestIntersectTranslatedRay(t *testing.T) {
	s := CreateSphere()
	r := ray.Ray{
		tuple.CreatePoint(0, 0, -5),
		tuple.CreateVector(0, 0, 1),
	}
	s.Transform = transformation.Translation(5, 0, 0)
	xs := s.Intersect(r)
	if len(xs.Xs) != 0 {
		t.Errorf("Translated intersection is wrong.")
	}
}

func TestNormalXAxis(t *testing.T) {
	s := CreateSphere()
	n := s.NormalAt(tuple.CreatePoint(1, 0, 0))
	if !tuple.CreateVector(1, 0, 0).Equal(n) {
		t.Errorf("Normal at X is wrong.")
	}
}

func TestNormalYAxis(t *testing.T) {
	s := CreateSphere()
	n := s.NormalAt(tuple.CreatePoint(0, 1, 0))
	if !tuple.CreateVector(0, 1, 0).Equal(n) {
		t.Errorf("Normal at Y is wrong.")
	}
}

func TestNormalZAxis(t *testing.T) {
	s := CreateSphere()
	n := s.NormalAt(tuple.CreatePoint(0, 0, 1))
	if !tuple.CreateVector(0, 0, 1).Equal(n) {
		t.Errorf("Normal at Z is wrong.")
	}
}

func TestNormalNonAxial(t *testing.T) {
	s := CreateSphere()
	n := s.NormalAt(tuple.CreatePoint(math.Sqrt(3)/3, math.Sqrt(3)/3, math.Sqrt(3)/3))
	if !tuple.CreateVector(math.Sqrt(3)/3, math.Sqrt(3)/3, math.Sqrt(3)/3).Equal(n) {
		t.Errorf("Non axial normal is wrong.")
	}
}

func TestNormalNormalized(t *testing.T) {
	s := CreateSphere()
	n := s.NormalAt(tuple.CreatePoint(math.Sqrt(3)/3, math.Sqrt(3)/3, math.Sqrt(3)/3))
	if !tuple.CreateVector(math.Sqrt(3)/3, math.Sqrt(3)/3, math.Sqrt(3)/3).Normalize().Equal(n) {
		t.Errorf("Non axial normal is wrong.")
	}
}

func TestNormalTranslated(t *testing.T) {
	s := CreateSphere()
	s.Transform = transformation.Translation(0, 1, 0)
	n := s.NormalAt(tuple.CreatePoint(0, 1.70711, -0.70711))
	if !tuple.CreateVector(0, 0.70711, -0.70711).Equal(n) {
		t.Errorf("Translated normal is wrong.")
	}
}

func TestNormalTransformed(t *testing.T) {
	s := CreateSphere()
	m := transformation.Scaling(1, 0.5, 1).Mul(transformation.RotationZ(math.Pi / 5))
	s.Transform = m
	n := s.NormalAt(tuple.CreatePoint(0, math.Sqrt(2)/2, -1*math.Sqrt(2)/2))
	if !tuple.CreateVector(0, 0.97014, -0.24254).Equal(n) {
		t.Errorf("Transformed normal is wrong.")
	}
}

func TestSphereDefaultMaterial(t *testing.T) {
	s := CreateSphere()
	m := s.Material
	if !material.CreateMaterial().Equal(m) {
		t.Errorf("Sphere should have default material.")
	}
}

func TestAssignMaterial(t *testing.T) {
	s := CreateSphere()
	m := material.CreateMaterial()
	m.Ambient = 1
	s.Material = m
	if !s.Material.Equal(m) {
		t.Errorf("Sphere has wrong material assigned.")
	}
}
