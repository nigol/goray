package main

import (
	"fmt"
	"goray/canvas"
	"goray/matrix"
	"goray/tuple"
	"os"
)

type Projectile struct {
	position tuple.Tuple
	velocity tuple.Tuple
}

type Environment struct {
	gravity tuple.Tuple
	wind    tuple.Tuple
}

func tick(env Environment, proj Projectile) Projectile {
	position := proj.position.Add(proj.velocity)
	velocity := proj.velocity.Add(env.gravity).Add(env.wind)
	result := Projectile{
		position,
		velocity,
	}
	return result
}

func matrixPlay() {
	a := matrix.Matrix{4, 4,
		[][]float64{
			{1, 0, 0, 0},
			{0, 1, 0, 0},
			{0, 0, 1, 0},
			{0, 0, 0, 1},
		}}
	fmt.Println(a.Inverse4x4().String())
	b := matrix.Matrix{4, 4,
		[][]float64{
			{-2, -8, 3, 5},
			{-3, 1, 7, 3},
			{1, 2, -9, 6},
			{-6, 7, 7, -9},
		}}
	fmt.Println(b.Mul(b.Inverse4x4()))
	fmt.Println(b.Transpose().Inverse4x4())
	fmt.Println(b.Inverse4x4().Transpose())
	t := tuple.CreateTuple(1, 2, -9, 6)
	fmt.Println(a.MulTuple(t))
}

func main() {
	p := Projectile{
		position: tuple.CreatePoint(0, 1, 0),
		velocity: tuple.CreateVector(1, 1.8, 0).Normalize().ScalarMul(9.25),
	}
	e := Environment{
		gravity: tuple.CreateVector(0, -0.1, 0),
		wind:    tuple.CreateVector(-0.01, 0, 0),
	}
	col := canvas.Color{0.0, 1.0, 0.0}
	can := canvas.CreateCanvas(900, 550)
	for p.position.Y >= 0 {
		p = tick(e, p)
		fmt.Printf("x: %f, y: %f\n", p.position.X, p.position.Y)
		can.WritePixel(int(p.position.X), can.Height-int(p.position.Y), col)
	}
	f, err := os.Create("out.ppm")
	if err != nil {
		panic("Can't create image file.")
	}
	defer f.Close()
	n, err := f.WriteString(can.Ppm())
	if err != nil {
		panic("Can't write to file.")
	}
	fmt.Printf("%d bytes writen.\n", n)
	matrixPlay()
}
