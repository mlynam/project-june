package shared

import (
	"github.com/go-gl/mathgl/mgl32"
)

// Object is the base game engine object type
type Object struct {
	Position mgl32.Vec3
	Scale    mgl32.Vec3
	Rotation mgl32.Vec4
}
