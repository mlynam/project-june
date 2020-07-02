package core

// Context of the currently rendering frame
type Context struct {
	delta  float64
	engine Engine
}

// NewContext returns a context with the parameters provided
func NewContext(delta float64, engine Engine) *Context {
	return &Context{
		delta,
		engine,
	}
}

// TimeDelta in seconds for the given context
func (c *Context) TimeDelta() float64 {
	return c.delta
}

// Platform for the given context
func (c *Context) Platform() Engine {
	return c.engine
}
