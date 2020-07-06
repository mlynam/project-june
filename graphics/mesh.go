package graphics

import (
	"unsafe"

	"github.com/go-gl/gl/v4.1-core/gl"
	"github.com/mlynam/project-june/engine"
	"github.com/mlynam/project-june/graphics/vertex"
)

// Mesh tracks vertex array data loaded into the graphics card memory
type Mesh struct {
	vertices []vertex.Vertex
	index    []uint

	// vao, vbo, and ibo are the vertex array object, vertex buffer object, and
	// index buffer object in graphics memory, respectively
	vao, vbo, ibo uint32

	count int32

	// world is the locatable position of this mesh
	world engine.Locatable

	attributes []vertex.Attribute
}

// New mesh object
func New(vertices []vertex.Vertex, index []uint, world engine.Locatable) *Mesh {
	return &Mesh{
		vertices: vertices,
		index:    index,
		world:    world,
	}
}

// AddAttribute to this mesh
func (m *Mesh) AddAttribute(a vertex.Attribute) {
	m.attributes = append(m.attributes, a)
}

// Render the mesh
func (m *Mesh) Render(g engine.Graphics) {
	// The first render attempt will load the mesh into the graphics memory
	if m.vao == 0 {
		m.load(g)
	} else {
		gl.BindVertexArray(m.vao)
		gl.BindBuffer(gl.ELEMENT_ARRAY_BUFFER, m.ibo)
	}

	index, ok := g.Uniform("model")
	if ok {
		transform := m.world.Locate()
		gl.UniformMatrix4fv(index, 1, false, &transform[0])
	}

	// var offset uint32 = 0
	gl.DrawElements(gl.TRIANGLES, m.count, gl.UNSIGNED_INT, nil)

	gl.BindBuffer(gl.ELEMENT_ARRAY_BUFFER, 0)
	gl.BindVertexArray(0)
}

func (m *Mesh) load(g engine.Graphics) {
	gl.GenVertexArrays(1, &m.vao)
	gl.GenBuffers(1, &m.vbo)
	gl.GenBuffers(1, &m.ibo)

	gl.BindVertexArray(m.vao)
	gl.BindBuffer(gl.ARRAY_BUFFER, m.vbo)
	gl.BindBuffer(gl.ELEMENT_ARRAY_BUFFER, m.ibo)

	for _, attr := range m.attributes {
		index, ok := g.Attribute(attr.Name)
		if ok {
			var offset int = attr.Offset
			gl.EnableVertexAttribArray(uint32(index))
			gl.VertexAttribPointer(uint32(index), int32(attr.Size), attr.Xtype, attr.Normalized, int32(attr.Stride), gl.Ptr(&offset))
			g.EnsureSuccessState()
		}
	}

	// Setup vertex data
	size := int(unsafe.Sizeof(m.vertices[0]))
	gl.BufferData(gl.ARRAY_BUFFER, len(m.vertices)*size, gl.Ptr(&m.vertices[0]), gl.STATIC_DRAW)

	// Setup index data
	size = int(unsafe.Sizeof(m.index[0]))
	gl.BufferData(gl.ELEMENT_ARRAY_BUFFER, len(m.index)*size, gl.Ptr(&m.index[0]), gl.STATIC_DRAW)

	// Unbind
	gl.BindBuffer(gl.ARRAY_BUFFER, 0)

	m.count = int32(len(m.index) / 4)
}
