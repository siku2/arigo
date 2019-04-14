package arigo

import (
	"encoding/json"
	"time"
)

type StatusName string

const (
	StatusActive    StatusName = "active"
	StatusWaiting   StatusName = "waiting"
	StatusPaused    StatusName = "paused"
	StatusError     StatusName = "error"
	StatusCompleted StatusName = "completed"
	StatusRemoved   StatusName = "removed"
)

type Status struct {
	GID                    string
	Status                 StatusName
	TotalLength            uint `json:",string"`
	CompletedLength        uint `json:",string"`
	UploadLength           uint `json:",string"`
	BitField               string
	DownloadSpeed          uint `json:",string"`
	UploadSpeed            uint `json:",string"`
	InfoHash               string
	NumSeeders             uint       `json:",string"`
	Seeder                 bool       `json:",string"`
	PieceLength            uint       `json:",string"`
	NumPieces              uint       `json:",string"`
	Connections            uint       `json:",string"`
	ErrorCode              ExitStatus `json:",string"`
	ErrorMessage           string
	FollowedBy             []string
	Following              string
	BelongsTo              string
	Dir                    string
	Files                  []File
	BitTorrent             BitTorrentStatus
	VerifiedLength         uint `json:",string"`
	VerifyIntegrityPending bool `json:",string"`
}

type UNIXTime struct {
	time.Time
}

func (t UNIXTime) MarshalJSON() ([]byte, error) {
	return json.Marshal(t.Unix())
}

func (t *UNIXTime) UnmarshalJSON(data []byte) error {
	var ts int64
	err := json.Unmarshal(data, &ts)
	if err != nil {
		return err
	}

	*t = UNIXTime{time.Unix(ts, 0)}
	return nil
}

type BitTorrentStatus struct {
	AnnounceList     []URI
	Comment          string
	CreationDateUNIX UNIXTime `json:",string"`
	Mode             string
	Info             BitTorrentStatusInfo
}

type BitTorrentStatusInfo struct {
	Name string
}
