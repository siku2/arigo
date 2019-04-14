package arigo

type Server struct {
	// Original URI.
	URI string
	// This is the URI currently used for downloading. If redirection is involved, currentUri and uri may differ.
	CurrentURI string
	// Download speed (byte/sec)
	DownloadSpeed uint `json:",string"`
}

type FileServers struct {
	// Index of the file, starting at 1, in the same order as files appear in the multi-file metalink.
	Index uint `json:",string"`
	// Slice of Servers
	Servers []Server
}
