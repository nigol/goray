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

func (m Material) Equal(o Material) bool {
	return m.Color.Equal(o.Color) &&
		m.Ambient == o.Ambient &&
		m.Diffuse == o.Diffuse &&
		m.Specular == o.Specular &&
		m.Shininess == o.Shininess
}
