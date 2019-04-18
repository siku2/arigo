package arigo

import (
	"encoding/json"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestSessionFormat(t *testing.T) {
	data := []byte(`{
		"sessionId": "cd6a3bc6a1de28eb5bfa181e5f6b916d44af31a9"
	}`)

	var info SessionInfo
	assert.NoError(t, json.Unmarshal(data, &info), "Couldn't unmarshal JSON")

	assert.Equal(t, SessionInfo{
		ID: "cd6a3bc6a1de28eb5bfa181e5f6b916d44af31a9",
	}, info)
}
