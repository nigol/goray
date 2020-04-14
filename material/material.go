package material

import (
	"goray/canvas"
	"goray/light"
	"goray/tuple"
	"math"
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

func (m Material) Lighting(light light.PointLight, point, eyev, normalv tuple.Tuple) canvas.Color {
	effectiveColor := m.Color.Product(light.Intensity)
	lightv := light.Position.Sub(point).Normalize()
	ambient := effectiveColor.ScalarMul(m.Ambient)

	// lightDotNormal = cosine of the angle between light vector and normal vector.
	// If negative, light is on the other side of the surface.
	lightDotNormal := lightv.Dot(normalv)
	diffuse := canvas.Color{0, 0, 0}
	specular := canvas.Color{0, 0, 0}
	if lightDotNormal > 0 {
		diffuse = effectiveColor.ScalarMul(m.Diffuse * lightDotNormal)

		// reflectDotEye = cosine of the angle between reflection vector and eye vector.
		// If negative, light reflects away.
		reflectv := lightv.Reflect(normalv)
		reflectDotEye := reflectv.Dot(eyev)
		specular = canvas.Color{0, 0, 0}
		if reflectDotEye > 0 {
			factor := math.Pow(reflectDotEye, m.Shininess)
			specular = light.Intensity.ScalarMul(m.Specular * factor)
		}
	}

	return ambient.Add(diffuse.Add(specular))
}
