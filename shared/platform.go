package shared

// Platform provides information about the platform the engine is running on
type Platform interface {
	AspectRatio() float32
	OpenGLProgram() uint32
	Name() string
}
