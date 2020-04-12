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
