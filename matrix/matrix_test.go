package matrix

import (
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
