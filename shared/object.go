package shared

import (
	"fmt"

	"github.com/go-gl/mathgl/mgl32"
)

// Object is the base game engine object type
type Object struct {
	Name     string
	Position mgl32.Vec3
	Scale    mgl32.Vec3
	Rotation mgl32.Vec4
	Children []*Object
	Parent   *Object
}

// Locate returns the identity matrix
func (o *Object) Locate() mgl32.Mat4 {
	t, s := o.Position, o.Scale
	scale := mgl32.Diag4(s.Vec4(1))
	rotation := o.Rotation.Quat().Mat4()
	translation := mgl32.Mat4{
		1, 0, 0, 0,
		0, 1, 0, 0,
		0, 0, 1, 0,
		t[0], t[1], t[2], 1,
	}

	location := scale.Mul4(rotation).Mul4(translation)
	uniform := mgl32.Vec3{s.X(), s.X(), s.X()}

	if uniform != s {
		return location.Inv().Transpose()
	}

	return location
}

// String renders the object as a string
func (o *Object) String() string {
	parent := "<nil>"

	if o.Parent != nil {
		parent = o.Parent.Name
	}

	val := struct {
		Name     string
		Parent   string
		Position mgl32.Vec3
		Scale    mgl32.Vec3
		Rotation mgl32.Vec4
	}{
		o.Name,
		parent,
		o.Position,
		o.Scale,
		o.Rotation,
	}

	return fmt.Sprintf("%+v", val)
}
