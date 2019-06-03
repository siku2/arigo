package arigo

// VersionInfo represents the version information sent by aria2
type VersionInfo struct {
	Version         string   `json:"version"`         // Version number of aria2 as a string.
	EnabledFeatures []string `json:"enabledFeatures"` // Slice of enabled features. Each feature is given as a string.
}
