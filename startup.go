package main

import (
	"github.com/mlynam/project-june/core"
	"github.com/mlynam/project-june/dataloader"
	"github.com/qmuntal/gltf"
)

const (
	source        = "assets/scene/BoxTaurus.gltf"
	meshLoadError = "Failed to load mesh data with name %v\n"
)

func config(c *core.Core) {
	doc, _ := gltf.Open(source)

	c.Scene = dataloader.LoadScene(doc)

	// for {
	// 	select {
	// 	case status := <-task.Status:
	// 		fmt.Println(status)
	// 	case <-task.Done:
	// 		c.Scene =
	// 		return
	// 	default:
	// 		glfw.PollEvents()
	// 	}
	// }
}
