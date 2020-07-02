package manager

import (
	"fmt"
	"log"

	"github.com/mlynam/project-june/engine"
	"github.com/mlynam/project-june/game"
	"github.com/mlynam/project-june/graphics"
	"github.com/mlynam/project-june/graphics/mesh"
	"github.com/mlynam/project-june/util"
	"github.com/qmuntal/gltf"
)

// Manager enables the current scene and world to load assets
type Manager struct{}

// LoadStartScene loads the scene with the given name and initializes a new world and scene
func (m *Manager) LoadStartScene(name string) (engine.Scene, engine.World) {
	path := fmt.Sprintf("assets/scenes/%v", name)
	doc, err := gltf.Open(path)
	if err != nil {
		panic(fmt.Errorf("Failed to load the start scene: %v", path))
	}

	var camera **graphics.Camera

	world := game.NewWorld()
	scene := game.NewScene(*camera)

	objects := make([]*game.Object, len(doc.Nodes))

	for i, node := range doc.Nodes {
		position := util.Vec3Convert64To32(node.Translation)
		scale := util.Vec3Convert64To32(node.Scale)
		rotation := util.Vec4Convert64ToQuat(node.Rotation)

		object := game.NewObject(node.Name, position, scale, rotation)

		if len(node.Children) > 0 {
			min := node.Children[0]
			max := node.Children[len(node.Children)-1] + 1

			for _, child := range objects[min:max] {
				object.AddChild(child)
			}
		}

		if node.Mesh != nil {
			for _, primitive := range doc.Meshes[*node.Mesh].Primitives {
				builder := mesh.Builder{}
				vertexID, ok := primitive.Attributes["POSITION"]

				if !ok {
					log.Print("WARN gltf failed to provide position data for primitive")
					continue
				}

				accessor := doc.Accessors[vertexID]
				view := doc.BufferViews[*accessor.BufferView]
				buffer := doc.Buffers[view.Buffer]
				slice := buffer.Data[view.ByteOffset : view.ByteOffset+view.ByteLength]

				builder.SetPositionData(view.ByteLength, slice)

				indexID := primitive.Indices
				if indexID != nil {
					accessor = doc.Accessors[*indexID]
					view = doc.BufferViews[*accessor.BufferView]
					buffer = doc.Buffers[view.Buffer]
					slice = buffer.Data[view.ByteOffset : view.ByteOffset+view.ByteLength]

					builder.SetIndexData(view.ByteLength, slice)
				}

				mesh := builder.Build(object)
				if mesh != nil {
					scene.AddRenderable(mesh)
				}
			}
		}

		if node.Name == "DefaultCamera" {
			for _, child := range node.Children {
				if idx := doc.Nodes[child].Camera; idx != nil {
					perspective := doc.Cameras[*idx].Perspective

					if perspective != nil {
						c := &graphics.Camera{
							Object:      object,
							FieldOfView: float32(perspective.Yfov),
							ZFar:        float32(*perspective.Zfar),
							ZNear:       float32(perspective.Znear),
						}

						camera = &c

						// We use the first camera we find in the DefaultCamera object
						break
					}
				}
			}
		}

		objects[i] = object
		world.AddObject(object)
	}

	return scene, world
}
