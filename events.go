package arigo

// DownloadEvent represents the event emitted by aria2 concerning downloads.
// It only contains the gid of the download.
type DownloadEvent struct {
	GID string
}

func (e *DownloadEvent) String() string {
	return e.GID
}
