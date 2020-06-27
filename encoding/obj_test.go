package encoding

import (
	"bytes"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestObj(t *testing.T) {
	var source = `# Blender ...
mtllib mat.mtl
o Cube
v 1.0 1.0 1.0
`

	reader := bytes.NewBufferString(source)
	var obj OBJ
	obj.UnmarshalBinary(reader.Bytes())

	assert.Equal(t, "Cube", obj.name)
	assert.Len(t, obj.verts, 1)
}
