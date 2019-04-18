package arigo

import (
	"encoding/json"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestServerFormat(t *testing.T) {
	data := []byte(`[{
		"index": "1",
		"servers": [{
			"currentUri": "http://example.org/file",
			"downloadSpeed": "10467",
			"uri": "http://example.org/file"
		}]
	}]`)

	var servers []FileServers
	assert.NoError(t, json.Unmarshal(data, &servers), "Couldn't unmarshal JSON")

	assert.Len(t, servers, 1)

	server := servers[0]

	assert.Equal(t, FileServers{
		Index: 1,
		Servers: []Server{{
			CurrentURI:    "http://example.org/file",
			DownloadSpeed: 10467,
			URI:           "http://example.org/file",
		}},
	}, server)
}
