package engine

// Locatable can be located with a world space matrix
type Locatable interface {
	Position() [3]float32
	// Locate something in world space as a tranformation matrix
	Locate() [16]float32
}
