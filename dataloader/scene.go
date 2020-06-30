package dataloader

import (
	"github.com/mlynam/project-june/core"
	"github.com/mlynam/project-june/graphics"
	"github.com/mlynam/project-june/shared"
	"github.com/qmuntal/gltf"
)

// LoadScene loads a document and signals when complete
func LoadScene(doc *gltf.Document) *core.Scene {
	objects := make([]*shared.Object, len(doc.Nodes))
	renderables := make([]*graphics.Renderable, 0)
	cameras := make([]*graphics.Camera, 0)

	for i, node := range doc.Nodes {
		objects[i] = &shared.Object{
			Name:     node.Name,
			Position: shared.Vec3Convert64To32(node.Translation),
			Scale:    shared.Vec3Convert64To32(node.Scale),
			Rotation: shared.Vec4Convert64to32(node.Rotation),
		}

		if len(node.Children) > 0 {
			min := node.Children[0]
			max := node.Children[len(node.Children)-1] + 1
			objects[i].Children = objects[min:max]

			for _, child := range objects[i].Children {
				child.Parent = objects[i]
			}
		}

		if node.Mesh != nil {
			model, ok := LoadModel(doc.Meshes[*node.Mesh], doc)

			if ok {
				renderables = append(renderables, &graphics.Renderable{Object: objects[i], Model: model})
			}
		}
	}

	return &core.Scene{
		Cameras:     cameras,
		Objects:     objects,
		Renderables: renderables,
	}
}
