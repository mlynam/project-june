package engine

type Scene interface {
	Renderables []Renderable
	Updatables []Updatable
}
