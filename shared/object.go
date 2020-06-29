package shared

import (
	"github.com/go-gl/mathgl/mgl32"
)

// Object is the base game engine object type
type Object struct {
	position mgl32.Vec3
	scale    mgl32.Vec3
	rotation mgl32.Vec4
}
