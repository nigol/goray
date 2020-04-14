package main

import (
	"fmt"
	"goray/canvas"
	"goray/light"
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
	object.Material.Color = canvas.Color{0.2, 1, 0.2}
	light := light.PointLight{tuple.CreatePoint(-10, 10, -10), canvas.Color{1, 1, 1}}
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
				point := ray.Position(xs.Hit().T)
				normal := xs.Hit().Object.NormalAt(point)
				eye := ray.Direction.ScalarMul(-1)
				col = xs.Hit().Object.GetMaterial().Lighting(light, point, eye, normal)
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
