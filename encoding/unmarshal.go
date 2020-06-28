package encoding

import (
	"bytes"
	"fmt"
	"io"

	"github.com/mlynam/project-june/shared"
)

// UnmarshalBinary attempts to read data as an obj
func (obj *OBJ) UnmarshalBinary(data []byte) error {
	b := bytes.NewBuffer(data)

	var (
		flag   string
		mtllib string
		usemtl string
		smooth string
		err    error
	)

	for {
		flag, err = peekFlag(b)
		if err == io.EOF {
			break
		}

		switch flag {
		case "o":
			fmt.Fscanln(b, &obj.name)
		case "v":
			v := shared.Vector3{}
			fmt.Fscanln(b, &v[0], &v[1], &v[2])
			obj.verts = append(obj.verts, v)
		case "vt":
			v := shared.Vector2{}
			fmt.Fscanln(b, &v[0], &v[1])
			obj.textureCoords = append(obj.textureCoords, v)
		case "vn":
			v := shared.Vector3{}
			fmt.Fscanln(b, &v[0], &v[1], &v[2])
			obj.normals = append(obj.normals, v)
		case "mtllib":
			fmt.Fscanln(b, &mtllib)
		case "usemtl":
			fmt.Fscanln(b, &usemtl)
		case "s":
			fmt.Fscanln(b, &smooth)
		default: // Skip anything not supported, like comments
			b.ReadBytes(byte('\n'))
		}
	}

	return io.EOF
}

func peekFlag(b *bytes.Buffer) (string, error) {
	var flag string
	defer b.UnreadByte()
	_, err := fmt.Fscanln(b, &flag)
	return flag, err
}
