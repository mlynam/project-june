package graphics

import (
	"fmt"

	"github.com/go-gl/gl/v4.1-core/gl"
	"github.com/mlynam/project-june/engine"
)

// Graphics elements
type Graphics struct {
	program uint32
	cache   map[string]int32
	scene   engine.Scene
}

// NewGraphics creates a new graphics instance
func NewGraphics(program uint32) *Graphics {
	return &Graphics{
		program: program,
		cache:   make(map[string]int32),
	}
}

// Clear the back buffer
func (g *Graphics) Clear() {
	gl.Clear(gl.COLOR_BUFFER_BIT | gl.DEPTH_BUFFER_BIT)
}

// Program returns the shader program currently attached to the graphics pipeline
func (g *Graphics) Program() uint32 {
	return g.program
}

// SetProgram sets the program handle
func (g *Graphics) SetProgram(handle uint32) {
	g.program = handle
	g.cache = make(map[string]int32)
}

// Attribute tries to get the index for the named attribute
func (g *Graphics) Attribute(name string) (int32, bool) {
	key := fmt.Sprintf("attribute:%v", name)
	cached, ok := g.cache[key]
	if ok {
		return cached, cached > -1
	}

	index := gl.GetAttribLocation(g.program, gl.Str(name+"\x00"))
	g.cache[key] = index
	return index, index > -1
}

// Uniform tries to get the index for the named uniform
func (g *Graphics) Uniform(name string) (int32, bool) {
	key := fmt.Sprintf("uniform:%v", name)
	cached, ok := g.cache[key]
	if ok {
		return cached, cached > -1
	}

	index := gl.GetUniformLocation(g.program, gl.Str(name+"\x00"))
	g.cache[key] = index
	return index, index > -1
}

// EnsureSuccessState checks the error state and panics if it finds an error
func (g *Graphics) EnsureSuccessState() {
	err := gl.GetError()
	if err != gl.NO_ERROR {
		panic(err)
	}
}

// SceneViewProjection from the scene camera
func (g *Graphics) SceneViewProjection() [16]float32 {
	return g.scene.SceneViewProjection()
}

// SetScene to be rendered
func (g *Graphics) SetScene(s engine.Scene) {
	g.scene = s
}
