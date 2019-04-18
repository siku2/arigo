package arigo

import (
	"encoding/json"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestVersionFormat(t *testing.T) {
	data := []byte(`{
		"enabledFeatures": [
			"Async DNS",
			"BitTorrent",
			"Firefox3 Cookie",
			"GZip",
			"HTTPS",
			"Message Digest",
			"Metalink",
			"XML-RPC"
		],
		"version": "1.11.0"
	}`)

	var version VersionInfo
	assert.NoError(t, json.Unmarshal(data, &version), "Couldn't unmarshal JSON")

	assert.Equal(t, VersionInfo{
		EnabledFeatures: []string{
			"Async DNS",
			"BitTorrent",
			"Firefox3 Cookie",
			"GZip",
			"HTTPS",
			"Message Digest",
			"Metalink",
			"XML-RPC",
		},
		Version: "1.11.0",
	}, version)
}
