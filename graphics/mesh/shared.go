package mesh

import (
	"log"

	"github.com/go-gl/gl/v4.1-core/gl"
	"github.com/mlynam/project-june/engine"
)

// Shared includes shared functionality for all mesh types
type Shared struct {
	world      engine.Locatable
	drawFailed bool
}

// SetupDraw sets up the graphics pipeline to draw the mesh. Returns true
// if the mesh is safe to draw
func (s *Shared) SetupDraw(g engine.Graphics) bool {
	index, ok := g.Uniform("model")
	if ok {
		transform := s.world.Locate()
		gl.UniformMatrix4fv(index, 1, false, &transform[0])
	}

	index, ok = g.Attribute("position")
	if !ok {
		log.Print("FAIL unable to draw mesh, 'position' attribute not found.")
		s.drawFailed = true
	} else {
		gl.VertexAttribPointer(uint32(index), 3, gl.FLOAT, false, 0, nil)
		gl.EnableVertexAttribArray(uint32(index))
	}

	return s.drawFailed
}
