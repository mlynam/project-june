package core

// State provides a type for the core state
type State string

// Core game state
const (
	Startup      State = "Startup"
	GameMenu     State = "Game Menu"
	LoadingLevel State = "Loading Level"
	LoadingSave  State = "Loading Save"
	Running      State = "Running"
	Shutdown     State = "Shutdown"
)

// State returns the current state of the core
func (c *core) State() State {
	return c.state
}
