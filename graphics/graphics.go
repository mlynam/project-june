package graphics

import "github.com/go-gl/gl/v4.1-core/gl"

// Graphics elements
type Graphics struct {
	program uint32
}

// Program returns the shader program currently attached to the graphics pipeline
func (g *Graphics) Program() uint32 {
	return g.program
}

// SetProgram sets the program handle
func (g *Graphics) SetProgram(handle uint32) {
	g.program = handle
}

// Attribute tries to get the index for the named attribute
func (g *Graphics) Attribute(name string) (int32, bool) {
	index := gl.GetAttribLocation(g.program, gl.Str(name))
	return index, index > -1
}

// Uniform tries to get the index for the named uniform
func (g *Graphics) Uniform(name string) (int32, bool) {
	index := gl.GetUniformLocation(g.program, gl.Str(name))
	return index, index > -1
}
