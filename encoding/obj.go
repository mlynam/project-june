package encoding

import "github.com/mlynam/project-june/shared"

// OBJ provides an in memory structure for the obj data type
type OBJ struct {
	verts, normals []shared.Vector3
	textureCoords  []shared.Vector2
	name, filename string
}
