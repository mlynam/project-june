package dataloader

import (
	"log"

	"github.com/go-gl/gl/v4.1-core/gl"
	"github.com/mlynam/project-june/graphics"
	"github.com/qmuntal/gltf"
)

// LoadModel loads a collection of primitives into a model that can later be rendered
func LoadModel(m *gltf.Mesh, d *gltf.Document) (*graphics.Model, bool) {
	model := graphics.Model{Meshes: make([]graphics.Mesh, len(m.Primitives))}
	for i, primitive := range m.Primitives {
		var mesh graphics.Mesh
		vertexID, ok := primitive.Attributes["POSITION"]

		if !ok {
			log.Print("WARN gltf failed to provide position data for primitive")
			return nil, false
		}

		var vbo uint32
		gl.GenBuffers(1, &vbo)

		ok = loadBufferData(&vertexID, d, vbo, gl.ARRAY_BUFFER)
		if !ok {
			return nil, false
		}

		indexID := primitive.Indices
		if indexID != nil {
			var ibo uint32
			gl.GenBuffers(1, &ibo)

			ok = loadBufferData(indexID, d, ibo, gl.ELEMENT_ARRAY_BUFFER)
			if !ok {
				return nil, false
			}

			mesh.IndexBufferID = &ibo
		}

		mesh.VertexBufferID = vbo
		mesh.VertexCount = d.Accessors[vertexID].Count

		model.Meshes[i] = mesh
	}

	return &model, true
}

func loadBufferData(accessorID *uint32, d *gltf.Document, bufferID uint32, target uint32) bool {
	accessor := d.Accessors[*accessorID]
	if accessor.BufferView == nil {
		return false
	}

	view := d.BufferViews[*accessor.BufferView]
	buffer := d.Buffers[view.Buffer]
	slice := buffer.Data[view.ByteOffset : view.ByteOffset+view.ByteLength]

	// TODO: detect animated models and set this to dynamic in that case
	var drawType uint32 = gl.STATIC_DRAW

	gl.BindBuffer(target, bufferID)
	gl.BufferData(target, int(view.ByteLength), gl.Ptr(slice), drawType)

	if target == gl.ARRAY_BUFFER {
		var vao uint32
		gl.GenVertexArrays(1, &vao)
		gl.BindVertexArray(vao)
		gl.EnableVertexArrayAttrib(vao, 0)
		gl.VertexAttribPointer(0, 3, gl.FLOAT, accessor.Normalized, int32(view.ByteStride), nil)
	}

	if err := gl.GetError(); err != gl.NO_ERROR {
		log.Printf("WARN OpenGL error: %x", err)
		return false
	}

	return true
}
