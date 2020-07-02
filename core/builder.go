package core

import (
	"errors"
	"fmt"

	"github.com/mlynam/project-june/shared"
)

// OnConfigureCore used to configure the core options
type OnConfigureCore func(*Config)

// Builder contains core builder info
type Builder struct {
	core   *core
	config *Config
}

// NewBuilder creates a default core builder
func NewBuilder() *Builder {
	core := &core{}
	return &Builder{
		core: core,
		config: &Config{
			Name:   "Project June",
			Width:  1920,
			Height: 1024,
		},
	}
}

// UseLoop allows callers to use a core loop action
func (b *Builder) UseLoop(action OnLoop) {
	ensureNotNilArg(action, "action")
	b.ensureBuilderValidCore()

	b.core.onLoop = action
}

// ConfigureCore allows callers to configure the core options
func (b *Builder) ConfigureCore(action OnConfigureCore) {
	ensureNotNilArg(action, "action")

	config := &Config{}
	action(config)

	b.config = config
}

// UseStartup allows callers to use a startup action
func (b *Builder) UseStartup(action OnCoreAction) {
	ensureNotNilArg(action, "action")
	b.ensureBuilderValidCore()

	b.core.onStartup = action
}

// Build and return the core platform
func (b *Builder) Build() shared.Platform {
	b.core.initWindow()

	return b.core
}

// UseShutdown allows callers to use a shutdown action
func (b *Builder) UseShutdown(action OnCoreAction) {
	ensureNotNilArg(action, "action")
	b.ensureBuilderValidCore()

	b.core.onShutdown = action
}

func (b *Builder) ensureBuilderValidCore() {
	if b.core == nil {
		panic(errors.New("invalid builder state: use 'builder.New()' to createa a builder"))
	}
}

func ensureNotNilArg(arg interface{}, name string) {
	if arg == nil {
		panic(fmt.Errorf("argument nil: %s", name))
	}
}
