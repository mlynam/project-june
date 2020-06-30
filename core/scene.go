package core

import (
	"github.com/mlynam/project-june/graphics"
	"github.com/mlynam/project-june/shared"
)

// Scene contains all scene data
type Scene struct {
	Cameras     []*graphics.Camera
	Renderables []*graphics.Renderable
	Objects     []*shared.Object
}
