package arigo

type Stats struct {
	// Overall download speed (byte/sec).
	DownloadSpeed uint `json:",string"`
	// Overall upload speed(byte/sec).
	UploadSpeed uint `json:",string"`
	// The number of active downloads.
	NumActive uint `json:",string"`
	// The number of waiting downloads.
	NumWaiting uint `json:",string"`
	// The number of stopped downloads in the current session.
	// This value is capped by the MaxDownloadResult option.
	NumStopped uint `json:",string"`
	// The number of stopped downloads in the current session and not capped by the MaxDownloadResult option.
	NumStoppedTotal uint `json:",string"`
}
