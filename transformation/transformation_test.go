package transformation

import (
	"goray/tuple"
	"testing"
)

func TestMulTranslationMatrix(t *testing.T) {
	p := tuple.CreatePoint(-3, 4, 5)
	tr := Translation(5, -3, 2)
	if !tuple.CreatePoint(2, 1, 7).Equal(tr.MulTuple(p)) {
		t.Errorf("Mul translation went wrong.")
	}
}

func TestMulInverseTranslationMatrix(t *testing.T) {
	p := tuple.CreatePoint(-3, 4, 5)
	tr := Translation(5, -3, 2)
	inv := tr.Inverse4x4()
	if !tuple.CreatePoint(-8, 7, 3).Equal(inv.MulTuple(p)) {
		t.Errorf("Mul inverse translation went wrong.")
	}
}

func TestMulTranslationMatrixVector(t *testing.T) {
	v := tuple.CreateVector(-3, 4, 5)
	tr := Translation(5, -3, 2)
	if !v.Equal(tr.MulTuple(v)) {
		t.Errorf("Mul translation changed vector.")
	}
}

func TestMulScalingMatrix(t *testing.T) {
	p := tuple.CreatePoint(-4, 6, 8)
	tr := Scaling(2, 3, 4)
	if !tuple.CreatePoint(-8, 18, 32).Equal(tr.MulTuple(p)) {
		t.Errorf("Mul scaling went wrong.")
	}
}

func TestMulScalingMatrixVector(t *testing.T) {
	v := tuple.CreateVector(-4, 6, 8)
	tr := Scaling(2, 3, 4)
	if !tuple.CreateVector(-8, 18, 32).Equal(tr.MulTuple(v)) {
		t.Errorf("Mul scaling vector went wrong.")
	}
}

func TestMulScalingInverseMatrixVector(t *testing.T) {
	v := tuple.CreateVector(-4, 6, 8)
	tr := Scaling(2, 3, 4)
	inv := tr.Inverse4x4()
	if !tuple.CreateVector(-2, 2, 2).Equal(inv.MulTuple(v)) {
		t.Errorf("Mul inverse scaling vector went wrong.")
	}
}

func TestMulScalingMatrixReflection(t *testing.T) {
	p := tuple.CreatePoint(2, 3, 4)
	tr := Scaling(-1, 1, 1)
	if !tuple.CreatePoint(-2, 3, 4).Equal(tr.MulTuple(p)) {
		t.Errorf("Mul scaling reflection went wrong.")
	}
}
