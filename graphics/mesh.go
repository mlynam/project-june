package graphics

import (
	"unsafe"

	"github.com/go-gl/gl/v4.1-core/gl"
	"github.com/mlynam/project-june/engine"
	"github.com/mlynam/project-june/graphics/vertex"
)

// Mesh tracks vertex array data loaded into the graphics card memory. It expects
// the basic vertex shader is linked to the current graphics program:
//		- assets/shaders/basic.vert
type Mesh struct {
	vertices []vertex.Vertex
	index    []uint32
	vao      uint32
	location [16]float32

	// world is the locatable position of this mesh
	world engine.Locatable
}

// New mesh object
func New(vertices []vertex.Vertex, index []uint32, world engine.Locatable) *Mesh {
	return &Mesh{
		vertices: vertices,
		index:    index,
		world:    world,
	}
}

// Render the mesh
func (m *Mesh) Render(g engine.Graphics) {
	// The first render attempt will load the mesh into the graphics memory
	if m.vao == 0 {
		m.load(g)
	}

	index, ok := g.Uniform("world")
	if ok {
		m.location = m.world.Locate()
		gl.UniformMatrix4fv(index, 1, false, &m.location[0])
	}

	gl.BindVertexArray(m.vao)
	gl.DrawElements(gl.TRIANGLES, int32(len(m.index)), gl.UNSIGNED_INT, gl.PtrOffset(0))
	gl.BindVertexArray(0)
}

func (m *Mesh) load(g engine.Graphics) {
	vert := vertex.Vertex{}
	size := unsafe.Sizeof(vert)

	gl.GenVertexArrays(1, &m.vao)

	var buffers [2]uint32
	gl.GenBuffers(2, &buffers[0])

	gl.BindVertexArray(m.vao)

	// Setup vertex data
	gl.BindBuffer(gl.ARRAY_BUFFER, buffers[0])
	gl.BufferData(gl.ARRAY_BUFFER, len(m.vertices)*int(size), gl.Ptr(m.vertices), gl.STATIC_DRAW)
	g.EnsureSuccessState()

	// Setup index data
	gl.BindBuffer(gl.ELEMENT_ARRAY_BUFFER, buffers[1])
	gl.BufferData(gl.ELEMENT_ARRAY_BUFFER, len(m.index)*4, gl.Ptr(&m.index[0]), gl.STATIC_DRAW)
	g.EnsureSuccessState()

	gl.EnableVertexAttribArray(0)
	gl.VertexAttribPointer(0, 3, gl.FLOAT, false, int32(size), gl.PtrOffset(vertex.PositionOffset))
	g.EnsureSuccessState()

	gl.EnableVertexAttribArray(1)
	gl.VertexAttribPointer(1, 3, gl.FLOAT, false, int32(size), gl.PtrOffset(vertex.ColorOffset))
	g.EnsureSuccessState()

	gl.BindVertexArray(0)
	gl.BindBuffer(gl.ARRAY_BUFFER, 0)
	gl.BindBuffer(gl.ELEMENT_ARRAY_BUFFER, 0)
}
