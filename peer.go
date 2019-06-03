package arigo

// Peer represents a torrent peer
type Peer struct {
	ID   string `json:"peerId"`      // Percent-encoded peer ID.
	IP   string `json:"ip"`          // IP address of the peer.
	Port uint16 `json:"port,string"` // Port number of the peer

	// Hexadecimal representation of the download progress of the peer.
	// The highest bit corresponds to the piece at index 0.
	// Set bits indicate the piece is available and unset bits indicate the piece is missing.
	// Any spare bits at the end are set to zero.
	BitField      string `json:"bitfield"`
	AmChoking     bool   `json:"amChoking,string"`     // true if aria2 is choking the peer. Otherwise false.
	PeerChoking   bool   `json:"peerChoking,string"`   // true if the peer is choking aria2. Otherwise false.
	DownloadSpeed uint   `json:"downloadSpeed,string"` // Download speed (byte/sec) that this client obtains from the peer
	UploadSpeed   uint   `json:"uploadSpeed,string"`   // Upload speed (byte/sec) that this client uploads to the peer
	Seeder        bool   `json:"seeder,string"`        // true if this peer is a seeder. Otherwise false
}
