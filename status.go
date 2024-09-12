package arigo

import (
	"encoding/json"
	"time"
)

// DownloadStatus represents the status of a download.
type DownloadStatus string

const (
	// StatusActive represents currently downloading/seeding downloads
	StatusActive DownloadStatus = "active"
	// StatusWaiting represents downloads in the queue
	StatusWaiting DownloadStatus = "waiting"
	// StatusPaused represents paused downloads
	StatusPaused DownloadStatus = "paused"
	// StatusError represents downloads that were stopped because of error
	StatusError DownloadStatus = "error"
	// StatusCompleted represents stopped and completed downloads
	StatusCompleted DownloadStatus = "completed"
	// StatusRemoved represents the downloads removed by user
	StatusRemoved DownloadStatus = "removed"
)

// Status holds information for a download.
type Status struct {
	GID             string         `json:"gid"`                    // gid of the download
	Status          DownloadStatus `json:"status"`                 // Download status
	TotalLength     uint           `json:"totalLength,string"`     // Total length of the download in bytes
	CompletedLength uint           `json:"completedLength,string"` // Completed length of the download in bytes
	UploadLength    uint           `json:"uploadLength,string"`    // Uploaded length of the download in bytes

	// Hexadecimal representation of the download progress.
	// The highest bit corresponds to the piece at index 0. Any set bits indicate loaded pieces,
	// while unset bits indicate not yet loaded and/or missing pieces.
	// Any overflow bits at the end are set to zero.
	// When the download was not started yet, this will be an empty string.
	BitField      string `json:"bitfield"`
	DownloadSpeed uint   `json:"downloadSpeed,string"` // Download speed of this download measured in bytes/sec
	UploadSpeed   uint   `json:"uploadSpeed,string"`   // Upload speed of this download measured in bytes/sec
	InfoHash      string `json:"infoHash"`             // InfoHash. BitTorrent only

	// The number of seeders aria2 has connected to. BitTorrent only
	NumSeeders uint `json:"numSeeders,string"`

	// true if the local endpoint is a seeder. Otherwise false. BitTorrent only
	Seeder       bool       `json:"seeder,string"`
	PieceLength  uint       `json:"pieceLength,string"` // Piece length in bytes
	NumPieces    uint       `json:"numPieces,string"`   // The number of pieces
	Connections  uint       `json:"connections,string"` // The number of peers/servers aria2 has connected to
	ErrorCode    ExitStatus `json:"errorCode,string"`   // The code of the last error for this item, if any.
	ErrorMessage string     `json:"errorMessage"`       // The human readable error message associated to ErrorCode

	// List of GIDs which are generated as the result of this download.
	// For example, when aria2 downloads a Metalink file, it generates downloads described in the Metalink
	// (see the --follow-metalink option). This value is useful to track auto-generated downloads.
	// If there are no such downloads, this will be an empty slice
	FollowedBy []string `json:"followedBy"`

	// The reverse link for followedBy.
	// A download included in followedBy has this object’s GID in its following
	// value.
	Following string `json:"following"`

	// GID of a parent download. Some downloads are a part of another download.
	// For example, if a file in a Metalink has BitTorrent resources,
	// the downloads of “.torrent” files are parts of that parent.
	// If this download has no parent, this will be an empty string
	BelongsTo  string           `json:"belongsTo"`
	Dir        string           `json:"dir"`        // Directory to save files
	Files      []File           `json:"files"`      // Slice of files.
	BitTorrent BitTorrentStatus `json:"bittorrent"` // Information retrieved from the .torrent (file). BitTorrent only

	// The number of verified number of bytes while the files are being has
	// checked.
	// This key exists only when this download is being hash checked
	VerifiedLength uint `json:"verifiedLength,string"`

	// true if this download is waiting for the hash check in a queue.
	VerifyIntegrityPending bool `json:"verifyIntegrityPending,string"`
}

// UNIXTime is a wrapper around time.Time that marshals to a unix timestamp.
type UNIXTime struct {
	time.Time
}

// MarshalJSON converts the time to a unix timestamp
func (t UNIXTime) MarshalJSON() ([]byte, error) {
	return json.Marshal(t.Unix())
}

// UnmarshalJSON loads a unix timestamp
func (t *UNIXTime) UnmarshalJSON(data []byte) error {
	var ts int64
	err := json.Unmarshal(data, &ts)
	if err != nil {
		return err
	}

	*t = UNIXTime{time.Unix(ts, 0)}
	return nil
}

// TorrentMode represents the file mode of the torrent
type TorrentMode string

const (
	// TorrentModeSingle represents the file mode single
	TorrentModeSingle TorrentMode = "single"
	// TorrentModeMulti represents the file mode multi
	TorrentModeMulti TorrentMode = "multi"
)

// BitTorrentStatus holds information for a BitTorrent download
type BitTorrentStatus struct {
	// List of lists of announce URIs.
	// If the torrent contains announce and no announce-list,
	// announce is converted to the announce-list format
	AnnounceList [][]string           `json:"announceList"`        // List of lists of announce URIs.
	Comment      string               `json:"comment"`             // The comment of the torrent
	CreationDate UNIXTime             `json:"creationDate,string"` // The creation time of the torrent
	Mode         TorrentMode          `json:"mode"`                // File mode of the torrent
	Info         BitTorrentStatusInfo `json:"info"`                // Information from the info dictionary
}

// A BitTorrentStatusInfo holds information from the info dictionary.
type BitTorrentStatusInfo struct {
	Name string `json:"name"` // name in info dictionary
}
