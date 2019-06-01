package arigo

import "fmt"

// Dial is a convenience method which connects to an aria2 RPC interface.
// It establishes a WebSocket connection to the given url and passes it
// to the NewClient() method. The client is also started.
func ExampleClient() {
	// arigo uses a WebSocket connection to communicate with the aria2 RPC.
	// This makes it possible to receive download events.
	// authToken is the secret string set on the aria2 server. Setting it to an empty string
	// indicates that no password should be used.
	client, err := Dial("ws://localhost:6800/jsonrpc", "")

	// err is returned if no connection to the RPC interface could be established.
	if err != nil {
		panic(err)
	}

	// client is now connected and can be used
	gid, err := client.AddURI(URIs("https://example.org/file"), nil)
	if err != nil {
		panic(err)
	}

	// gid isn't just a string but a GID instance.
	// this makes it possible to use all the client methods you would normally
	// pass a gid to directly on the GID instance.
	// WaitForDownload() is a special method which waits for the download to complete
	// (using the aria2 events)
	if err = gid.WaitForDownload(); err != nil {
		panic(err)
	}

	// Get the status of the now completed download.
	// The TellStatus() method accepts the keys of the Status struct
	// which are to be requested. Please note that the keys must match
	// the aria2 keys, which may differ from the keys used in the golang
	// representation.
	// For the aria2 status documentation see: https://aria2.github.io/manual/en/html/aria2c.html#aria2.tellStatus
	status, err := gid.TellStatus("status")
	if err != nil {
		panic(err)
	}

	fmt.Println(status.Status)
}
