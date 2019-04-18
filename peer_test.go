package arigo

import (
	"encoding/json"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestPeerFormat(t *testing.T) {
	data := []byte(`[{
		"amChoking": "true",
		"bitfield": "ffffffffffffffffffffffffffffffffffffffff",
		"downloadSpeed": "10602",
		"ip": "10.0.0.9",
		"peerChoking": "false",
		"peerId": "aria2%2F1%2E10%2E5%2D%87%2A%EDz%2F%F7%E6",
		"port": "6881",
		"seeder": "true",
		"uploadSpeed": "0"
	}, {
		"amChoking": "false",
		"bitfield": "ffffeff0fffffffbfffffff9fffffcfff7f4ffff",
		"downloadSpeed": "8654",
		"ip": "10.0.0.30",
		"peerChoking": "false",
		"peerId": "bittorrent client758",
		"port": "37842",
		"seeder": "false",
		"uploadSpeed": "6890"
	}]`)

	var peers []Peer
	assert.NoError(t, json.Unmarshal(data, &peers), "Couldn't unmarshal JSON")

	assert.Len(t, peers, 2)

	firstPeer := peers[0]
	assert.Equal(t, Peer{
		AmChoking:     true,
		BitField:      "ffffffffffffffffffffffffffffffffffffffff",
		DownloadSpeed: uint(10602),
		IP:            "10.0.0.9",
		PeerChoking:   false,
		ID:            "aria2%2F1%2E10%2E5%2D%87%2A%EDz%2F%F7%E6",
		Port:          6881,
		Seeder:        true,
		UploadSpeed:   0,
	}, firstPeer)

	secondPeer := peers[1]
	assert.Equal(t, Peer{
		AmChoking:     false,
		BitField:      "ffffeff0fffffffbfffffff9fffffcfff7f4ffff",
		DownloadSpeed: uint(8654),
		IP:            "10.0.0.30",
		PeerChoking:   false,
		ID:            "bittorrent client758",
		Port:          37842,
		Seeder:        false,
		UploadSpeed:   6890,
	}, secondPeer)
}
