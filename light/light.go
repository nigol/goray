package light

import (
	"goray/canvas"
	"goray/tuple"
)

type PointLight struct {
	Position tuple.Tuple
	Intensity canvas.Color	
}