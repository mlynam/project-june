package engine

// PlatformProvider creates platform types
type PlatformProvider interface {
	NewWindow(Settings) Window
	NewTimer(Settings) Timer
	NewPlatform(Settings) Platform
}

// Platform provides platform level access
type Platform interface {
	PollEvents()
	Terminate()
}

// Timer gets high performance time
type Timer interface {
	GetTime() float64
}

// Window the engine will use
type Window interface {
	ShouldClose() bool
	SwapBuffers()
}
