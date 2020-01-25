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

func RotationZ(r float64) matrix.Matrix {
	return matrix.Matrix{4, 4,
		[][]float64{
			{math.Cos(r), -1 * math.Sin(r), 0, 0},
			{math.Sin(r), math.Cos(r), 0, 0},
			{0, 0, 1, 0},
			{0, 0, 0, 1},
		}}
}

func Shearing(xy float64, xz float64, yx float64, yz float64, zx float64, zy float64) matrix.Matrix {
	return matrix.Matrix{4, 4,
		[][]float64{
			{1, xy, xz, 0},
			{yx, 1, yz, 0},
			{zx, zy, 1, 0},
			{0, 0, 0, 1},
		}}
}
