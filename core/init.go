package core

import "github.com/mlynam/project-june/graphics"

// Init contains the initialization values for the core
type Init struct {
	Name     string
	Width    int
	Height   int
	Graphics graphics.Init
}
