package mesh

import (
	"log"

	"github.com/mlynam/project-june/graphics/gl"
	"github.com/mlynam/project-june/graphics/shader"
	"github.com/mlynam/project-june/shared"
)

// Shared includes shared functionality for all mesh types
type Shared struct {
	world      shared.Locatable
	drawFailed bool
}

// SetupDraw sets up the graphics pipeline to draw the mesh. Returns true
// if the mesh is safe to draw
func (s *Shared) SetupDraw(p *shader.Program) bool {
	index, ok := p.Uniform("model")
	if ok {
		transform := s.world.Locate()
		gl.UniformMatrix4fv(index, 1, false, &transform[0])
	}

	index, ok = p.Attribute("position")
	if !ok {
		log.Print("FAIL unable to draw mesh, 'position' attribute not found.")
		s.drawFailed = true
	} else {
		gl.VertexAttribPointer(uint32(index), 3, gl.FLOAT, false, 0, nil)
		gl.EnableVertexAttribArray(uint32(index))
	}

	return s.drawFailed
}
