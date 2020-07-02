package engine

// PlatformProvider creates platform types
type PlatformProvider interface {
	NewWindow(Settings) Window
	NewTimer(Settings) Timer
}

// Timer gets high performance time
type Timer interface {
	GetTime() float64
}

// Window the engine will use
type Window interface {
	ShouldClose() bool
}
