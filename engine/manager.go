package engine

// ManagerProvider loads scenes
type ManagerProvider interface {
	New(Settings) Manager
}

// Manager ties together world and scene logic
type Manager interface {
	LoadStartScene(string) (Scene, World)
}
