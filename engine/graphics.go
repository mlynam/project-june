package engine

// GraphicsProvider creates a graphics engine
type GraphicsProvider interface {
	New(Settings) Graphics
}

// Graphics engine
type Graphics interface {
	// Program is the shader program currently attached to the graphics pipeline
	Program() uint32
	Clear()

	SetProgram(uint32)
	SetScene(Scene)
	Attribute(name string) (int32, bool)
	Uniform(name string) (int32, bool)
	SceneViewProjection() [16]float32
	EnsureSuccessState()
}
