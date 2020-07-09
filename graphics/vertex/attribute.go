package vertex

// Attribute stores data about a particular vertex attribute, such as position
//	- Name is the name of the attribute. This should match the name used in
//		any graphics program.
//
//	- XType is the component type. If your position is made up of float(x,y,z),
//	then your component type is float.
//
//	- Stride is the stride between two of this attribute. If your vertex data
//	is interleaved, you need to set a stride.
//
//	- Offset is the offset to use into the vertex buffer to find this attribute. If
//	your data is tightly packed, you probably need an offset.
//
//	- Normalized is a flag indicating that your components are normalized so the
//	graphics pipeline knows how to load the data into a graphics program
type Attribute struct {
	Name       string
	Size       uint32
	Xtype      uint32
	Stride     uintptr
	Offset     int
	Normalized bool
}
