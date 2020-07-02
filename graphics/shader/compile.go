package shader

import (
	"fmt"
	"strings"

	"github.com/go-gl/gl/v4.1-core/gl"
)

func (shader *Shader) compile() bool {
	glShader := gl.CreateShader(uint32(shader.shaderType))

	csource, free := gl.Strs(shader.source)
	gl.ShaderSource(glShader, 1, csource, nil)
	free()
	gl.CompileShader(glShader)

	var status int32
	gl.GetShaderiv(glShader, gl.COMPILE_STATUS, &status)
	if status == gl.FALSE {
		var logLength int32
		gl.GetShaderiv(glShader, gl.INFO_LOG_LENGTH, &logLength)

		log := strings.Repeat("\x00", int(logLength+1))
		gl.GetShaderInfoLog(glShader, logLength, nil, gl.Str(log))

		fmt.Printf("failed to compile\n%v: %v", shader.source, log)
		return false
	}

	shader.shaderID = glShader
	return true
}
