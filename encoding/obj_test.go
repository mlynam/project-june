package encoding

import (
	"bytes"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestObj(t *testing.T) {
	var source = `# Blender ...
		# www.blender.org
		mtllib cube.mtl
		o Cube
		v 1.000000 1.000000 -1.000000
		v 1.000000 -1.000000 -1.000000
		v 1.000000 1.000000 1.000000
		v 1.000000 -1.000000 1.000000
		v -1.000000 1.000000 -1.000000
		v -1.000000 -1.000000 -1.000000
		v -1.000000 1.000000 1.000000
		v -1.000000 -1.000000 1.000000
		vt 0.875000 0.500000
		vt 0.625000 0.750000
		vt 0.625000 0.500000
		vt 0.375000 1.000000
		vt 0.375000 0.750000
		vt 0.625000 0.000000
		vt 0.375000 0.250000
		vt 0.375000 0.000000
		vt 0.375000 0.500000
		vt 0.125000 0.750000
		vt 0.125000 0.500000
		vt 0.625000 0.250000
		vt 0.875000 0.750000
		vt 0.625000 1.000000
		vn 0.0000 1.0000 0.0000
		vn 0.0000 0.0000 1.0000
		vn -1.0000 0.0000 0.0000
		vn 0.0000 -1.0000 0.0000
		vn 1.0000 0.0000 0.0000
		vn 0.0000 0.0000 -1.0000
		usemtl Material
		s off
		f 5/1/1 3/2/1 1/3/1
		f 3/2/2 8/4/2 4/5/2
		f 7/6/3 6/7/3 8/8/3
		f 2/9/4 8/10/4 6/11/4
		f 1/3/5 4/5/5 2/9/5
		f 5/12/6 2/9/6 6/7/6
		f 5/1/1 7/13/1 3/2/1
		f 3/2/2 7/14/2 8/4/2
		f 7/6/3 5/12/3 6/7/3
		f 2/9/4 4/5/4 8/10/4
		f 1/3/5 3/2/5 4/5/5
		f 5/12/6 1/3/6 2/9/6
	`

	reader := bytes.NewBufferString(source)
	var obj OBJ
	obj.UnmarshalBinary(reader.Bytes())

	assert.Equal(t, "Cube", obj.name)
	assert.Len(t, obj.verts, 8)
	assert.Len(t, obj.normals, 6)
	assert.Len(t, obj.textureCoords, 14)
}
