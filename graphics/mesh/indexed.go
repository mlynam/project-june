package mesh

import (
	"github.com/go-gl/gl/v4.1-core/gl"
	"github.com/mlynam/project-june/engine"
)

// Indexed draws by indexed vertices
type Indexed struct {
	*Shared
	vertexBufferObjectID uint32
	indexBufferObjectID  uint32
	elementCount         int32
}

// Render an indexed mesh
func (i *Indexed) Render(g engine.Graphics) {
	ok := i.SetupDraw(g)

	if ok {
		gl.BindBuffer(gl.ELEMENT_ARRAY_BUFFER, i.indexBufferObjectID)
		gl.DrawElements(gl.TRIANGLES, i.elementCount, gl.UNSIGNED_SHORT, gl.Ptr(0))
	}
}
