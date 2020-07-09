package engine

// Scene can be rendered
type Scene interface {
	SetupScene(Graphics)
	Renderables() []Renderable
	SceneViewProjection() [16]float32
}
