package matrix

import (
	"goray/common"
	"goray/tuple"
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

func (m Matrix) MulTuple(t tuple.Tuple) tuple.Tuple {
	mt := Matrix{4, 1,
		[][]float64{
			{t.X},
			{t.Y},
			{t.Z},
			{t.W},
		}}
	mr := m.Mul(mt)
	return tuple.CreateTuple(mr.D[0][0], mr.D[1][0], mr.D[2][0], mr.D[3][0])
}

func (m Matrix) Transpose() Matrix {
	mr := Matrix{m.Rows, m.Cols,
		[][]float64{
			{0, 0, 0, 0},
			{0, 0, 0, 0},
			{0, 0, 0, 0},
			{0, 0, 0, 0},
		}}
	for i := 0; i < m.Rows; i++ {
		for j := 0; j < m.Cols; j++ {
			mr.D[j][i] = m.D[i][j]
		}
	}
	return mr
}

func (m Matrix) Determinant2x2() float64 {
	return m.D[0][0]*m.D[1][1] - m.D[0][1]*m.D[1][0]
}

func (m Matrix) Submatrix(r int, c int) Matrix {
	mr := Matrix{m.Rows - 1, m.Cols - 1,
		[][]float64{
			{0, 0, 0, 0},
			{0, 0, 0, 0},
			{0, 0, 0, 0},
			{0, 0, 0, 0},
		}}
	for i := 0; i < m.Rows; i++ {
		if i == r {
			continue
		}
		for j := 0; j < m.Cols; j++ {
			if j == c {
				continue
			}
			fr := i
			if i > r {
				fr = fr - 1
			}
			fc := j
			if j > c {
				fc = fc - 1
			}
			mr.D[fr][fc] = m.D[i][j]
		}
	}
	return mr
}

func (m Matrix) Minor3x3(r int, c int) float64 {
	return m.Submatrix(r, c).Determinant2x2()
}

func (m Matrix) Cofactor3x3(r int, c int) float64 {
	result := m.Minor3x3(r, c)
	if (r+c)%2 == 1 {
		result = result * -1
	}
	return result
}

func (m Matrix) Determinant3x3() float64 {
	result := 0.0
	for i := 0; i < 3; i++ {
		result = result + m.D[0][i]*m.Cofactor3x3(0, i)
	}
	return result
}

func (m Matrix) Minor4x4(r int, c int) float64 {
	return m.Submatrix(r, c).Determinant3x3()
}

func (m Matrix) Cofactor4x4(r int, c int) float64 {
	result := m.Minor4x4(r, c)
	if (r+c)%2 == 1 {
		result = result * -1
	}
	return result
}

func (m Matrix) Determinant4x4() float64 {
	result := 0.0
	for i := 0; i < 4; i++ {
		result = result + m.D[0][i]*m.Cofactor4x4(0, i)
	}
	return result
}

func (m Matrix) Invertible4x4() bool {
	return m.Determinant4x4() != 0
}
