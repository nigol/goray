package main

import (
	"math"
	"testing"
)

func TestCreateTuple(t *testing.T) {
	// Test for point w = 1.0
	tup := CreateTuple(4.3, -4.2, 3.1, 1.0)
	comp := tup.x == 4.3 && tup.y == -4.2 && tup.z == 3.1 && tup.w == 1.0
	if !comp {
		t.Errorf("Tuple is not a point")
	}
	// Test for vector w = 0.0
	tup = CreateTuple(4.3, -4.2, 3.1, 0.0)
	comp = tup.x == 4.3 && tup.y == -4.2 && tup.z == 3.1 && tup.w == 0.0
	if !comp {
		t.Errorf("Tuple is not a vector")
	}
}

func TestCreatePoint(t *testing.T) {
	tup := CreatePoint(4.3, -4.2, 3.1)
	comp := tup.x == 4.3 && tup.y == -4.2 && tup.z == 3.1 && tup.w == 1.0
	if !comp {
		t.Errorf("Tuple is not a point")
	}
}

func TestCreateVector(t *testing.T) {
	tup := CreateVector(4.3, -4.2, 3.1)
	comp := tup.x == 4.3 && tup.y == -4.2 && tup.z == 3.1 && tup.w == 0.0
	if !comp {
		t.Errorf("Tuple is not a vector")
	}
}

func TestEqual(t *testing.T) {
	t1 := CreateTuple(1, 2, 3, 4)
	t2 := CreateTuple(1, 2, 3, 4)
	if !t1.Equal(t2) {
		t.Errorf("Tuples should be equal")
	}
	t1 = CreateTuple(1, 2, 3, 4)
	t2 = CreateTuple(12, 2, 3, 4)
	if t1.Equal(t2) {
		t.Errorf("Tuples should be different")
	}
}

func TestAddTuple(t *testing.T) {
	t1 := CreateTuple(3, -2, 5, 1)
	t2 := CreateTuple(-2, 3, 1, 0)
	t3 := CreateTuple(1, 1, 6, 1)
	t4 := t1.Add(t2)
	if !t3.Equal(t4) {
		t.Errorf("Tuples should be equal: %f, %f, %f, %f", t4.x, t4.y, t4.z, t4.w)
	}
}

func TestSubPointPoint(t *testing.T) {
	t1 := CreatePoint(3, 2, 1)
	t2 := CreatePoint(5, 6, 7)
	t3 := CreateVector(-2, -4, -6)
	t4 := t1.Sub(t2)
	if !t3.Equal(t4) {
		t.Errorf("Tuples should be equal: %f, %f, %f, %f", t4.x, t4.y, t4.z, t4.w)
	}
}

func TestSubPointVector(t *testing.T) {
	t1 := CreatePoint(3, 2, 1)
	t2 := CreateVector(5, 6, 7)
	t3 := CreatePoint(-2, -4, -6)
	t4 := t1.Sub(t2)
	if !t3.Equal(t4) {
		t.Errorf("Tuples should be equal: %f, %f, %f, %f", t4.x, t4.y, t4.z, t4.w)
	}
}

func TestSubVectorVector(t *testing.T) {
	t1 := CreateVector(3, 2, 1)
	t2 := CreateVector(5, 6, 7)
	t3 := CreateVector(-2, -4, -6)
	t4 := t1.Sub(t2)
	if !t3.Equal(t4) {
		t.Errorf("Tuples should be equal: %f, %f, %f, %f", t4.x, t4.y, t4.z, t4.w)
	}
}

func TestNegateTuple(t *testing.T) {
	t1 := CreateTuple(1, -2, 3, -4)
	t2 := CreateTuple(-1, 2, -3, 4)
	t3 := t1.Negate()
	if !t2.Equal(t3) {
		t.Errorf("Tuples should be equal: %f, %f, %f, %f", t3.x, t3.y, t3.z, t3.w)
	}
}

