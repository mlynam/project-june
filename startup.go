package main

import (
	"log"

	"github.com/mlynam/project-june/core"
	"github.com/mlynam/project-june/graphics"
	"github.com/mlynam/project-june/graphics/renderable"
	"github.com/qmuntal/gltf"
)

const (
	source        = "assets/scene/BoxTaurus.gltf"
	meshLoadError = "Failed to load mesh data with name %v\n"
)

func config(c *core.Core) {
	doc, _ := gltf.Open(source)
	meshes := make([]*graphics.Render, 0)

	for _, mesh := range doc.Meshes {
		for _, primitive := range mesh.Primitives {
			renderable, ok := renderable.PrimitiveToUnindexed(primitive, doc)

			if ok {
				meshes = append(meshes, &renderable)
			} else {
				log.Printf(meshLoadError, mesh.Name)
			}
		}
	}

	c.Renderables = meshes
}
