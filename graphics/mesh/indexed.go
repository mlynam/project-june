package mesh

import (
	"github.com/mlynam/project-june/graphics/gl"
	"github.com/mlynam/project-june/graphics/shader"
)

// Indexed draws by indexed vertices
type Indexed struct {
	*Shared
	vertexBufferObjectID uint32
	indexBufferObjectID  uint32
	elementCount         int32
}

// Render an indexed mesh
func (i *Indexed) Render(program *shader.Program) {
	ok := i.SetupDraw(program)

	if ok {
		gl.BindBuffer(gl.ELEMENT_ARRAY_BUFFER, i.indexBufferObjectID)
		gl.DrawElements(gl.TRIANGLES, i.elementCount, gl.UNSIGNED_SHORT, gl.Ptr(0))
	}
}
