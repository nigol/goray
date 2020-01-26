package transformation

import (
	"goray/tuple"
	"math"
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

func TestRotationX(t *testing.T) {
	p := tuple.CreatePoint(0, 1, 0)
	hq := RotationX(math.Pi / 4)
	fq := RotationX(math.Pi / 2)
	if !tuple.CreatePoint(0, math.Sqrt(2)/2, math.Sqrt(2)/2).Equal(hq.MulTuple(p)) {
		t.Errorf("Rotation X half quarter went wrong.")
	}
	if !tuple.CreatePoint(0, 0, 1).Equal(fq.MulTuple(p)) {
		t.Errorf("Rotation X full quarter went wrong.")
	}
}

func TestRotationXInverse(t *testing.T) {
	p := tuple.CreatePoint(0, 1, 0)
	hq := RotationX(math.Pi / 4)
	inv := hq.Inverse4x4()
	if !tuple.CreatePoint(0, math.Sqrt(2)/2, -1*math.Sqrt(2)/2).Equal(inv.MulTuple(p)) {
		t.Errorf("Rotation X inverse half quarter went wrong.")
	}
}

func TestRotationY(t *testing.T) {
	p := tuple.CreatePoint(0, 0, 1)
	hq := RotationY(math.Pi / 4)
	fq := RotationY(math.Pi / 2)
	if !tuple.CreatePoint(math.Sqrt(2)/2, 0, math.Sqrt(2)/2).Equal(hq.MulTuple(p)) {
		t.Errorf("Rotation Y half quarter went wrong.")
	}
	if !tuple.CreatePoint(1, 0, 0).Equal(fq.MulTuple(p)) {
		t.Errorf("Rotation Y full quarter went wrong.")
	}
}

func TestRotationZ(t *testing.T) {
	p := tuple.CreatePoint(0, 1, 0)
	hq := RotationZ(math.Pi / 4)
	fq := RotationZ(math.Pi / 2)
	if !tuple.CreatePoint(-1*math.Sqrt(2)/2, math.Sqrt(2)/2, 0).Equal(hq.MulTuple(p)) {
		t.Errorf("Rotation Z half quarter went wrong.")
	}
	if !tuple.CreatePoint(-1, 0, 0).Equal(fq.MulTuple(p)) {
		t.Errorf("Rotation Z full quarter went wrong.")
	}
}

func TestShearingXtoY(t *testing.T) {
	p := tuple.CreatePoint(2, 3, 4)
	tr := Shearing(1, 0, 0, 0, 0, 0)
	if !tuple.CreatePoint(5, 3, 4).Equal(tr.MulTuple(p)) {
		t.Errorf("Shearing X to Y went wrong.")
	}
}

func TestShearingXtoZ(t *testing.T) {
	p := tuple.CreatePoint(2, 3, 4)
	tr := Shearing(0, 1, 0, 0, 0, 0)
	if !tuple.CreatePoint(6, 3, 4).Equal(tr.MulTuple(p)) {
		t.Errorf("Shearing X to Z went wrong.")
	}
}

func TestShearingYtoX(t *testing.T) {
	p := tuple.CreatePoint(2, 3, 4)
	tr := Shearing(0, 0, 1, 0, 0, 0)
	if !tuple.CreatePoint(2, 5, 4).Equal(tr.MulTuple(p)) {
		t.Errorf("Shearing Y to X went wrong.")
	}
}

func TestShearingYtoZ(t *testing.T) {
	p := tuple.CreatePoint(2, 3, 4)
	tr := Shearing(0, 0, 0, 1, 0, 0)
	if !tuple.CreatePoint(2, 7, 4).Equal(tr.MulTuple(p)) {
		t.Errorf("Shearing Y to Z went wrong.")
	}
}

func TestShearingZtoX(t *testing.T) {
	p := tuple.CreatePoint(2, 3, 4)
	tr := Shearing(0, 0, 0, 0, 1, 0)
	if !tuple.CreatePoint(2, 3, 6).Equal(tr.MulTuple(p)) {
		t.Errorf("Shearing Z to X went wrong.")
	}
}

func TestShearingZtoY(t *testing.T) {
	p := tuple.CreatePoint(2, 3, 4)
	tr := Shearing(0, 0, 0, 0, 0, 1)
	if !tuple.CreatePoint(2, 3, 7).Equal(tr.MulTuple(p)) {
		t.Errorf("Shearing Z to Y went wrong.")
	}
}

func TestTransforationFolding(t *testing.T) {
	p := tuple.CreatePoint(1, 0, 1)
	a := RotationX(math.Pi / 2)
	b := Scaling(5, 5, 5)
	c := Translation(10, 5, 7)
	p2 := a.MulTuple(p)
	if !tuple.CreatePoint(1, -1, 0).Equal(p2) {
		t.Errorf("A rotation wrong.")
	}
	p3 := b.MulTuple(p2)
	if !tuple.CreatePoint(5, -5, 0).Equal(p3) {
		t.Errorf("B scaling wrong.")
	}
	p4 := c.MulTuple(p3)
	if !tuple.CreatePoint(15, 0, 7).Equal(p4) {
		t.Errorf("C translation wrong.")
	}
	tr := c.Mul(b.Mul(a))
	p5 := tr.MulTuple(p)
	if !tuple.CreatePoint(15, 0, 7).Equal(p5) {
		t.Errorf("T transformation wrong.")
	}
}
