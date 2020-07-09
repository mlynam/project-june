package game

import (
	"github.com/go-gl/mathgl/mgl32"
	"github.com/mlynam/project-june/engine"
)

// Rotator rotates on update
type Rotator struct {
	*Object
	angle float64
}

var (
	unitY = mgl32.Vec3{0, 1, 0}.Normalize()
)

// Update the rotating cube
func (r *Rotator) Update(c *engine.Context) {
	r.angle += c.Delta()
	rotation := mgl32.QuatRotate(float32(r.angle), unitY)
	r.rotation = rotation
}
