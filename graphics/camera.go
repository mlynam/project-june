package graphics

import (
	"github.com/go-gl/gl/v4.1-core/gl"
	"github.com/go-gl/mathgl/mgl32"
	"github.com/mlynam/project-june/shared"
)

// Camera provides data required to setup a scene camera
type Camera struct {
	shared.Object
	FieldOfView float32
	ZFar        float32
	ZNear       float32
}

// SetUniformCamera sets the `camera` input in the shader program
func (c *Camera) SetUniformCamera(program uint32) {
	location := gl.GetUniformLocation(program, gl.Str("camera"))
	rotation := c.Rotation.Quat().Mat4()
	position := mgl32.Mat4FromRows(mgl32.Vec4{1}, mgl32.Vec4{0, 1}, mgl32.Vec4{0, 0, 1}, c.Position.Vec4(1))
	camera := rotation.Mul4(position)

	gl.UniformMatrix4fv(location, 1, false, &camera[0])
}
