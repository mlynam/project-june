package graphics

import (
	"fmt"
	"io/ioutil"
	"log"
	"strings"

	"github.com/go-gl/gl/v4.1-core/gl"
	"github.com/mlynam/project-june/engine"
	"github.com/mlynam/project-june/graphics/shader"
)

// Provider creates a graphics engine
type Provider struct{}

// New creates a new graphics type
func (p *Provider) New(s engine.Settings) engine.Graphics {
	if err := gl.Init(); err != nil {
		panic(err)
	}

	version := gl.GoStr(gl.GetString(gl.VERSION))
	log.Println("INFO OpenGL version", version)

	program := gl.CreateProgram()

	shaders := map[shader.Type]string{
		shader.Vertex:   "assets/shaders/basic.vert",
		shader.Fragment: "assets/shaders/basic.frag",
	}

	for t, path := range shaders {
		bytes, err := ioutil.ReadFile(path)
		if err != nil {
			panic(err)
		}

		source := string(bytes)
		shader, ok := shader.New(source, t)
		if !ok {
			panic("Failed to create shader")
		}

		gl.AttachShader(program, shader.GetID())
	}

	gl.LinkProgram(program)

	ensureProgramLinkSuccess(program)

	return &Graphics{program}
}

func ensureProgramLinkSuccess(program uint32) {
	var status int32
	gl.GetProgramiv(program, gl.LINK_STATUS, &status)
	if status == gl.FALSE {
		var logLength int32
		gl.GetProgramiv(program, gl.INFO_LOG_LENGTH, &logLength)

		log := strings.Repeat("\x00", int(logLength+1))
		gl.GetProgramInfoLog(program, logLength, nil, gl.Str(log))

		panic(fmt.Sprintf("failed to link program: %v", log))
	}
}
