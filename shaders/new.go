package shaders

// New creates and compiles a new shader
func New(source string, shaderType ShaderType) (Shader, bool) {
	shader := Shader{source, shaderType, 0}
	ok := shader.compile()

	return shader, ok
}
