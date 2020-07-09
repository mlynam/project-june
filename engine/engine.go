package engine

import "runtime"

// engine runs the game
type engine struct {
	providers *HostProviders
}

// Run the game starting with the given entry scene
func (e *engine) Run(entry string) {
	runtime.LockOSThread()

	settings := e.providers.Settings.New()
	window := e.providers.Platform.NewWindow(settings)
	timer := e.providers.Platform.NewTimer(settings)
	platform := e.providers.Platform.NewPlatform(settings)
	graphics := e.providers.Graphics.New(settings)
	manager := e.providers.Manager.New(settings, graphics)

	scene, world := manager.LoadStartScene(entry)

	graphics.SetScene(scene)
	context := &Context{
		previousFrameTime: timer.GetTime(),
		frameTime:         timer.GetTime(),
	}

	defer platform.Terminate()
	for !window.ShouldClose() {
		context.frameTime = timer.GetTime()
		done := make(chan bool, 1)

		scene.SetupScene(graphics)

		e.update(done, world, context)
		e.render(scene, graphics)

		<-done

		e.synchronize(world)

		window.SwapBuffers()
		platform.PollEvents()

		context.previousFrameTime = context.frameTime
	}
}

func (e *engine) update(done chan bool, world World, context *Context) {
	for _, object := range world.Objects() {
		object.Update(context)
	}

	done <- true
}

func (e *engine) render(s Scene, g Graphics) {
	for _, renderable := range s.Renderables() {
		renderable.Render(g)
	}
}

func (e *engine) synchronize(w World) {
	for _, object := range w.Objects() {
		object.Synchronize()
	}
}
