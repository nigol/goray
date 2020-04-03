package ray

import (
	"goray/transformation"
	"goray/tuple"
	"testing"
)

func TestTranslate(t *testing.T) {
	r := Ray{
		tuple.CreatePoint(1, 2, 3),
		tuple.CreateVector(0, 1, 0),
	}
	m := transformation.Translation(3, 4, 5)
	r2 := r.Transform(m)
	if !tuple.CreatePoint(4, 6, 8).Equal(r2.Origin) || !tuple.CreateVector(0, 1, 0).Equal(r2.Direction) {
		t.Errorf("Ray translation is wrong.")
	}
}

func TestScale(t *testing.T) {
	r := Ray{
		tuple.CreatePoint(1, 2, 3),
		tuple.CreateVector(0, 1, 0),
	}
	m := transformation.Scaling(2, 3, 4)
	r2 := r.Transform(m)
	if !tuple.CreatePoint(2, 6, 12).Equal(r2.Origin) || !tuple.CreateVector(0, 3, 0).Equal(r2.Direction) {
		t.Errorf("Ray scaling is wrong.")
	}
}
