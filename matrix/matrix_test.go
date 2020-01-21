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

func TestMatrixSubmatrix3x3(t *testing.T) {
	m1 := Matrix{3, 3,
		[][]float64{
			{1, 5, 0},
			{-3, 2, 7},
			{0, 6, -3},
		}}
	m2 := Matrix{2, 2,
		[][]float64{
			{-3, 2},
			{0, 6},
		}}
	mt := m1.Submatrix(0, 2)
	if !mt.Equal(m2) {
		t.Errorf("Submatrix 3x3 is incorrect. %f %f %f %f", mt.D[0][0], mt.D[0][1], mt.D[1][0], mt.D[1][1])
	}
}

func TestMatrixSubmatrix4x4(t *testing.T) {
	m1 := Matrix{4, 4,
		[][]float64{
			{-6, 1, 1, 6},
			{-8, 5, 8, 6},
			{-1, 0, 8, 2},
			{-7, 1, -1, 1},
		}}
	m2 := Matrix{3, 3,
		[][]float64{
			{-6, 1, 6},
			{-8, 8, 6},
			{-7, -1, 1},
		}}
	mt := m1.Submatrix(2, 1)
	if !mt.Equal(m2) {
		t.Errorf("Submatrix 4x4 is incorrect. \n %f %f %f\n %f %f %f\n %f %f %f\n", mt.D[0][0], mt.D[0][1], mt.D[0][2], mt.D[1][0], mt.D[1][1], mt.D[1][2], mt.D[2][0], mt.D[2][1], mt.D[2][2])
	}
}

func TestMatrixMinor3x3(t *testing.T) {
	m1 := Matrix{3, 3,
		[][]float64{
			{3, 5, 0},
			{2, -1, -7},
			{6, -1, 5},
		}}
	mt := m1.Submatrix(1, 0)
	det := mt.Determinant2x2()
	if det != 25 || m1.Minor3x3(1, 0) != 25 {
		t.Errorf("Minor of 3x3 is incorrect. %f", m1.Minor3x3(1, 0))
	}
}

func TestMatrixCofactor3x3(t *testing.T) {
	m1 := Matrix{3, 3,
		[][]float64{
			{3, 5, 0},
			{2, -1, -7},
			{6, -1, 5},
		}}
	min1 := m1.Minor3x3(0, 0)
	if min1 != -12 || m1.Cofactor3x3(0, 0) != -12 {
		t.Errorf("Cofactor of 3x3 is incorrect. %f", m1.Cofactor3x3(0, 0))
	}
	min2 := m1.Minor3x3(1, 0)
	if min2 != 25 || m1.Cofactor3x3(1, 0) != -25 {
		t.Errorf("Cofactor of 3x3 is incorrect. %f", m1.Cofactor3x3(1, 0))
	}
}

func TestMatrixDeterminant3x3(t *testing.T) {
	m1 := Matrix{3, 3,
		[][]float64{
			{1, 2, 6},
			{-5, 8, -4},
			{2, 6, 4},
		}}
	co1 := m1.Cofactor3x3(0, 0)
	if co1 != 56 {
		t.Errorf("Cofactor of 3x3 is incorrect. %f", m1.Cofactor3x3(0, 0))
	}
	co2 := m1.Cofactor3x3(0, 1)
	if co2 != 12 {
		t.Errorf("Cofactor of 3x3 is incorrect. %f", m1.Cofactor3x3(0, 1))
	}
	co3 := m1.Cofactor3x3(0, 2)
	if co3 != -46 {
		t.Errorf("Cofactor of 3x3 is incorrect. %f", m1.Cofactor3x3(0, 2))
	}
	d := m1.Determinant3x3()
	if d != -196 {
		t.Errorf("Determinant of 3x3 is incorrect. %f", m1.Determinant3x3())
	}
}

