package graphics

// Renderable types implement a render method that will render something
type Renderable interface {
	Render(uint32)
}
