package game

import "github.com/mlynam/project-june/engine"

// World containing the game objects
type World struct {
	objects []engine.Object
}

// NewWorld with no objects
func NewWorld() *World {
	return &World{
		objects: make([]engine.Object, 0),
	}
}

// Objects of the world
func (w *World) Objects() []engine.Object {
	return w.objects
}

// AddObject to the world
func (w *World) AddObject(o engine.Object) {
	w.objects = append(w.objects, o)
}
