package game

import (
	"github.com/go-gl/mathgl/mgl32"
	"github.com/mlynam/project-june/engine"
)

// Object is the base game object
type Object struct {
	name     string
	position mgl32.Vec3
	scale    mgl32.Vec3
	rotation mgl32.Quat
	parent   engine.Locatable
	children []*Object

	location [16]float32
}

// NewObject object from the given propeties
func NewObject(name string, position, scale mgl32.Vec3, rotation mgl32.Quat) *Object {
	object := &Object{
		name:     name,
		position: position,
		scale:    scale,
		rotation: rotation,
		location: mgl32.Ident4(),
		parent:   nil,
		children: make([]*Object, 0),
	}

	object.Synchronize()

	return object
}

// Update the object
func (o *Object) Update(c *engine.Context) {
}

// Synchronize the object
func (o *Object) Synchronize() {
	parent := mgl32.Ident4()
	if o.parent != nil {
		parent = parent.Mul4(o.parent.Locate())
	}

	t, s := o.position, o.scale
	scale := mgl32.Diag4(s.Vec4(1))
	rotation := o.rotation.Mat4()
	translation := mgl32.Translate3D(t[0], t[1], t[2])

	o.location = scale.Mul4(rotation).Mul4(parent).Mul4(translation)
}

// Locate the game object in the world
func (o *Object) Locate() [16]float32 {
	return o.location
}

// Position of the object
func (o *Object) Position() [3]float32 {
	return o.position
}

// AddChild to this object
func (o *Object) AddChild(c *Object) {
	c.parent = o
	o.children = append(o.children, c)
}
