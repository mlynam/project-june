package graphics

import (
	"github.com/mlynam/project-june/graphics/shader"
)

// Renderable types implement a render method that will render something
type Renderable interface {
	Render(*shader.Program)
}
