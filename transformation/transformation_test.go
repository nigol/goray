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
