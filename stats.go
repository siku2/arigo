package arigo

// Stats holds aria2 statistics
type Stats struct {
	DownloadSpeed uint `json:"downloadSpeed,string"` // Overall download speed (byte/sec).
	UploadSpeed   uint `json:"uploadSpeed,string"`   // Overall upload speed(byte/sec).
	NumActive     uint `json:"numActive,string"`     // The number of active downloads.
	NumWaiting    uint `json:"numWaiting,string"`    // The number of waiting downloads.

	// The number of stopped downloads in the current session.
	// This value is capped by the MaxDownloadResult option.
	NumStopped uint `json:"numStopped,string"`

	// The number of stopped downloads in the current session and not capped by the MaxDownloadResult option.
	NumStoppedTotal uint `json:"numStoppedTotal,string"`
}
