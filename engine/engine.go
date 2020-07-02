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
	graphics := e.providers.Graphics.New(settings)
	manager := e.providers.Manager.New(settings)

	scene, world := manager.LoadStartScene(entry)

	context := &Context{
		previousFrameTime: timer.GetTime(),
		frameTime:         timer.GetTime(),
	}

	for !window.ShouldClose() {
		context.frameTime = timer.GetTime()
		done := make(chan bool, 1)

		go e.update(done, world, context)
		e.render(scene, graphics)

		<-done

		e.synchronize(world)
		context.previousFrameTime = context.frameTime
	}
}

func (e *engine) update(done chan bool, world World, context *Context) {
	for _, updatable := range world.Objects() {
		updatable.Update(context)
	}

	done <- true
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
