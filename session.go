package arigo

// SessionInfo holds the session information of aria2
type SessionInfo struct {
	// Session ID, which is generated each time when aria2 is invoked
	ID string `json:"sessionId"`
}
