package arigo

// URIStatus represents the status of an uri.
type URIStatus string

const (
	// URIUsed represents the state of the uri being used
	URIUsed URIStatus = "used"
	// URIWaiting represents the state of the uri waiting in the queue
	URIWaiting URIStatus = "waiting"
)

// URI represents a uri used in a download
type URI struct {
	URI    string    `json:"uri"`
	Status URIStatus `json:"status"` // Status of the uri
}
