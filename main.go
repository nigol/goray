package main

import (
	"fmt"
	"goray/canvas"
	"goray/object"
	"goray/ray"
	"goray/tuple"
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
	rayOrigin := tuple.CreatePoint(0, 0, -5)
	const wallZ float64 = 10
	const wallSize float64 = 7
	canvasPixels := 100
	pixelSize := wallSize / float64(canvasPixels)
	half := wallSize / 2
	can := canvas.CreateCanvas(canvasPixels, canvasPixels)
	object := object.CreateSphere()
	for y := 0; y < canvasPixels; y++ {
		worldY := half - pixelSize*float64(y)
		for x := 0; x < canvasPixels; x++ {
			worldX := -half + pixelSize*float64(x)
			position := tuple.CreatePoint(worldX, worldY, wallZ)
			ray := ray.Ray{
				rayOrigin,
				position.Sub(rayOrigin).Normalize(),
			}
			xs := object.Intersect(ray)
			if xs.Defined() {
				can.WritePixel(x, y, col)
			}
		}
	}
	n, err := canvasToFile(can)
	if err != nil {
		panic("Error writing to file.")
	}
	fmt.Printf("%d bytes writen.\n", n)
}
