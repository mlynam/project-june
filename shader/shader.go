package shader

import "github.com/go-gl/gl/v4.6-core/gl"

// Type is strongly typed shader type
type Type uint32

// Shader types
const (
	Vertex   Type = gl.VERTEX_SHADER
	Fragment Type = gl.FRAGMENT_SHADER
)

// Shader provides structure around pipeline shaders
type Shader struct {
	source     string
	shaderType Type
	shaderID   uint32
}

// GetID gets the platform shader ID
func (s *Shader) GetID() uint32 {
	return s.shaderID
}
