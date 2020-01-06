package matrix

import (
	"goray/tuple"
	"testing"
)

func TestCreateMatrix4x4(t *testing.T) {
	m := Matrix{4, 4,
		[][]float64{
			{1, 2, 3, 4},
			{5.5, 6.5, 7.5, 8.5},
			{9, 10, 11, 12},
			{13.5, 14.5, 15.5, 16.5},
		}}
	comp := m.D[0][0] == 1 && m.D[0][3] == 4 && m.D[1][0] == 5.5 &&
		m.D[1][2] == 7.5 && m.D[2][2] == 11 && m.D[3][0] == 13.5 &&
		m.D[3][2] == 15.5
	if !comp {
		t.Errorf("Matrix definition is wrong.")
	}
}

func TestCreateMatrix2x2(t *testing.T) {
	m := Matrix{2, 2,
		[][]float64{
			{-3, 5},
			{1, -2},
		}}
	comp := m.D[0][0] == -3 && m.D[0][1] == 5 && m.D[1][0] == 1 &&
		m.D[1][1] == -2
	if !comp {
		t.Errorf("Matrix definition is wrong.")
	}
}

func TestCreateMatrix3x3(t *testing.T) {
	m := Matrix{3, 3,
		[][]float64{
			{-3, 5, 0},
			{1, -2, -7},
			{0, 1, 1},
		}}
	comp := m.D[0][0] == -3 && m.D[1][1] == -2 && m.D[2][2] == 1
	if !comp {
		t.Errorf("Matrix definition is wrong.")
	}
}

func TestMatrixEqual(t *testing.T) {
	m1 := Matrix{4, 4,
		[][]float64{
			{1, 2, 3, 4},
			{5, 6, 7, 8},
			{9, 8, 7, 6},
			{5, 4, 3, 2},
		}}
	m2 := Matrix{4, 4,
		[][]float64{
			{1, 2, 3, 4},
			{5, 6, 7, 8},
			{9, 8, 7, 6},
			{5, 4, 3, 2},
		}}
	m3 := Matrix{4, 4,
		[][]float64{
			{2, 3, 4, 5},
			{6, 7, 8, 9},
			{8, 7, 6, 5},
			{4, 3, 2, 1},
		}}
	if !m1.Equal(m2) {
		t.Errorf("Matrices should be equal.")
	}
	if m1.Equal(m3) {
		t.Errorf("Matrices should not be equal.")
	}
}

func TestMatrixMul(t *testing.T) {
	m1 := Matrix{4, 4,
		[][]float64{
			{1, 2, 3, 4},
			{5, 6, 7, 8},
			{9, 8, 7, 6},
			{5, 4, 3, 2},
		}}
	m2 := Matrix{4, 4,
		[][]float64{
			{-2, 1, 2, 3},
			{3, 2, 1, -1},
			{4, 3, 6, 5},
			{1, 2, 7, 8},
		}}
	m3 := Matrix{4, 4,
		[][]float64{
			{20, 22, 50, 48},
			{44, 54, 114, 108},
			{40, 58, 110, 102},
			{16, 26, 46, 42},
		}}
	m4 := m1.Mul(m2)
	if !m3.Equal(m4) {
		t.Errorf("Matrices should be equal.")
	}
}

func TestMatrixMulTuple(t *testing.T) {
	m1 := Matrix{4, 4,
		[][]float64{
			{1, 2, 3, 4},
			{2, 4, 4, 2},
			{8, 6, 4, 1},
			{0, 0, 0, 1},
		}}
	t1 := tuple.CreateTuple(1, 2, 3, 1)
	t2 := tuple.CreateTuple(18, 24, 33, 1)
	t3 := m1.MulTuple(t1)
	if !t3.Equal(t2) {
		t.Errorf("Tuples should be equal.")
	}
}

func TestMatrixMulIdentity(t *testing.T) {
	m1 := Matrix{4, 4,
		[][]float64{
			{0, 1, 2, 4},
			{1, 2, 4, 8},
			{2, 4, 8, 16},
			{4, 8, 16, 32},
		}}
	m2 := Matrix{4, 4,
		[][]float64{
			{1, 0, 0, 0},
			{0, 1, 0, 0},
			{0, 0, 1, 0},
			{0, 0, 0, 1},
		}}
	m3 := m1.Mul(m2)
	if !m3.Equal(m1) {
		t.Errorf("Matrices should be equal.")
	}
}

func TestMatrixMulTupleIdentity(t *testing.T) {
	t1 := tuple.CreateTuple(1, 2, 3, 4)
	m1 := Matrix{4, 4,
		[][]float64{
			{1, 0, 0, 0},
			{0, 1, 0, 0},
			{0, 0, 1, 0},
			{0, 0, 0, 1},
		}}
	t2 := m1.MulTuple(t1)
	if !t1.Equal(t2) {
		t.Errorf("Tuples should be equal.")
	}
}

func TestMatrixTranspose(t *testing.T) {
	m1 := Matrix{4, 4,
		[][]float64{
			{0, 9, 3, 0},
			{9, 8, 0, 8},
			{1, 8, 5, 3},
			{0, 0, 5, 8},
		}}
	m2 := Matrix{4, 4,
		[][]float64{
			{0, 9, 1, 0},
			{9, 8, 8, 0},
			{3, 0, 5, 5},
			{0, 8, 3, 8},
		}}
	m3 := m1.Transpose()
	if !m3.Equal(m2) {
		t.Errorf("Matrices should be equal.")
	}
}

func TestMatrixTransposeIdentity(t *testing.T) {
	m1 := Matrix{4, 4,
		[][]float64{
			{1, 0, 0, 0},
			{0, 1, 0, 0},
			{0, 0, 1, 0},
			{0, 0, 0, 1},
		}}
	m3 := m1.Transpose()
	if !m3.Equal(m1) {
		t.Errorf("Matrices should be equal.")
	}
}

func TestMatrixDeterminant2x2(t *testing.T) {
	m1 := Matrix{2, 2,
		[][]float64{
			{1, 5},
			{-3, 2},
		}}
	if m1.Determinant2x2() != 17 {
		t.Errorf("2x2 determinant should be 17.")
	}
}
