package arigo

// VersionInfo represents the version information sent by aria2
type VersionInfo struct {
	Version         string   // Version number of aria2 as a string.
	EnabledFeatures []string // Slice of enabled features. Each feature is given as a string.
}
