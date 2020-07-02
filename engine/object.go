package engine

// Object is a world object
type Object interface {
	Locatable

	Update(*Context)

	// Synchronize perform any internal updates needed for the next render
	Synchronize()
}
