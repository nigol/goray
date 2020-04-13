package material

import (
	"goray/canvas"
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
