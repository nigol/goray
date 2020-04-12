package material

import (
	"goray/canvas"
)

type Material struct {
	Color     canvas.Color
	Ambient   float64
	Diffuse   float64
	Specular  float64
	Shininess float64
}

func CreateMaterial() Material {
	return Material{canvas.Color{1, 1, 1}, 0.1, 0.9, 0.9, 200}
}
