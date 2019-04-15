package arigo

// Stats holds aria2 statistics
type Stats struct {
	DownloadSpeed uint `json:",string"` // Overall download speed (byte/sec).
	UploadSpeed   uint `json:",string"` // Overall upload speed(byte/sec).
	NumActive     uint `json:",string"` // The number of active downloads.
	NumWaiting    uint `json:",string"` // The number of waiting downloads.

	// The number of stopped downloads in the current session.
	// This value is capped by the MaxDownloadResult option.
	NumStopped uint `json:",string"`

	// The number of stopped downloads in the current session and not capped by the MaxDownloadResult option.
	NumStoppedTotal uint `json:",string"`
}
