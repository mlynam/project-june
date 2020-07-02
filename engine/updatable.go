package engine

// Updatable will be updated during the engine loop
type Updatable interface {
	Update(*Context)

	// Synchronize perform any internal updates needed for the next render
	Synchronize()
}
