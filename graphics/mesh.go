package graphics

import (
	"fmt"

	"github.com/go-gl/mathgl/mgl32"
	"github.com/mlynam/project-june/shared"
)

// Mesh holds information about mesh data loaded into the graphics card
type Mesh struct {
	VertexBufferID uint32
	IndexBufferID  *uint32
	VertexCount    uint32
}

// Render a mesh
func (m *Mesh) Render(o shared.Locatable) {
	world := o.Locate()

	if world == mgl32.Ident4() {
		fmt.Print("Found identiy matrix")
	}
}
