package core

import "github.com/go-gl/mathgl/mgl32"

// Locatable types can be located with a world space transformation matrix
type Locatable interface {
	Locate() mgl32.Mat4
}
