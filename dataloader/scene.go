package dataloader

import (
	"github.com/mlynam/project-june/core"
	"github.com/mlynam/project-june/game"
	"github.com/mlynam/project-june/graphics"
	"github.com/mlynam/project-june/shared"
	"github.com/qmuntal/gltf"
)

// LoadScene loads a document and signals when complete
func LoadScene(doc *gltf.Document, program uint32) *core.Scene {
	objects := make([]*shared.Object, len(doc.Nodes))
	renderables := make([]graphics.Renderable, 0)
	var camera *graphics.Camera

	for i, node := range doc.Nodes {
		object := &shared.Object{
			Name:     node.Name,
			Position: shared.Vec3Convert64To32(node.Translation),
			Scale:    shared.Vec3Convert64To32(node.Scale),
			Rotation: shared.Vec4Convert64ToQuat(node.Rotation),
		}

		if len(node.Children) > 0 {
			min := node.Children[0]
			max := node.Children[len(node.Children)-1] + 1
			object.Children = objects[min:max]

			for _, child := range object.Children {
				child.Parent = object
			}
		}

		if node.Mesh != nil {
			model, ok := LoadModel(doc.Meshes[*node.Mesh], doc, program)

			if ok {
				for _, mesh := range model.Meshes {
					mesh.World = object
					renderables = append(renderables, &mesh)
				}
			}
		}

		if node.Name == "DefaultCamera" {
			for _, child := range node.Children {
				if idx := doc.Nodes[child].Camera; idx != nil {
					perspective := doc.Cameras[*idx].Perspective

					if perspective != nil {
						camera = &graphics.Camera{
							Object:      object,
							FieldOfView: float32(perspective.Yfov),
							ZFar:        float32(*perspective.Zfar),
							ZNear:       float32(perspective.Znear),
						}

						// We use the first camera we find in the DefaultCamera object
						break
					}
				}
			}
		}

		objects[i] = object
	}

	updatable := make([]shared.Updatable, len(objects))
	for i, object := range objects {
		if object.Name == "Cube" {
			updatable[i] = &game.Rotator{
				Object: object,
			}
		} else {
			updatable[i] = object
		}
	}

	return &core.Scene{
		Camera:      camera,
		Updatables:  updatable,
		Renderables: renderables,
	}
}
