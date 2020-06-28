package graphics

import "github.com/mlynam/project-june/shader"

// Init contains the initialization values for graphics
type Init struct {
	Shaders map[shader.Type]string
}
