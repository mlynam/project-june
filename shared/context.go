package shared

// Context of the currently rendering frame
type Context struct {
	delta    float64
	platform Platform
}

// NewContext returns a context with the parameters provided
func NewContext(delta float64, platform Platform) *Context {
	return &Context{
		delta,
		platform,
	}
}

// TimeDelta in seconds for the given context
func (c *Context) TimeDelta() float64 {
	return c.delta
}

// Platform for the given context
func (c *Context) Platform() Platform {
	return c.platform
}
