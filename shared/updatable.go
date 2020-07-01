package shared

// Updatable types are updatable by the engine
type Updatable interface {
	Update(*Context)
}
