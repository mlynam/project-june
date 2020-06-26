package core

// Init contains the initialization values for the core
type Init struct {
	Name   string
	Width  int
	Height int
}

// New is a function that initializes the core
func New(init *Init) *Core {
	c := Core{}

	c.state = Startup
	c.name = init.Name
	c.width = init.Width
	c.height = init.Height

	return c.initWindow()
}
