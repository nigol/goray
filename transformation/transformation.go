package transformation

import (
	"goray/matrix"
	"math"
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

func RotationX(r float64) matrix.Matrix {
	return matrix.Matrix{4, 4,
		[][]float64{
			{1, 0, 0, 0},
			{0, math.Cos(r), -1 * math.Sin(r), 0},
			{0, math.Sin(r), math.Cos(r), 0},
			{0, 0, 0, 1},
		}}
}

func RotationY(r float64) matrix.Matrix {
	return matrix.Matrix{4, 4,
		[][]float64{
			{math.Cos(r), 0, math.Sin(r), 0},
			{0, 1, 0, 0},
			{-1 * math.Sin(r), 0, math.Cos(r), 0},
			{0, 0, 0, 1},
		}}
}
