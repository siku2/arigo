// Package wsrpc provides a ReadWriteCloser which operates on a WebSocket
// connection.
package wsrpc

import (
	"github.com/gorilla/websocket"
	"io"
)

// ReadWriteCloser is a rwc based on WebSockets
type ReadWriteCloser struct {
	ws *websocket.Conn
	r  io.Reader
	w  io.WriteCloser
}

// NewReadWriteCloser creates a new rwc from a WebSocket connection
func NewReadWriteCloser(ws *websocket.Conn) ReadWriteCloser {
	return ReadWriteCloser{ws: ws}
}

// Read reads from the WebSocket into p
func (rwc *ReadWriteCloser) Read(p []byte) (n int, err error) {
	if rwc.r == nil {
		_, rwc.r, err = rwc.ws.NextReader()
		if err != nil {
			return 0, err
		}
	}

	for n = 0; n < len(p); {
		var m int
		m, err = rwc.r.Read(p[n:])
		n += m
		if err == io.EOF {
			// done
			rwc.r = nil
			break
		}
		if err != nil {
			break
		}
	}

	return
}

// Write writes the provided bytes to the WebSocket
func (rwc *ReadWriteCloser) Write(p []byte) (n int, err error) {
	if rwc.w == nil {
		rwc.w, err = rwc.ws.NextWriter(websocket.TextMessage)
		if err != nil {
			return 0, err
		}
	}

	for n = 0; n < len(p); {
		m, err := rwc.w.Write(p)
		n += m
		if err != nil {
			break
		}
	}
	if err != nil || n == len(p) {
		err = rwc.w.Close()
		rwc.w = nil
	}

	return
}

// Close the rwc and the underlying WebSocket connection
func (rwc *ReadWriteCloser) Close() error {
	var err error
	if rwc.w != nil {
		if err = rwc.w.Close(); err != nil {
			return err
		}
		rwc.w = nil
	}
	if rwc.ws != nil {
		err = rwc.ws.Close()
		rwc.ws = nil
	}
	return err
}