func TestMultipleScalar(t *testing.T) {
	t1 := CreateTuple(1, -2, 3, -4)
	t2 := CreateTuple(3.5, -7, 10.5, -14)
	t3 := t1.ScalarMul(3.5)
	if !t2.Equal(t3) {
		t.Errorf("Tuples should be equal: %f, %f, %f, %f", t3.x, t3.y, t3.z, t3.w)
	}
}

func TestMultipleFraction(t *testing.T) {
	t1 := CreateTuple(1, -2, 3, -4)
	t2 := CreateTuple(0.5, -1, 1.5, -2)
	t3 := t1.ScalarMul(0.5)
	if !t2.Equal(t3) {
		t.Errorf("Tuples should be equal: %f, %f, %f, %f", t3.x, t3.y, t3.z, t3.w)
	}
}

func TestDivisionScalar(t *testing.T) {
	t1 := CreateTuple(1, -2, 3, -4)
	t2 := CreateTuple(0.5, -1, 1.5, -2)
	t3 := t1.ScalarDiv(2)
	if !t2.Equal(t3) {
		t.Errorf("Tuples should be equal: %f, %f, %f, %f", t3.x, t3.y, t3.z, t3.w)
	}
}

func TestMagnitude(t *testing.T) {
	t1 := CreateVector(1, 0, 0)
	t2 := CreateVector(0, 1, 0)
	t3 := CreateVector(0, 0, 1)
	t4 := CreateVector(1, 2, 3)
	t5 := CreateVector(-1, -2, -3)
	if t1.Magnitude() != 1.0 {
		t.Errorf("Magnitude should be 1.0, given: %f", t1.Magnitude())
	}
	if t2.Magnitude() != 1.0 {
		t.Errorf("Magnitude should be 1.0, given: %f", t2.Magnitude())
	}
	if t3.Magnitude() != 1.0 {
		t.Errorf("Magnitude should be 1.0, given: %f", t3.Magnitude())
	}
	if t4.Magnitude() != math.Sqrt(14) {
		t.Errorf("Magnitude should be %f, given: %f", math.Sqrt(14), t4.Magnitude())
	}
	if t5.Magnitude() != math.Sqrt(14) {
		t.Errorf("Magnitude should be %f, given: %f", math.Sqrt(14), t5.Magnitude())
	}
}

func TestNormalize(t *testing.T) {
	t1 := CreateVector(4, 0, 0)
	t2 := CreateVector(1, 0, 0)
	t3 := t1.Normalize()
	if !t2.Equal(t3) {
		t.Errorf("Tuples should be equal: %f, %f, %f, %f", t3.x, t3.y, t3.z, t3.w)
	}
	t1 = CreateVector(1, 2, 3)
	t2 = CreateVector(0.26726, 0.53452, 0.80178)
	t3 = t1.Normalize()
	if !t2.Equal(t3) {
		t.Errorf("Tuples should be equal: %f, %f, %f, %f", t3.x, t3.y, t3.z, t3.w)
	}
	if t3.Magnitude() != 1.0 {
		t.Errorf("Magnitude should be %f, given: %f", 1.0, t3.Magnitude())
	}
}

// Dot (scalar) product
func TestDot(t *testing.T) {
	t1 := CreateVector(1, 2, 3)
	t2 := CreateVector(2, 3, 4)
	if t1.Dot(t2) != 20.0 {
		t.Errorf("Dot product should be %f, given: %f", 20.0, t1.Dot(t2))
	}
}

// Cross product - defined only for vector
func TestCross(t *testing.T) {
	t1 := CreateVector(1, 2, 3)
	t2 := CreateVector(2, 3, 4)
	t3 := t1.Cross(t2)
	t4 := t2.Cross(t1)
	t5 := CreateVector(-1, 2, -1)
	t6 := CreateVector(1, -2, 1)
	if !t3.Equal(t5) {
		t.Errorf("Tuples should be equal: %f, %f, %f, %f", t3.x, t3.y, t3.z, t3.w)
	}
	if !t4.Equal(t6) {
		t.Errorf("Tuples should be equal: %f, %f, %f, %f", t4.x, t4.y, t4.z, t4.w)
	}
}
