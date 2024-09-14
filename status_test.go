package arigo

import (
	"encoding/json"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestStatusFormat(t *testing.T) {
	data := []byte(`{
		"bitfield": "0000000000",
		"bittorrent": {
			"announceList": [
				[
					"http://tracker.example.com:80/announce"
				]
			],
			"comment": "Torrent downloaded from https://example.com",
			"creationDate": 1520238244,
			"info": {
				"name": "example.mkv"
			},
			"mode": "multi"
		},
		"completedLength": "901120",
		"connections": "1",
		"dir": "/downloads",
		"downloadSpeed": "15158",
		"files": [{
			"index": "1",
			"length": "34896138",
			"completedLength": "34896138",
			"path": "/downloads/file",
			"selected": "true",
			"uris": [{"status": "used", "uri": "http://example.org/file"}]
		}],
		"gid": "2089b05ecca3d829",
		"numPieces": "34",
		"pieceLength": "1048576",
		"status": "active",
		"totalLength": "34896138",
		"uploadLength": "0",
		"uploadSpeed": "0"
	}`)

	var status Status
	if err := json.Unmarshal(data, &status); err != nil {
		t.Error(err)
	}

	assert.EqualValues(t, "0000000000", status.BitField)
	assert.EqualValues(t, 901120, status.CompletedLength)
	assert.EqualValues(t, 1, status.Connections)
	assert.EqualValues(t, "/downloads", status.Dir)
	assert.EqualValues(t, 15158, status.DownloadSpeed)
	assert.EqualValues(t, "2089b05ecca3d829", status.GID)
	assert.EqualValues(t, 34, status.NumPieces)
	assert.EqualValues(t, 1048576, status.PieceLength)
	assert.EqualValues(t, StatusActive, status.Status)
	assert.EqualValues(t, 34896138, status.TotalLength)
	assert.EqualValues(t, 0, status.UploadLength)
	assert.EqualValues(t, 0, status.UploadSpeed)

	assert.EqualValues(t, 1, len(status.Files))
	file1 := status.Files[0]
	assert.EqualValues(t, 1, file1.Index)
	assert.EqualValues(t, 34896138, file1.Length)
	assert.EqualValues(t, 34896138, file1.CompletedLength)
	assert.EqualValues(t, "/downloads/file", file1.Path)
	assert.EqualValues(t, true, file1.Selected)
	assert.Equal(t, []URI{{Status: URIUsed, URI: "http://example.org/file"}}, file1.URIs)
}
