package arigo

// GID provides an object oriented approach to arigo.
// Instead of calling the methods on the client directly,
// you can call them on the GID instance.
type GID struct {
	client *Client
	GID    string // gid of the download
}

func (gid *GID) String() string {
	return gid.GID
}

// Subscribe subscribes to the given event but only dispatches events concerning
// this GID.
func (gid *GID) Subscribe(evtType EventType, listener EventListener) UnsubscribeFunc {
	return gid.client.Subscribe(evtType, func(event *DownloadEvent) {
		if event.GID == gid.GID {
			listener(event)
		}
	})
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

// GetPeers returns a list of peers of the download denoted by gid.
// This method is for BitTorrent only.
// The response is a slice of Peers.
func (gid *GID) GetPeers() ([]Peer, error) {
	return gid.client.GetPeers(gid.GID)
}

// GetServers returns currently connected HTTP(S)/FTP/SFTP servers of the download denoted by gid.
// Returns a slice of FileServers.
func (gid *GID) GetServers() ([]FileServers, error) {
	return gid.client.GetServers(gid.GID)
}

// ChangePosition changes the position of the download denoted by gid in the queue.
//
// If how is SetPositionStart, it moves the download to a position relative to the beginning of the queue.
// If how is SetPositionRelative, it moves the download to a position relative to the current position.
// If how is SetPositionEnd, it moves the download to a position relative to the end of the queue.
// If the destination position is less than 0 or beyond the end of the queue,
// it moves the download to the beginning or the end of the queue respectively.
//
// The response is an integer denoting the resulting position.
func (gid *GID) ChangePosition(pos int, how PositionSetBehaviour) (int, error) {
	return gid.client.ChangePosition(gid.GID, pos, how)
}

// ChangeURIAt removes the URIs in delUris from and appends the URIs in addUris to download denoted by gid.
// A download can contain multiple files and URIs are attached to each file.
// fileIndex is used to select which file to remove/attach given URIs. fileIndex is 1-based.
// position is used to specify where URIs are inserted in the existing waiting URI list. position is 0-based.
// When position is nil, URIs are appended to the back of the list.
//
// This method first executes the removal and then the addition.
// position is the position after URIs are removed, not the position when this method is called.
// When removing an URI, if the same URIs exist in download, only one of them is removed for each URI in delUris.
//
// Returns two integers.
// The first integer is the number of URIs deleted.
// The second integer is the number of URIs added.
func (gid *GID) ChangeURIAt(fileIndex uint, delURIs []string, addURIs []string, position uint) (uint, uint, error) {
	return gid.client.ChangeURIAt(gid.GID, fileIndex, delURIs, addURIs, position)
}

// ChangeURI removes the URIs in delUris from and appends the URIs in addUris to download denoted by gid.
// A download can contain multiple files and URIs are attached to each file.
// fileIndex is used to select which file to remove/attach given URIs. fileIndex is 1-based.
// position is used to specify where URIs are inserted in the existing waiting URI list. position is 0-based.
// URIs are appended to the back of the list.
//
// This method first executes the removal and then the addition.
// position is the position after URIs are removed, not the position when this method is called.
// When removing an URI, if the same URIs exist in download, only one of them is removed for each URI in delUris.
//
// Returns two integers.
// The first integer is the number of URIs deleted.
// The second integer is the number of URIs added.
func (gid *GID) ChangeURI(fileIndex uint, delURIs []string, addURIs []string) (uint, uint, error) {
	return gid.client.ChangeURI(gid.GID, fileIndex, delURIs, addURIs)
}

// GetOptions returns Options of the download denoted by gid.
// Note that this method does not return options which have no default value and have not been set on the command-line,
// in configuration files or RPC methods.
func (gid *GID) GetOptions() (Options, error) {
	return gid.client.GetOptions(gid.GID)
}

// ChangeOptions changes options of the download denoted by gid dynamically.
//
// Except for following options, all options are available:
// 	- DryRun
//  - MetalinkBaseURI
//  - ParameterizedURI
//  - Pause
//  - PieceLength
//  - RPCSaveUploadMetadata
//
// Except for the following options, changing the other options of active download makes it restart
// (restart itself is managed by aria2, and no user intervention is required):
// 	- BtMaxPeers
// 	- BtRequestPeerSpeedLimit
// 	- BtRemoveUnselectedFile
// 	- ForceSave
// 	- MaxDownloadLimit
// 	- MaxUploadLimit
func (gid *GID) ChangeOptions(changes Options) error {
	return gid.client.ChangeOptions(gid.GID, changes)
}

// RemoveDownloadResult removes a completed/error/removed download denoted by gid from memory.
func (gid *GID) RemoveDownloadResult() error {
	return gid.client.RemoveDownloadResult(gid.GID)
}
