package shaders

import "github.com/go-gl/gl/v4.1-core/gl"

// ShaderType is strongly typed shader type
type ShaderType uint32

// Shader types
const (
	Vertex   ShaderType = gl.VERTEX_SHADER
	Fragment ShaderType = gl.FRAGMENT_SHADER
)

// Shader provides structure around pipeline shaders
type Shader struct {
	source     string
	shaderType ShaderType
	shaderID   uint32
}

// GetID gets the platform shader ID
func (s *Shader) GetID() uint32 {
	return s.shaderID
}
