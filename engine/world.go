package engine

// World contains all updatable elements
type World interface {
	Objects() []Object
}
