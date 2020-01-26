package main

import (
	"fmt"
	"goray/canvas"
	"goray/transformation"
	"goray/tuple"
	"math"
	"os"
)

func canvasToFile(can canvas.Canvas) (n int, err error) {
	f, e1 := os.Create("out.ppm")
	if e1 != nil {
		return 0, e1
	}
	defer f.Close()
	n, e2 := f.WriteString(can.Ppm())
	return n, e2
}

func main() {
	col := canvas.Color{0.0, 1.0, 0.0}
	can := canvas.CreateCanvas(200, 200)
	step := 2.0 * math.Pi / 12
	fmt.Printf("step: %f\n", step)
	for i := 0; i < 12; i++ {
		p := tuple.CreatePoint(0, 1, 0)
		t := transformation.Translation(80, 80, 0).Mul(transformation.Scaling(80, 80, 0).Mul(transformation.RotationZ(step * float64(i))))
		pt := t.MulTuple(p)
		fmt.Printf("x: %f, y: %f\n", pt.X, pt.Y)
		can.WritePixel(int(pt.X), int(pt.Y), col)
	}
	n, err := canvasToFile(can)
	if err != nil {
		panic("Error writing to file.")
	}
	fmt.Printf("%d bytes writen.\n", n)
}
