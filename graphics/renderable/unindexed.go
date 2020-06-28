package renderable

import (
	"github.com/go-gl/gl/v4.1-core/gl"
	"github.com/mlynam/project-june/graphics"
	"github.com/qmuntal/gltf"
)

// Unindexed renderables use simple vertex arrays to render data
type Unindexed struct {
	VertexBufferID   uint32
	VertexArrayID    uint32
	VertexByteOffset uint32
	VertexByteLength uint32
	VertexByteSize   uint32
}

// Render renders an unindexed renderable
func (m Unindexed) Render() {
	gl.BindVertexArray(m.VertexArrayID)
	gl.DrawArrays(gl.TRIANGLES, 0, int32(m.VertexByteLength/m.VertexByteSize))
}

// PrimitiveToUnindexed converts a gltf primitive into a renderable
func PrimitiveToUnindexed(p *gltf.Primitive, doc *gltf.Document) (graphics.Render, bool) {
	unindexed := Unindexed{}
	accessorID := p.Attributes["POSITION"]
	accessor := doc.Accessors[accessorID]
	view := doc.BufferViews[*accessor.BufferView]
	buffer := doc.Buffers[view.Buffer]

	if accessor.ComponentType != gltf.ComponentFloat {
		return unindexed, false
	}

	var vbo uint32
	gl.GenBuffers(1, &vbo)
	gl.BindBuffer(gl.ARRAY_BUFFER, vbo)
	gl.BufferData(gl.ARRAY_BUFFER, int(view.ByteLength), gl.Ptr(buffer.Data), gl.STATIC_DRAW)
	gl.BindBuffer(gl.ARRAY_BUFFER, 0)

	var vao uint32
	gl.GenVertexArrays(1, &vao)
	gl.BindVertexArray(vao)
	gl.EnableVertexArrayAttrib(vao, 0)
	gl.BindBuffer(gl.ARRAY_BUFFER, vbo)
	gl.VertexAttribPointer(0, 3, gl.FLOAT, false, 0, nil)

	unindexed.VertexBufferID = vbo
	unindexed.VertexArrayID = vao
	unindexed.VertexByteLength = view.ByteLength
	unindexed.VertexByteOffset = view.ByteOffset
	unindexed.VertexByteSize = 12

	err := gl.GetError()
	if err != gl.NO_ERROR {
	} else {
		return unindexed, true
	}

	return unindexed, false
}
