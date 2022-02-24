# arigo
[![Build Status](https://github.com/myanimestream/arigo/actions/workflows/build.yml/badge.svg)](https://github.com/myanimestream/arigo/actions/workflows/build.yml)
[![Go Report Card](https://goreportcard.com/badge/github.com/myanimestream/arigo)](https://goreportcard.com/report/github.com/myanimestream/arigo)
[![Documentation](https://godoc.org/github.com/github.com/myanimestream/arigo?status.svg)](http://godoc.org/github.com/myanimestream/arigo)

A client library for the aria2 RPC interface.

Features:
- Complete implementation of all supported methods.
- Bidirectional communication over WebSocket which makes it 
possible to receive events and know when a download has completed.


arigo currently doesn't start and manage the aria2c process for you.
To start aria2 use the following command:
```bash
aria2c --enable-rpc --rpc-listen-all
```

If aria2 is not installed then head on to
https://aria2.github.io/ and follow the instructions there.

## Example
```go
package main

import (
	"fmt"
	"github.com/myanimestream/arigo"
)

func main() {
	c, err := arigo.Dial("ws://localhost:6800/jsonrpc", "")
	if err != nil {
		panic(err)
	}
	
	// Download is a utility method which calls AddURI and 
	// then waits for the download to complete
	status, err := c.Download(arigo.URIs("https://example.org/file"), nil)
	if err != nil {
		panic(err)
	}
	
	// download complete
	fmt.Println(status.GID)
}
```