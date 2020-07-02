package mesh

import (
	"github.com/go-gl/gl/v4.1-core/gl"
	"github.com/mlynam/project-june/engine"
)

// Arrayed meshes draw without an index
type Arrayed struct {
	*Shared
	vertexArrayBufferID uint32
	elementCount        int32
}

// Render the arrayed mesh
func (a *Arrayed) Render(g engine.Graphics) {
	ok := a.SetupDraw(g)

	if ok {
		gl.BindVertexArray(a.vertexArrayBufferID)
		gl.DrawArrays(gl.TRIANGLES, 0, a.elementCount)
	}
}
