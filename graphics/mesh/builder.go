package mesh

import (
	"github.com/go-gl/gl/v4.1-core/gl"
	"github.com/mlynam/project-june/engine"
)

// Builder provides functions to build a mesh and load it into memory
type Builder struct {
	position         []byte
	bytesPerPosition int
	indices          []byte
	bytesPerIndex    int
	elementCount     int32
}

// SetPositionData sets the position data
func (b *Builder) SetPositionData(byteLength uint32, data []byte) *Builder {
	b.position = data
	b.bytesPerPosition = int(byteLength)
	return b
}

// SetIndexData sets the index data
func (b *Builder) SetIndexData(bytesPerIndex uint32, data []byte) *Builder {
	b.indices = data
	b.bytesPerIndex = int(bytesPerIndex)
	return b
}

// Build the mesh and tie it to a locatable object
func (b *Builder) Build(world engine.Locatable) engine.Renderable {
	indexable := b.isIndexable()
	var mesh engine.Renderable
	var vertexArrayBuffer uint32

	if !indexable {
		gl.GenVertexArrays(1, &vertexArrayBuffer)
		gl.BindVertexArray(vertexArrayBuffer)
	}

	var vertexBuffer uint32
	gl.GenBuffers(1, &vertexBuffer)
	gl.BindBuffer(gl.ARRAY_BUFFER, vertexBuffer)
	gl.BufferData(gl.ARRAY_BUFFER, b.bytesPerPosition, gl.Ptr(&b.position[0]), gl.STATIC_DRAW)

	if indexable {
		var indexBuffer uint32
		gl.GenBuffers(1, &indexBuffer)
		gl.BindBuffer(gl.ELEMENT_ARRAY_BUFFER, indexBuffer)
		gl.BufferData(gl.ELEMENT_ARRAY_BUFFER, b.bytesPerIndex, gl.Ptr(&b.indices[0]), gl.STATIC_DRAW)

		mesh = &Indexed{
			Shared: &Shared{
				world: world,
			},
			vertexBufferObjectID: vertexBuffer,
			indexBufferObjectID:  indexBuffer,
			elementCount:         b.elementCount,
		}
	} else {
		mesh = &Arrayed{
			Shared: &Shared{
				world: world,
			},
			vertexArrayBufferID: vertexArrayBuffer,
			elementCount:        b.elementCount,
		}
	}

	return mesh
}

func (b *Builder) isIndexable() bool {
	return len(b.indices) > 0 && b.bytesPerIndex > 0
}
