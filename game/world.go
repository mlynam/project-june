package game

import (
	"github.com/mlynam/project-june/graphics"
	"github.com/mlynam/project-june/graphics/shader"
	"github.com/mlynam/project-june/shared"
)

// World contains the base elements that make up the game world
type World struct {
	camera     *graphics.Camera
	renderable []graphics.Renderable
	updatable  []shared.Updatable
}

// NewWorld creates a new game world given the base elements
func NewWorld(camera *graphics.Camera, renderables []graphics.Renderable, updatables []shared.Updatable) *World {
	return &World{
		camera:     camera,
		renderable: renderables,
		updatable:  updatables,
	}
}

// Update the world with the given context
func (w *World) Update(c *shared.Context) {
	for _, updatable := range w.updatable {
		updatable.Update(c)
	}
}

// Render the world with the given shader program
func (w *World) Render(p *shader.Program) {
	for _, renderable := range w.renderable {
		renderable.Render(p)
	}
}
