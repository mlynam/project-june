package meshloader

import (
	"bytes"
	"encoding/binary"
	"fmt"
	"log"
	"strings"

	"github.com/go-gl/gl/v4.1-core/gl"
	"github.com/mlynam/project-june/engine"
	"github.com/mlynam/project-june/graphics"
	"github.com/mlynam/project-june/graphics/vertex"
	"github.com/qmuntal/gltf"
)

var (
	sizes = map[gltf.AccessorType]uint32{
		gltf.AccessorScalar: 1,
		gltf.AccessorVec2:   2,
		gltf.AccessorVec3:   3,
		gltf.AccessorVec4:   4,
		gltf.AccessorMat2:   4,
		gltf.AccessorMat3:   9,
		gltf.AccessorMat4:   16,
	}

	xtypes = map[gltf.ComponentType]uint32{
		gltf.ComponentFloat:  gl.FLOAT,
		gltf.ComponentByte:   gl.BYTE,
		gltf.ComponentUbyte:  gl.UNSIGNED_BYTE,
		gltf.ComponentShort:  gl.SHORT,
		gltf.ComponentUshort: gl.UNSIGNED_SHORT,
		gltf.ComponentUint:   gl.UNSIGNED_INT,
	}
)

var (
	supported = []string{
		"POSITION",
		"NORMAL",
		"COLOR",
		"TEXCOORD_0",
	}
)

// LoadFromGLTFDoc loads a mesh from a gltf primitive into graphics memory
func LoadFromGLTFDoc(doc *gltf.Document, p *gltf.Primitive, locatable engine.Locatable) *graphics.Mesh {
	if p.Indices == nil {
		panic(fmt.Errorf("unindexed mesh unsupported"))
	}

	var (
		positions *bytes.Reader
		normals   *bytes.Reader
		colors    *bytes.Reader
		uvs       *bytes.Reader
		reference vertex.Vertex
		err       error

		vertices = make([]vertex.Vertex, 0)
	)

	for {
		vertex := vertex.Vertex{}

		if positions != nil {
			err = binary.Read(positions, binary.LittleEndian, &vertex.Position)
		}

		if normals != nil {
			err = binary.Read(normals, binary.LittleEndian, &vertex.Normal)
		}
	}

	data := make([]byte, 0)
	attributes := make([]vertex.Attribute, 0)

	for _, attr := range supported {
		i, ok := p.Attributes[attr]
		if !ok {
			continue
		}

		accessor := doc.Accessors[i]
		if accessor.BufferView == nil {
			log.Printf("WARN invalid accessor data for %v", accessor.Name)
			continue
		}

		view := doc.BufferViews[*accessor.BufferView]
		buffer := doc.Buffers[view.Buffer]
		slice := buffer.Data[view.ByteOffset : view.ByteOffset+view.ByteLength]

		name := strings.ToLower(attr)

		attributes = append(attributes, vertex.Attribute{
			Name:       name,
			Normalized: accessor.Normalized,
			Offset:     len(data),
			Stride:     view.ByteStride,
			Size:       sizes[accessor.Type],
			Xtype:      xtypes[accessor.ComponentType],
		})

		data = append(data, slice...)
	}

	accessor := doc.Accessors[*p.Indices]
	view := doc.BufferViews[*accessor.BufferView]
	buffer := doc.Buffers[view.Buffer]
	index := buffer.Data[view.ByteOffset : view.ByteOffset+view.ByteLength]

	mesh := graphics.New(data, index, locatable)

	for _, attr := range attributes {
		mesh.AddAttribute(attr)
	}

	return mesh
}

func tryReadComponent(b *bytes.Reader, data interface{}) {

}
