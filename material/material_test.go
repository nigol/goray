package material

import (
	"goray/canvas"
	"goray/light"
	"goray/tuple"
	"math"
	"testing"
)

func TestDefaultMaterial(t *testing.T) {
	m := CreateMaterial()
	if !(m.Color.Equal(canvas.Color{1, 1, 1}) && m.Ambient == 0.1 && m.Diffuse == 0.9 && m.Specular == 0.9 && m.Shininess == 200) {
		t.Errorf("Default material is wrong.")
	}
}

func TestEqual(t *testing.T) {
	m1 := CreateMaterial()
	m2 := CreateMaterial()
	if !m1.Equal(m2) {
		t.Errorf("Materials should be the same.")
	}
	m2.Shininess = 0.5
	if m1.Equal(m2) {
		t.Errorf("Materials should be different.")
	}
}

func TestLightingEyeBetweenLightSurface(t *testing.T) {
	m := CreateMaterial()
	position := tuple.CreatePoint(0, 0, 0)
	eyev := tuple.CreateVector(0, 0, -1)
	normalv := tuple.CreateVector(0, 0, -1)
	light := light.PointLight{tuple.CreatePoint(0, 0, -10), canvas.Color{1, 1, 1}}
	result := m.Lighting(light, position, eyev, normalv)
	if !result.Equal(canvas.Color{1.9, 1.9, 1.9}) {
		t.Errorf("Lighting with the eye betweeen the light and the surface is wrong.")
	}
}

func TestLightingEyeBetweenLightSurface45(t *testing.T) {
	m := CreateMaterial()
	position := tuple.CreatePoint(0, 0, 0)
	eyev := tuple.CreateVector(0, math.Sqrt(2)/2, -1*math.Sqrt(2)/2)
	normalv := tuple.CreateVector(0, 0, -1)
	light := light.PointLight{tuple.CreatePoint(0, 0, -10), canvas.Color{1, 1, 1}}
	result := m.Lighting(light, position, eyev, normalv)
	if !result.Equal(canvas.Color{1.0, 1.0, 1.0}) {
		t.Errorf("Lighting with the eye betweeen the light and the surface, eye at 45, is wrong.")
	}
}

func TestLightingEyeOppositeSurface45(t *testing.T) {
	m := CreateMaterial()
	position := tuple.CreatePoint(0, 0, 0)
	eyev := tuple.CreateVector(0, 0, -1)
	normalv := tuple.CreateVector(0, 0, -1)
	light := light.PointLight{tuple.CreatePoint(0, 10, -10), canvas.Color{1, 1, 1}}
	result := m.Lighting(light, position, eyev, normalv)
	if !result.Equal(canvas.Color{0.7364, 0.7364, 0.7364}) {
		t.Errorf("Lighting with eye opposite surface, light at 45, is wrong.")
	}
}

func TestLightingEyeInPathReflection(t *testing.T) {
	m := CreateMaterial()
	position := tuple.CreatePoint(0, 0, 0)
	eyev := tuple.CreateVector(0, -1*math.Sqrt(2)/2, -1*math.Sqrt(2)/2)
	normalv := tuple.CreateVector(0, 0, -1)
	light := light.PointLight{tuple.CreatePoint(0, 10, -10), canvas.Color{1, 1, 1}}
	result := m.Lighting(light, position, eyev, normalv)
	if !result.Equal(canvas.Color{1.6364, 1.6364, 1.6364}) {
		t.Errorf("Lighting with eye in the path of the reflection vector is wrong.")
	}
}

func TestLightingLightBehindSurface(t *testing.T) {
	m := CreateMaterial()
	position := tuple.CreatePoint(0, 0, 0)
	eyev := tuple.CreateVector(0, 0, -1)
	normalv := tuple.CreateVector(0, 0, -1)
	light := light.PointLight{tuple.CreatePoint(0, 0, 10), canvas.Color{1, 1, 1}}
	result := m.Lighting(light, position, eyev, normalv)
	if !result.Equal(canvas.Color{0.1, 0.1, 0.1}) {
		t.Errorf("Lighting with the light behind surface is wrong.")
	}
}
