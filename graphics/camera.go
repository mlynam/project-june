package graphics

import (
	"github.com/go-gl/gl/v4.1-core/gl"
	"github.com/go-gl/mathgl/mgl32"
	"github.com/mlynam/project-june/engine"
)

// Camera provides data required to setup a scene camera
type Camera struct {
	engine.Object
	FieldOfView float32
	ZFar        float32
	ZNear       float32
	AspectRatio float32
}

// SetTransformations sets the `camera` and `projection` uniform mat4
//  values in the shader program
func (c *Camera) SetTransformations(program uint32) {
	index := gl.GetUniformLocation(program, gl.Str("camera\x00"))
	if index > -1 {
		camera := c.View()
		gl.UniformMatrix4fv(index, 1, false, &camera[0])
	}

	index = gl.GetUniformLocation(program, gl.Str("projection\x00"))
	if index > -1 {
		projection := c.Projection()
		gl.UniformMatrix4fv(index, 1, false, &projection[0])
	}
}

// View returns transformation matrix from the point of view of the camera
func (c *Camera) View() mgl32.Mat4 {
	eye := c.Position()
	return mgl32.LookAtV(eye, mgl32.Vec3{}, mgl32.Vec3{0, 1, 0})
}

// Projection returns the projection matrix generated by the camera
func (c *Camera) Projection() mgl32.Mat4 {
	return mgl32.Perspective(c.FieldOfView, c.AspectRatio, c.ZNear, c.ZFar)
}
