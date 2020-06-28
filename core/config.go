package core

// StartupConfigAction is a function used to configure the core during startup
type StartupConfigAction func(c *Core)

// Config the core with the given action
func (c *Core) Config(a StartupConfigAction) *Core {
	a(c)
	return c
}
