package arigo

// Server represents an endpoint which data is being downloaded from
type Server struct {
	URI string `json:"uri"` // Original URI.

	// This is the URI currently used for downloading.
	// If redirection is involved, currentUri and uri may differ.
	CurrentURI    string `json:"currentUri"`
	DownloadSpeed uint   `json:"downloadSpeed,string"` // Download speed (byte/sec)
}

// FileServers holds the servers of a file
type FileServers struct {
	// Index of the file, starting at 1,
	// in the same order as files appear in the multi-file metalink.
	Index   uint     `json:"index,string"`
	Servers []Server `json:"servers"` // Slice of Servers which are used for the file
}
