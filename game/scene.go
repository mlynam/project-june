package game

import (
	"github.com/mlynam/project-june/engine"
	"github.com/mlynam/project-june/graphics"
)

// Scene contains data to be rendered
type Scene struct {
	camera      *graphics.Camera
	renderables []engine.Renderable
}

// NewScene with no renderables
func NewScene(c *graphics.Camera) *Scene {
	return &Scene{
		camera:      c,
		renderables: make([]engine.Renderable, 0),
	}
}

// Renderables to render
func (s *Scene) Renderables() []engine.Renderable {
	return s.renderables
}

// AddRenderable to the scene
func (s *Scene) AddRenderable(r engine.Renderable) {
	s.renderables = append(s.renderables, r)
}
