package arigo

import (
	"encoding/json"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestURIFormat(t *testing.T) {
	data := []byte(`[{
		"status": "used",
		"uri": "http://example.org/file"
	}]`)

	var uris []URI
	assert.NoError(t, json.Unmarshal(data, &uris), "couldn't unmarshal json")

	assert.Len(t, uris, 1)

	uri := uris[0]
	assert.Equal(t, URIUsed, uri.Status)
	assert.Equal(t, "http://example.org/file", uri.URI)
}
