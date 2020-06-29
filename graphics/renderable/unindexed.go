package renderable

import (
	"github.com/go-gl/gl/v4.6-core/gl"
	"github.com/mlynam/project-june/graphics"
	"github.com/qmuntal/gltf"
)

// Unindexed renderables use simple vertex arrays to render data
type Unindexed struct {
	VertexBufferID uint32
	VertexArrayID  uint32
	VertexCount    uint32
}

// Render renders an unindexed renderable
func (m Unindexed) Render() {
	gl.BindVertexArray(m.VertexArrayID)
	gl.DrawArrays(gl.TRIANGLES, 0, int32(m.VertexCount))
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

	length := int(view.ByteLength)
	slice := buffer.Data[view.ByteOffset : view.ByteOffset+view.ByteLength]

	var vbo uint32
	gl.GenBuffers(1, &vbo)
	gl.BindBuffer(gl.ARRAY_BUFFER, vbo)
	gl.BufferData(gl.ARRAY_BUFFER, length, gl.Ptr(slice), gl.STATIC_DRAW)

	var vao uint32
	gl.GenVertexArrays(1, &vao)
	gl.BindVertexArray(vao)
	gl.EnableVertexArrayAttrib(vao, 0)
	gl.BindBuffer(gl.ARRAY_BUFFER, vbo)
	gl.VertexAttribPointer(0, 3, gl.FLOAT, accessor.Normalized, int32(view.ByteStride), nil)

	unindexed.VertexBufferID = vbo
	unindexed.VertexArrayID = vao
	unindexed.VertexCount = accessor.Count

	err := gl.GetError()
	if err != gl.NO_ERROR {
	} else {
		return unindexed, true
	}

	return unindexed, false
}
