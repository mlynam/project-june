package shader

import "github.com/mlynam/project-june/graphics/gl"

// Program is a graphics pipeline program with shaders attached
type Program struct {
	handle uint32
}

// SetProgram sets the program handle
func (p *Program) SetProgram(handle uint32) {
	p.handle = handle
}

// Attribute tries to get the index for the named attribute
func (p *Program) Attribute(name string) (int32, bool) {
	index := gl.GetAttribLocation(p.handle, gl.Str(name))
	return index, index > -1
}

// Uniform tries to get the index for the named uniform
func (p *Program) Uniform(name string) (int32, bool) {
	index := gl.GetUniformLocation(p.handle, gl.Str(name))
	return index, index > -1
}
