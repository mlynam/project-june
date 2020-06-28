package shader

// New creates and compiles a new shader
func New(source string, shaderType Type) (Shader, bool) {
	shader := Shader{source + "\x00", shaderType, 0}
	ok := shader.compile()

	return shader, ok
}
