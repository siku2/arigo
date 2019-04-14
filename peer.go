package arigo

type Peer struct {
	// Percent-encoded peer ID.
	PeerId string
	// IP address of the peer.
	IP string
	// Port number of the peer
	Port uint16 `json:",string"`
	// Hexadecimal representation of the download progress of the peer.
	// The highest bit corresponds to the piece at index 0.
	// Set bits indicate the piece is available and unset bits indicate the piece is missing.
	// Any spare bits at the end are set to zero.
	BitField string
	// true if aria2 is choking the peer. Otherwise false.
	AmChoking bool `json:",string"`
	// true if the peer is choking aria2. Otherwise false.
	PeerChoking bool `json:",string"`
	// Download speed (byte/sec) that this client obtains from the peer
	DownloadSpeed uint `json:",string"`
	// Upload speed (byte/sec) that this client uploads to the peer
	UploadSpeed uint `json:",string"`
	// true if this peer is a seeder. Otherwise false
	Seeder bool `json:",string"`
}
