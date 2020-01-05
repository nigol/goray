package matrix

import (
	"goray/common"
)

type Matrix struct {
	Rows int
	Cols int
	D    [][]float64
}

func (m Matrix) Equal(o Matrix) bool {
	result := m.Rows == o.Rows && m.Cols == m.Cols
	if result {
		for i := 0; i < m.Rows; i++ {
			for j := 0; j < m.Cols; j++ {
				result = result && common.Equal(m.D[i][j], o.D[i][j])
			}
		}
	}
	return result
}

func (m Matrix) Mul(o Matrix) Matrix {
	mr := Matrix{m.Rows, o.Cols,
		[][]float64{
			{0, 0, 0, 0},
			{0, 0, 0, 0},
			{0, 0, 0, 0},
			{0, 0, 0, 0},
		}}
	for i := 0; i < m.Rows; i++ {
		for j := 0; j < o.Cols; j++ {
			mr.D[i][j] = m.D[i][0]*o.D[0][j] +
				m.D[i][1]*o.D[1][j] +
				m.D[i][2]*o.D[2][j] +
				m.D[i][3]*o.D[3][j]
		}
	}
	return mr
}
