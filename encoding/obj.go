package encoding

import "github.com/mlynam/project-june/shared"

// OBJ provides an in memory structure for the obj data type
type OBJ struct {
	verts, normals []shared.Vector3
	textureCoords  []shared.Vector2
	name, filename string
	faces          []OBJFace
}

// OBJFace describes a face in the model
type OBJFace struct {
	vertex   []shared.Vector3
	normal   []shared.Vector3
	uv       []shared.Vector2
	material string
	smooth   bool
}
