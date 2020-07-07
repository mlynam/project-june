package meshloader

import (
	"bytes"
	"encoding/binary"
	"fmt"
	"io"

	"github.com/go-gl/gl/v4.1-core/gl"
	"github.com/mlynam/project-june/engine"
	"github.com/mlynam/project-june/graphics"
	"github.com/mlynam/project-june/graphics/vertex"
	"github.com/qmuntal/gltf"
)

var (
	sizes = map[gltf.AccessorType]int{
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

// LoadFromGLTFDoc loads a mesh from a gltf primitive into graphics memory
func LoadFromGLTFDoc(doc *gltf.Document, p *gltf.Primitive, locatable engine.Locatable) *graphics.Mesh {
	if p.Indices == nil {
		panic(fmt.Errorf("unindexed mesh unsupported"))
	}

	var (
		positions *bytes.Reader
		normals   *bytes.Reader
		uvs       *bytes.Reader
		reference vertex.Vertex

		vertices   = make([]vertex.Vertex, 0)
		indices    = make([]uint, 0)
		attributes = make([]vertex.Attribute, 0)
	)

	// Setup position
	i, ok := p.Attributes["POSITION"]
	if !ok {
		panic("position data not found")
	}

	positions, accessor := attributeReader(i, doc)
	attributes = append(attributes, reference.AttributeFor("position", accessor.Normalized))

	// Normals and UVs are optional and therefore we don't panic if we don't find them
	i, ok = p.Attributes["NORMAL"]
	if ok {
		normals, accessor = attributeReader(i, doc)
		attributes = append(attributes, reference.AttributeFor("normal", accessor.Normalized))
	}
	i, ok = p.Attributes["TEXCOORD_0"]
	if ok {
		uvs, accessor = attributeReader(i, doc)
		attributes = append(attributes, reference.AttributeFor("uv", accessor.Normalized))
	}

	for len(vertices) < int(accessor.Count) {
		vertex := vertex.Vertex{}

		tryReadComponent(positions, &vertex.Position)
		tryReadComponent(normals, &vertex.Normal)
		tryReadComponent(uvs, &vertex.UV)

		vertices = append(vertices, vertex)
	}

	index, accessor := attributeReader(*p.Indices, doc)
	for len(indices) < int(accessor.Count) {
		var value uint
		switch accessor.ComponentType {
		case gltf.ComponentUbyte:
			var data uint8
			tryReadComponent(index, &data)
			value = uint(data)
		case gltf.ComponentUshort:
			var data uint16
			tryReadComponent(index, &data)
			value = uint(data)
		case gltf.ComponentUint:
			var data uint32
			tryReadComponent(index, &data)
			value = uint(data)
		}

		indices = append(indices, value)
	}

	mesh := graphics.New(vertices, indices, locatable)

	for _, attr := range attributes {
		mesh.AddAttribute(attr)
	}

	return mesh
}

func attributeReader(accessorIndex uint32, doc *gltf.Document) (*bytes.Reader, *gltf.Accessor) {
	accessor := doc.Accessors[accessorIndex]
	view := doc.BufferViews[*accessor.BufferView]
	buffer := doc.Buffers[view.Buffer]
	slice := buffer.Data[view.ByteOffset : view.ByteOffset+view.ByteLength]
	return bytes.NewReader(slice), accessor
}

func tryReadComponent(b *bytes.Reader, data interface{}) {
	if b == nil {
		return
	}

	err := binary.Read(b, binary.LittleEndian, data)

	switch err {
	case io.EOF, io.ErrUnexpectedEOF:
		panic("reader EOF")
	}
}
