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
            for j:= 0; j < m.Cols; j++ {
                result = result && common.Equal(m.D[i][j], o.D[i][j])
            }
        }
    }
    return result
}
