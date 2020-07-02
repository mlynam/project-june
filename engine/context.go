package engine

// Context of the current frame
type Context struct {
	Window   Window
	Graphics Graphics
}

// Delta time for the frame
func (c *Context) Delta() float64 {
	return 0
}
