package engine

// Context of the current frame
type Context struct {
	previousFrameTime float64
	frameTime         float64
}

// Delta returns the frame delta in seconds
func (c *Context) Delta() float64 {
	return c.frameTime - c.previousFrameTime
}
