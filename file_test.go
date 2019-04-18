package arigo

import (
	"encoding/json"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestFileFormat(t *testing.T) {
	data := []byte(`[{
		"index": "1",
		"length": "34896138",
		"completedLength": "34896138",
		"path": "/downloads/file",
		"selected": "true",
		"uris": [{
			"status": "used",
			"uri": "http://example.org/file"
		}]
	}]`)

	var files []File
	assert.NoError(t, json.Unmarshal(data, &files), "couldn't unmarshal json")

	assert.Len(t, files, 1)

	file := files[0]
	assert.Equal(t, 1, file.Index)
	assert.Equal(t, uint(34896138), file.Length)
	assert.Equal(t, uint(34896138), file.CompletedLength)
	assert.Equal(t, "/downloads/file", file.Path)
	assert.Equal(t, true, file.Selected)

	uris := file.URIs
	assert.Len(t, uris, 1)
	uri := uris[0]

	assert.Equal(t, URIUsed, uri.Status)
	assert.Equal(t, "http://example.org/file", uri.URI)
}
