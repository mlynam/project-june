package settings

// Settings contains application settings
type Settings struct {
	name       string
	resolution [2]uint16
}

// Name for the application
func (s *Settings) Name() string {
	return s.name
}

// Resolution for the application
func (s *Settings) Resolution() [2]uint16 {
	return s.resolution
}
