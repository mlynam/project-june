package vertex

import (
	"fmt"
	"reflect"
	"strings"

	"github.com/go-gl/gl/v4.1-core/gl"
)

var reference = Vertex{}
var vertexType = reflect.TypeOf(Vertex{})

// Vertex data
type Vertex struct {
	Position [3]float32
	Normal   [3]float32
	UV       [2]float32
	Color    [3]float32
}

// AttributeFor the vertex
func (v *Vertex) AttributeFor(name string, normalized bool) Attribute {
	normalName := strings.ToLower(name)
	field, found := vertexType.FieldByNameFunc(func(n string) bool {
		return normalName == strings.ToLower(n)
	})

	if !found {
		panic(fmt.Errorf("%v attribute not found", name))
	}

	attribute := Attribute{
		Name:       normalName,
		Normalized: normalized,
		Stride:     vertexType.Size(),
		Offset:     int(field.Offset),
		Xtype:      gl.FLOAT,
	}

	switch normalName {
	case "position", "normal", "color":
		attribute.Size = 3
	default:
		attribute.Size = 2
	}

	return attribute
}
