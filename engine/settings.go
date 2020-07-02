package engine

// SettingsProvider gets the engine settings
type SettingsProvider interface {
	New() Settings
}

// Settings for the engine
type Settings interface {
	Name() string
	Resolution() [2]uint16
}
