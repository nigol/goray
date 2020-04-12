package light

import (
    "testing"
    "goray/canvas"
    "goray/tuple"
)

func TestLightPositionIntensity(t *testing.T) {
    intensity := canvas.Color{1, 1, 1}
    position := tuple.CreatePoint(0, 0, 0)
    light := PointLight{position, intensity}
    if light.Position != position || light.Intensity != intensity {
        t.Errorf("PointLight is wrong.")
    }
}