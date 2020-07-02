package core

// Platform provides information about the platform the engine is running on
type Engine interface {
	Run()
	AspectRatio() float32
	Name() string
}
