# arigo
[![Go Report Card](https://goreportcard.com/badge/github.com/myanimestream/arigo)](https://goreportcard.com/report/github.com/myanimestream/arigo)
[![Documentation](https://godoc.org/github.com/github.com/myanimestream/arigo?status.svg)](http://godoc.org/github.com/myanimestream/arigo)

A client library for the aria2 RPC interface.

Features:
 - Complete implementation of all methods and their return values
 - Connection over WebSocket which makes it possible to receive events
 - Object-oriented approach for individual downloads


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
	
	// Download is an utility method which calls AddURI and 
	// then waits for the download to complete
	status, err := c.Download(arigo.URIs("https://example.org/file"), nil)
	if err != nil {
		panic(err)
	}
	
	// download complete
	fmt.Println(status.GID)
}
```