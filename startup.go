package main

import (
	"github.com/mlynam/project-june/core"
	"github.com/mlynam/project-june/dataloader"
	"github.com/qmuntal/gltf"
)

const source = "assets/scene/BoxTaurus.gltf"

func config(c *core.Core) {
	doc, _ := gltf.Open(source)

	c.Scene = dataloader.LoadScene(doc, c.OpenGLProgram())
	if c.Scene.Camera == nil {
		panic("No valid DefaultCamera found")
	}

	c.Scene.Camera.AspectRatio = c.AspectRatio()
	c.Scene.Updatables = append(c.Scene.Updatables, c)
}
