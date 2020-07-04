package engine

// engine runs the game
type engine struct {
	providers *HostProviders
}

// Run the game starting with the given entry scene
func (e *engine) Run(entry string) {
	settings := e.providers.Settings.New()
	window := e.providers.Platform.NewWindow(settings)
	timer := e.providers.Platform.NewTimer(settings)
	platform := e.providers.Platform.NewPlatform(settings)
	graphics := e.providers.Graphics.New(settings)
	manager := e.providers.Manager.New(settings, graphics)

	scene, world := manager.LoadStartScene(entry)

	context := &Context{
		previousFrameTime: timer.GetTime(),
		frameTime:         timer.GetTime(),
	}

	defer platform.Terminate()
	for !window.ShouldClose() {
		context.frameTime = timer.GetTime()

		graphics.Clear()
		scene.SetupScene(graphics)

		e.update(world, context)
		e.render(scene, graphics)
		e.synchronize(world)

		window.SwapBuffers()
		platform.PollEvents()

		context.previousFrameTime = context.frameTime
	}
}

func (e *engine) update(world World, context *Context) {
	for _, o := range world.Objects() {
		o.Update(context)
	}
}

func (e *engine) render(s Scene, g Graphics) {
	for _, renderable := range s.Renderables() {
		renderable.Render(g)
	}
}

func (e *engine) synchronize(w World) {
	for _, u := range w.Objects() {
		u.Synchronize()
	}
}
