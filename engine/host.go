package engine

// Host can run
type Host interface {
	Run(string)
}

// HostProviders provide core engine elements
type HostProviders struct {
	Graphics GraphicsProvider
	Platform PlatformProvider
	Manager  ManagerProvider
	Settings SettingsProvider
}

// NewHost with the given providers
func NewHost(p *HostProviders) Host {
	return &engine{p}
}
