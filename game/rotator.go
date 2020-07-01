package game

import (
	"github.com/go-gl/mathgl/mgl32"
	"github.com/mlynam/project-june/shared"
)

// Rotator rotates on update
type Rotator struct {
	*shared.Object
	angle float64
}

// Update the rotating cube
func (r *Rotator) Update(c *shared.Context) {
	r.angle += c.TimeDelta()
	rotation := mgl32.QuatRotate(float32(r.angle), mgl32.Vec3{0, 1, 1}.Normalize())
	r.Rotation = rotation
}
