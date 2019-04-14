package arigo

// TODO complete implementation

// GID provides an object oriented approach to arigo.
// Instead of calling the methods on the client directly,
// you can call them on the GID instance.
type GID struct {
	client *Client
	GID    string
}

func (gid *GID) String() string {
	return gid.GID
}

// Delete removes the download from disk as well as from aria2.
func (gid *GID) Delete() error {
	return gid.client.Delete(gid.GID)
}

// WaitForDownload waits for the download to finish.
func (gid *GID) WaitForDownload() error {
	return gid.client.WaitForDownload(gid.GID)
}

// Remove removes the download.
// If the specified download is in progress, it is first stopped.
// The status of the removed download becomes removed.
func (gid *GID) Remove() error {
	return gid.client.Remove(gid.GID)
}

// ForceRemove removes the download.
// This method behaves just like Remove() except that this method removes the download
// without performing any actions which take time, such as contacting BitTorrent trackers to
// unregister the download first.
func (gid *GID) ForceRemove() error {
	return gid.client.ForceRemove(gid.GID)
}

// Pause pauses the download.
// The status of paused download becomes paused. If the download was active,
// the download is placed in the front of the queue. While the status is paused,
// the download is not started. To change status to waiting, use the Unpause() method.
func (gid *GID) Pause() error {
	return gid.client.Pause(gid.GID)
}

// ForcePause pauses the download.
// This method behaves just like Pause() except that this method pauses downloads
// without performing any actions which take time, such as contacting BitTorrent trackers to
// unregister the download first.
func (gid *GID) ForcePause() error {
	return gid.client.ForcePause(gid.GID)
}

// Unpause changes the status of the download from paused to waiting,
// making the download eligible to be restarted.
func (gid *GID) Unpause() error {
	return gid.client.Unpause(gid.GID)
}

// TellStatus returns the progress of the download.
// If specified, the response only contains only the keys passed to the method.
// If keys is empty, the response contains all keys.
// This is useful when you just want specific keys and avoid unnecessary transfers.
func (gid *GID) TellStatus(keys ...string) (Status, error) {
	return gid.client.TellStatus(gid.GID, keys...)
}

// GetURIs returns the URIs used in the download.
// The response is a slice of URI.
func (gid *GID) GetURIs() ([]URI, error) {
	return gid.client.GetURIs(gid.GID)
}

// GetFiles returns the file list of the download.
// The response is a slice of File.
func (gid *GID) GetFiles() ([]File, error) {
	return gid.client.GetFiles(gid.GID)
}
