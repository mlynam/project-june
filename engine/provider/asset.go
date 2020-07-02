package provider

// Asset provider loads scenes
type Asset interface {
	LoadScene(string)
}
