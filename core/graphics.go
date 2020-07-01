package core

import (
	"fmt"
	"io/ioutil"
	"log"
	"strings"

	"github.com/go-gl/gl/v4.1-core/gl"
	"github.com/mlynam/project-june/graphics"
	"github.com/mlynam/project-june/shader"
)

func (c *Core) initGraphics(init *graphics.Init) *Core {
	if err := gl.Init(); err != nil {
		panic(err)
	}

	version := gl.GoStr(gl.GetString(gl.VERSION))
	log.Println("INFO OpenGL version", version)

	program := gl.CreateProgram()

	for t, path := range init.Shaders {
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

		c.shaders = append(c.shaders, shader)
	}

	gl.LinkProgram(program)

	ensureProgramLinkSuccess(program)

	c.glProgram = program

	return c
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
