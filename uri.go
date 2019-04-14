package arigo

type URIStatus string

const (
	URIUsed    URIStatus = "used"
	URIWaiting URIStatus = "waiting"
)

type URI struct {
	URI    string
	Status URIStatus
}
