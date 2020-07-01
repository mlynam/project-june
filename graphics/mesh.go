package graphics

import (
	"github.com/go-gl/gl/v4.1-core/gl"
	"github.com/mlynam/project-june/shared"
)

// Mesh holds information about mesh data loaded into the graphics card
type Mesh struct {
	VertexArrayID uint32
	VertexCount   int32
	IndexBufferID *uint32
	World         shared.Locatable
}

// Render a mesh
func (m *Mesh) Render(program uint32) {
	index := gl.GetUniformLocation(program, gl.Str("model\x00"))
	if index > -1 {
		world := m.World.Locate()
		gl.UniformMatrix4fv(index, 1, false, &world[0])
	}

	gl.BindVertexArray(m.VertexArrayID)
	gl.DrawArrays(gl.TRIANGLES, 0, m.VertexCount)
}
