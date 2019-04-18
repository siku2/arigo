package arigo

import (
	"encoding/json"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestStatsFormat(t *testing.T) {
	data := []byte(`{
		"downloadSpeed": "21846",
		"numActive": "2",
		"numStopped": "0",
		"numWaiting": "0",
		"uploadSpeed": "0"
	}`)

	var stats Stats
	assert.NoError(t, json.Unmarshal(data, &stats), "Couldn't unmarshal JSON")

	assert.Equal(t, Stats{
		DownloadSpeed: 21846,
		NumActive:     2,
		NumStopped:    0,
		NumWaiting:    0,
		UploadSpeed:   0,
	}, stats)
}
