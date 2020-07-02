package shader

import "github.com/go-gl/gl/v4.1-core/gl"

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

// New creates and compiles a new shader
func New(source string, shaderType Type) (Shader, bool) {
	shader := Shader{source + "\x00", shaderType, 0}
	ok := shader.compile()

	return shader, ok
}

// GetID gets the platform shader ID
func (s *Shader) GetID() uint32 {
	return s.shaderID
}
