package dataloader

import (
	"log"

	"github.com/go-gl/gl/v4.1-core/gl"
	"github.com/mlynam/project-june/graphics"
	"github.com/qmuntal/gltf"
)

// LoadModel loads a collection of primitives into a model that can later be rendered
func LoadModel(m *gltf.Mesh, d *gltf.Document, program uint32) (*graphics.Model, bool) {
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

		accessor := d.Accessors[vertexID]
		if accessor.BufferView == nil {
			log.Print("WARN Failed to load buffer view")
			return nil, false
		}

		view := d.BufferViews[*accessor.BufferView]
		buffer := d.Buffers[view.Buffer]
		slice := buffer.Data[view.ByteOffset : view.ByteOffset+view.ByteLength]

		var vao uint32
		gl.GenVertexArrays(1, &vao)
		gl.BindVertexArray(vao)

		gl.BindBuffer(gl.ARRAY_BUFFER, vbo)
		gl.BufferData(gl.ARRAY_BUFFER, int(view.ByteLength), gl.Ptr(slice), gl.STATIC_DRAW)

		// Setup the position attribute
		positionIndex := uint32(gl.GetAttribLocation(program, gl.Str("position\x00")))
		gl.EnableVertexAttribArray(positionIndex)
		gl.VertexAttribPointer(positionIndex, 3, gl.FLOAT, false, 0, nil)

		mesh.VertexArrayID = vao
		mesh.VertexCount = int32(accessor.Count)

		model.Meshes[i] = mesh
	}

	return &model, true
}
