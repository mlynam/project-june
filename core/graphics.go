package core

import (
	"io/ioutil"
	"log"

	"github.com/go-gl/gl/v4.1-core/gl"
	"github.com/mlynam/project-june/graphics"
	"github.com/mlynam/project-june/shader"
)

func (c *Core) initGraphics(init *graphics.Init) *Core {
	if err := gl.Init(); err != nil {
		panic(err)
	}

	version := gl.GoStr(gl.GetString(gl.VERSION))
	log.Println("OpenGL version", version)

	prog := gl.CreateProgram()

	for t, path := range init.Shaders {
		bytes, err := ioutil.ReadFile(path)
		if err != nil {
			panic(err)
		}

		source := string(bytes)
		shader, ok := shader.New(source, t)
		if !ok {
			panic("Failed to create vertex shader")
		}

		gl.AttachShader(prog, shader.GetID())

		c.shaders = append(c.shaders, shader)
	}

	gl.LinkProgram(prog)

	c.glProgram = prog

	return c
}
