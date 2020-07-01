package shared

import (
	"github.com/go-gl/mathgl/mgl32"
)

// Object is the base game engine object type
type Object struct {
	Children []*Object
	Parent   Locatable
	Name     string
	Position mgl32.Vec3
	Scale    mgl32.Vec3
	Rotation mgl32.Quat
}

// Update the object
func (o *Object) Update(c *Context) {
}

// Locate returns the world space transform
func (o *Object) Locate() mgl32.Mat4 {
	parent := mgl32.Ident4()
	if o.Parent != nil {
		parent = parent.Mul4(o.Parent.Locate())
	}

	t, s := o.Position, o.Scale
	scale := mgl32.Diag4(s.Vec4(1))
	rotation := o.Rotation.Mat4()
	translation := mgl32.Mat4{
		1, 0, 0, 0,
		0, 1, 0, 0,
		0, 0, 1, 0,
		t[0], t[1], t[2], 1,
	}

	return scale.Mul4(rotation).Mul4(parent).Mul4(translation)
}
