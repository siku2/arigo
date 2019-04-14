package arigo

type VersionInfo struct {
	// Version number of aria2 as a string.
	Version string
	// Slice of enabled features. Each feature is given as a string.
	EnabledFeatures []string
}
