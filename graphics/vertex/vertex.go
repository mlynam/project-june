package vertex

import "unsafe"

// Vertex data
type Vertex struct {
	Position [3]float32
	Color    [3]float32
}

var reference = Vertex{}

// Offset values
var (
	PositionOffset = int(unsafe.Offsetof(reference.Position))
	ColorOffset    = int(unsafe.Offsetof(reference.Color))
)
