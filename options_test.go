package arigo

import (
	"encoding/json"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestOptionFormat(t *testing.T) {
	data := []byte(`{
		"allow-overwrite": "false",
		"allow-piece-length-change": "false",
		"always-resume": "true",
		"async-dns": "true"
	}`)

	var options Options
	assert.NoError(t, json.Unmarshal(data, &options), "Couldn't unmarshal JSON")

	assert.Equal(t, Options{
		AllowOverwrite:         false,
		AllowPieceLengthChange: false,
		AlwaysResume:           true,
		AsyncDNS:               true,
	}, options)
}
