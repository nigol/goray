package transformation

import (
	"goray/matrix"
)

func Translation(x float64, y float64, z float64) matrix.Matrix {
	return matrix.Matrix{4, 4,
		[][]float64{
			{1, 0, 0, x},
			{0, 1, 0, y},
			{0, 0, 1, z},
			{0, 0, 0, 1},
		}}
}

func Scaling(x float64, y float64, z float64) matrix.Matrix {
	return matrix.Matrix{4, 4,
		[][]float64{
			{x, 0, 0, 0},
			{0, y, 0, 0},
			{0, 0, z, 0},
			{0, 0, 0, 1},
		}}
}
