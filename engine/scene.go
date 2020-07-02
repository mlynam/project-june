package engine

// Scene can be rendered
type Scene interface {
	Renderables() []Renderable
}
