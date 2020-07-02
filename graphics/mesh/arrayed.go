package mesh

import (
	"github.com/mlynam/project-june/graphics/gl"
	"github.com/mlynam/project-june/graphics/shader"
)

// Arrayed meshes draw without an index
type Arrayed struct {
	*Shared
	vertexArrayBufferID uint32
	elementCount        int32
}

// Render the arrayed mesh
func (a *Arrayed) Render(program *shader.Program) {
	ok := a.SetupDraw(program)

	if ok {
		gl.BindVertexArray(a.vertexArrayBufferID)
		gl.DrawArrays(gl.TRIANGLES, 0, a.elementCount)
	}
}
