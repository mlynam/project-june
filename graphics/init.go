package graphics

import "github.com/mlynam/project-june/shaders"

// Init contains the initialization values for graphics
type Init struct {
	Shaders map[shaders.ShaderType]string
}
