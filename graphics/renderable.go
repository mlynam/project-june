package graphics

import "github.com/mlynam/project-june/shared"

// Renderable composes an object and a model
type Renderable struct {
	*shared.Object
	*Model
}

// Render renders the meshes contained in the renderable model
func (r *Renderable) Render() {
	for _, mesh := range r.Model.Meshes {
		mesh.Render(r)
	}
}