func TestMatrixDeterminant4x4(t *testing.T) {
	m1 := Matrix{4, 4,
		[][]float64{
			{-2, -8, 3, 5},
			{-3, 1, 7, 3},
			{1, 2, -9, 6},
			{-6, 7, 7, -9},
		}}
	co1 := m1.Cofactor4x4(0, 0)
	if co1 != 690 {
		t.Errorf("Cofactor of 4x4 is incorrect. %f", m1.Cofactor4x4(0, 0))
	}
	d := m1.Determinant4x4()
	if d != -4071 {
		t.Errorf("Determinant of 4x4 is incorrect. %f", m1.Determinant4x4())
	}
}

func TestMatrixInvertible4x4(t *testing.T) {
	m1 := Matrix{4, 4,
		[][]float64{
			{6, 4, 4, 4},
			{5, 5, 7, 6},
			{4, -9, 3, -7},
			{9, 1, 7, -6},
		}}
	m2 := Matrix{4, 4,
		[][]float64{
			{-4, 2, -2, -3},
			{9, 6, 2, 6},
			{0, -5, 1, -5},
			{0, 0, 0, 0},
		}}
	d1 := m1.Determinant4x4()
	if d1 != -2120 {
		t.Errorf("Determinant of 4x4 is incorrect. %f", m1.Determinant4x4())
	}
	if !m1.Invertible4x4() {
		t.Errorf("Matrix should be invertible.")
	}
	d2 := m2.Determinant4x4()
	if d2 != 0 {
		t.Errorf("Determinant of 4x4 is incorrect. %f", m2.Determinant4x4())
	}
	if m2.Invertible4x4() {
		t.Errorf("Matrix should be invertible.")
	}
}

func TestMatrixInverse4x4_1(t *testing.T) {
	m1 := Matrix{4, 4,
		[][]float64{
			{-5, 2, 6, -8},
			{1, -5, 1, 8},
			{7, 7, -6, -7},
			{1, -3, 7, 4},
		}}
	m2 := Matrix{4, 4,
		[][]float64{
			{0.21805, 0.45113, 0.24060, -0.04511},
			{-0.80827, -1.45677, -0.44361, 0.52068},
			{-0.07895, -0.22368, -0.05263, 0.19737},
			{-0.52256, -0.81391, -0.30075, 0.30639},
		}}
    m3 := m1.Inverse4x4()
	d := m1.Determinant4x4()
	if d != 532 {
		t.Errorf("Determinant of 4x4 is incorrect. %f", m1.Determinant4x4())
	}
	if m1.Cofactor4x4(2, 3) != -160 {
		t.Errorf("Cofactor of 4x4 is incorrect. %f", m1.Cofactor4x4(2, 3))
	}
	if m3.D[3][2] != -160.0/532.0 {
		t.Errorf("Should be %f.\n%s", -160.0/532.0, m3.String())
	}
	if m1.Cofactor4x4(3, 2) != 105 {
		t.Errorf("Cofactor of 4x4 is incorrect. %f", m1.Cofactor4x4(3, 2))
	}
	if m3.D[2][3] != 105.0/532 {
		t.Errorf("Should be %f.\n%s", 105.0/532.0, m3.String())
	}
	if !m2.Equal(m3) {
		t.Errorf("Should be %s.\n%s", m2.String(), m3.String())
	}
}

func TestMatrixInverse4x4_2(t *testing.T) {
	a := Matrix{4, 4,
		[][]float64{
			{3, -9, 7, 3},
			{3, -8, 2, -9},
			{-4, 4, 4, 1},
			{-6, 5, -1, 1},
		}}
	b := Matrix{4, 4,
		[][]float64{
			{8, 2, 2, 2},
			{3, -1, 7, 0},
			{7, 0, 5, 4},
			{6, -2, 0, 5},
		}}
    c := a.Mul(b)
	if !a.Equal(c.Mul(b.Inverse4x4())) {
		t.Errorf("Should be %s.\n%s", a.String(), c.Mul(b.Inverse4x4()).String())
	}
}
