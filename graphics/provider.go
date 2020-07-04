package graphics

import (
	"io/ioutil"
	"log"

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

	g := NewGraphics(program)

	gl.UseProgram(program)
	gl.Enable(gl.DEPTH_TEST)
	gl.DepthFunc(gl.LESS)
	gl.ClearColor(100/256.0, 149/256.0, 237/256.0, 1.0)

	g.EnsureSuccessState()

	return g
}
