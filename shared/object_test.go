package shared

import (
	"testing"

	"github.com/go-gl/mathgl/mgl32"
	"github.com/stretchr/testify/assert"
)

// Testing constants
var (
	upY Object = Object{
		Name:     "Up by 1j",
		Position: mgl32.Vec3{0, 1, 0},
		Scale:    mgl32.Vec3{1, 1, 1},
	}

	vertex mgl32.Vec4 = mgl32.Vec4{.5, .5, 0, 1}
)

func TestObjectLocateTransformMatrix(t *testing.T) {
	transform := upY.Locate()
	position := transform.Mul4x1(vertex)

	assert.Equal(t, mgl32.Vec4{.5, 1.5, 0, 1}, position)
}
