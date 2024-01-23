package arigo

import (
	"context"
	"encoding/base64"
	"encoding/json"
	"errors"
	"net/http"
	"os"

	"github.com/cenkalti/rpc2"
	"github.com/cenkalti/rpc2/jsonrpc"
	"github.com/gorilla/websocket"
	"github.com/siku2/arigo/internal/pkg/wsrpc"
	"github.com/siku2/arigo/pkg/aria2proto"
)

const (
	// QueueEndPosition represents the end of the download queue.
	QueueEndPosition = ^uint(0)
)

var (
	// ErrDownloadError represents an error that occurred during the download
	ErrDownloadError = errors.New("download encountered error")
	// ErrDownloadStopped is the error returned when a download is stopped
	ErrDownloadStopped = errors.New("download stopped")
)

// URIs creates a string slice from the given uris.
// This is a convenience function for the various client
// methods that accept a slice of URIs (strings).
func URIs(uris ...string) []string {
	return uris
}

// Client represents a connection to an aria2 rpc interface over websocket.
type Client struct {
	rpcClient *rpc2.Client
	closed    bool

	authToken string

	evtTarget eventTarget
}

// NewClient creates a new client.
// The client needs to be manually ran
// using the Run method.
func NewClient(rpcClient *rpc2.Client, authToken string) *Client {
	client := &Client{
		rpcClient: rpcClient,
		authToken: authToken,
		closed:    false,
	}

	rpcClient.Handle(aria2proto.OnDownloadStart, client.onDownloadStart)
	rpcClient.Handle(aria2proto.OnDownloadPause, client.onDownloadPause)
	rpcClient.Handle(aria2proto.OnDownloadStop, client.onDownloadStop)
	rpcClient.Handle(aria2proto.OnDownloadComplete, client.onDownloadComplete)
	rpcClient.Handle(aria2proto.OnDownloadError, client.onDownloadError)
	rpcClient.Handle(aria2proto.OnBTDownloadComplete, client.onBTDownloadComplete)

	return client
}

// DialContext creates a new connection to an aria2 rpc interface.
// It returns a new client.
func DialContext(ctx context.Context, url string, authToken string) (client *Client, err error) {
	dialer := websocket.Dialer{}

	ws, _, err := dialer.DialContext(ctx, url, http.Header{})
	if err != nil {
		return
	}

	rwc := wsrpc.NewReadWriteCloser(ws)
	codec := jsonrpc.NewJSONCodec(&rwc)
	rpcClient := rpc2.NewClientWithCodec(codec)

	client = NewClient(rpcClient, authToken)
	go client.Run()

	return
}

// Dial creates a new connection to an aria2 rpc interface.
// It returns a new client.
func Dial(url string, authToken string) (client *Client, err error) {
	return DialContext(context.Background(), url, authToken)
}

// Run runs the underlying rpcClient.
// There's no need to call this if the client
// was created using the Dial function.
func (c *Client) Run() {
	c.rpcClient.Run()
}

// Close closes the connection to the aria2 rpc interface.
// The client becomes unusable after that point.
func (c *Client) Close() error {
	c.closed = true

	return c.rpcClient.Close()
}

func (c *Client) onDownloadStart(_ *rpc2.Client, event *DownloadEvent, _ *interface{}) error {
	c.evtTarget.Dispatch(StartEvent, event)
	return nil
}
func (c *Client) onDownloadPause(_ *rpc2.Client, event *DownloadEvent, _ *interface{}) error {
	c.evtTarget.Dispatch(PauseEvent, event)
	return nil
}
func (c *Client) onDownloadStop(_ *rpc2.Client, event *DownloadEvent, _ *interface{}) error {
	c.evtTarget.Dispatch(StopEvent, event)
	return nil
}
func (c *Client) onDownloadComplete(_ *rpc2.Client, event *DownloadEvent, _ *interface{}) error {
	c.evtTarget.Dispatch(CompleteEvent, event)
	return nil
}
func (c *Client) onDownloadError(_ *rpc2.Client, event *DownloadEvent, _ *interface{}) error {
	c.evtTarget.Dispatch(ErrorEvent, event)
	return nil
}
func (c *Client) onBTDownloadComplete(_ *rpc2.Client, event *DownloadEvent, _ *interface{}) error {
	c.evtTarget.Dispatch(BTCompleteEvent, event)
	return nil
}

// Subscribe registers the given listener for an event.
// The listener will be called every time the event occurs.
func (c *Client) Subscribe(evtType EventType, listener EventListener) UnsubscribeFunc {
	return c.evtTarget.Subscribe(evtType, listener)
}

// WaitForDownload waits for a download denoted by its gid to finish.
func (c *Client) WaitForDownload(gid string) error {
	channel := make(chan error, 1)

	sendResponse := func(err error) EventListener {
		return func(ev *DownloadEvent) {
			if ev.GID == gid {
				channel <- err
			}
		}
	}

	stopUnsub := c.Subscribe(StopEvent, sendResponse(ErrDownloadStopped))
	completeUnsub := c.Subscribe(CompleteEvent, sendResponse(nil))
	errUnsub := c.Subscribe(ErrorEvent, sendResponse(ErrDownloadError))

	err := <-channel

	stopUnsub()
	completeUnsub()
	errUnsub()

	return err
}

// Download adds a new download and waits for it to complete.
// It returns the status of the finished download.
func (c *Client) Download(uris []string, options *Options) (status Status, err error) {
	return c.DownloadWithContext(context.Background(), uris, options)
}

// DownloadWithContext adds a new download and waits for it to complete.
// The passed context can be used to cancel the download.
// It returns the status of the finished download.
func (c *Client) DownloadWithContext(ctx context.Context, uris []string, options *Options) (status Status, err error) {
	gid, err := c.AddURI(uris, options)
	if err != nil {
		return
	}

	downloadDone := make(chan error, 1)

	go func() {
		downloadDone <- gid.WaitForDownload()
	}()

	select {
	case <-downloadDone:
		status, err = gid.TellStatus()
		if err != nil {
			return
		}
	case <-ctx.Done():
		_ = gid.Delete()
		err = ctx.Err()
	}

	return
}

// Delete removes the download denoted by gid and deletes all corresponding files.
// This is not an aria2 method.
func (c *Client) Delete(gid string) (err error) {
	err = c.Remove(gid)
	if err != nil {
		return
	}

	files, err := c.GetFiles(gid)
	if err == nil {
		for _, file := range files {
			_ = os.Remove(file.Path)
		}
	}

	return
}

// GetGID creates a GID struct which you can use to interact with the download directly
func (c *Client) GetGID(gid string) GID {
	return GID{c, gid}
}

func (c *Client) getArgs(args ...interface{}) []interface{} {
	if c.authToken == "" {
		if len(args) != 0 {
			return args
		}
		return []interface{}{}
	} else {
		tokenArg := "token:" + c.authToken
		return append([]interface{}{tokenArg}, args...)
	}
}

// AddURIAtPosition adds a new download at a specific position in the queue.
// uris is a slice of HTTP/FTP/SFTP/BitTorrent URIs pointing to the same resource.
// If you mix URIs pointing to different resources,
// then the download may fail or be corrupted without aria2 complaining.
//
// When adding BitTorrent Magnet URIs, uris must have only one element and it should be BitTorrent Magnet URI.
//
// The new download will be inserted at position in the waiting queue.
// If position is nil or position is larger than the current size of the queue,
// the new download is appended to the end of the queue.
//
// This method returns the GID of the newly registered download.
func (c *Client) AddURIAtPosition(uris []string, position uint, options *Options) (GID, error) {
	args := c.getArgs(uris)

	if options != nil {
		args = append(args, options)
	}

	if position != QueueEndPosition {
		args = append(args, position)
	}

	var reply string
	err := c.rpcClient.Call(aria2proto.AddURI, args, &reply)

	return c.GetGID(reply), err
}

// AddURI adds a new download.
// uris is a slice of HTTP/FTP/SFTP/BitTorrent URIs (strings) pointing to the same resource.
// If you mix URIs pointing to different resources,
// then the download may fail or be corrupted without aria2 complaining.
//
// When adding BitTorrent Magnet URIs, uris must have only one element and it should be BitTorrent Magnet URI.
//
// The new download is appended to the end of the queue.
//
// This method returns the GID of the newly registered download.
func (c *Client) AddURI(uris []string, options *Options) (GID, error) {
	return c.AddURIAtPosition(uris, QueueEndPosition, options)
}

// AddTorrentAtPosition adds a BitTorrent download at a specific position in the queue.
// If you want to add a BitTorrent Magnet URI, use the AddURI() method instead.
// torrent must be the contents of the “.torrent” file.
// uris is an array of URIs (string). uris is used for Web-seeding.
//
// For single file torrents, the URI can be a complete URI pointing to the resource;
// if URI ends with /, name in torrent file is added. For multi-file torrents,
// name and path in torrent are added to form a URI for each file.
//
// The new download will be inserted at position in the waiting queue.
// If position is nil or position is larger than the current size of the queue,
// the new download is appended to the end of the queue.
//
// This method returns the GID of the newly registered download.
func (c *Client) AddTorrentAtPosition(torrent []byte, uris []string, position uint, options *Options) (GID, error) {
	encodedTorrent := base64.StdEncoding.EncodeToString(torrent)
	args := c.getArgs(encodedTorrent, uris)

	if options != nil {
		args = append(args, options)
	}

	if position != QueueEndPosition {
		args = append(args, position)
	}

	var reply string
	err := c.rpcClient.Call(aria2proto.AddTorrent, args, &reply)

	return c.GetGID(reply), err
}

// AddTorrent adds a BitTorrent download by uploading a “.torrent” file.
// If you want to add a BitTorrent Magnet URI, use the AddURI() method instead.
// torrent must be the contents of the “.torrent” file.
// uris is an array of URIs (string). uris is used for Web-seeding.
//
// For single file torrents, the URI can be a complete URI pointing to the resource;
// if URI ends with /, name in torrent file is added. For multi-file torrents,
// name and path in torrent are added to form a URI for each file.
//
// The new download is appended to the end of the queue.
//
// This method returns the GID of the newly registered download.
func (c *Client) AddTorrent(torrent []byte, uris []string, options *Options) (GID, error) {
	return c.AddTorrentAtPosition(torrent, uris, QueueEndPosition, options)
}

// AddMetalinkAtPosition adds a Metalink download at a specific position in the queue by uploading a “.metalink” file.
// metalink is the contents of the “.metalink” file.
//
// The new download will be inserted at position in the waiting queue.
// If position is nil or position is larger than the current size of the queue,
// the new download is appended to the end of the queue.
//
// This method returns an array of GIDs of newly registered downloads.
func (c *Client) AddMetalinkAtPosition(metalink []byte, position uint, options *Options) ([]GID, error) {
	encodedMetalink := base64.StdEncoding.EncodeToString(metalink)
	args := c.getArgs(encodedMetalink)

	if options != nil {
		args = append(args, options)
	}

	if position != QueueEndPosition {
		args = append(args, position)
	}

	var reply []string
	err := c.rpcClient.Call(aria2proto.AddMetalink, args, &reply)

	gids := make([]GID, len(reply))
	for _, rawGID := range reply {
		gids = append(gids, c.GetGID(rawGID))
	}

	return gids, err
}

// AddMetalink adds a Metalink download by uploading a “.metalink” file.
// metalink is the contents of the “.metalink” file.
//
// The new download is appended to the end of the queue.
//
// This method returns an array of GIDs of newly registered downloads.
func (c *Client) AddMetalink(metalink []byte, options *Options) ([]GID, error) {
	return c.AddMetalinkAtPosition(metalink, QueueEndPosition, options)
}

// Remove removes the download denoted by gid.
// If the specified download is in progress, it is first stopped.
// The status of the removed download becomes removed.
func (c *Client) Remove(gid string) error {
	return c.rpcClient.Call(aria2proto.Remove, c.getArgs(gid), nil)
}

// ForceRemove removes the download denoted by gid.
// This method behaves just like Remove() except that this method removes the download
// without performing any actions which take time, such as contacting BitTorrent trackers to
// unregister the download first.
func (c *Client) ForceRemove(gid string) error {
	return c.rpcClient.Call(aria2proto.ForceRemove, c.getArgs(gid), nil)
}

// Pause pauses the download denoted by gid.
// The status of paused download becomes paused. If the download was active,
// the download is placed in the front of the queue. While the status is paused,
// the download is not started. To change status to waiting, use the Unpause() method.
func (c *Client) Pause(gid string) error {
	return c.rpcClient.Call(aria2proto.Pause, c.getArgs(gid), nil)
}

// PauseAll is equal to calling Pause() for every active/waiting download.
func (c *Client) PauseAll() error {
	return c.rpcClient.Call(aria2proto.PauseAll, c.getArgs(), nil)
}

// ForcePause pauses the download denoted by gid.
// This method behaves just like Pause() except that this method pauses downloads
// without performing any actions which take time, such as contacting BitTorrent trackers to
// unregister the download first.
func (c *Client) ForcePause(gid string) error {
	return c.rpcClient.Call(aria2proto.ForcePause, c.getArgs(gid), nil)
}

// ForcePauseAll is equal to calling ForcePause() for every active/waiting download.
func (c *Client) ForcePauseAll() error {
	return c.rpcClient.Call(aria2proto.ForcePauseAll, c.getArgs(), nil)
}

// Unpause changes the status of the download denoted by gid from paused to waiting,
// making the download eligible to be restarted.
func (c *Client) Unpause(gid string) error {
	return c.rpcClient.Call(aria2proto.Unpause, c.getArgs(gid), nil)
}

// UnpauseAll is equal to calling Unpause() for every paused download.
func (c *Client) UnpauseAll() error {
	return c.rpcClient.Call(aria2proto.UnpauseAll, c.getArgs(), nil)
}

// TellStatus returns the progress of the download denoted by gid.
//
// If specified, the returned Status only contains the keys passed to the method.
// This is useful when you just want specific keys and avoid unnecessary transfers.
func (c *Client) TellStatus(gid string, keys ...string) (Status, error) {
	var reply Status
	// convert nil to empty slice, prevent error The parameter at 1 has wrong type.
	if len(keys) == 0 {
		keys = make([]string, 0)
	}
	err := c.rpcClient.Call(aria2proto.TellStatus, c.getArgs(gid, keys), &reply)

	return reply, err
}

// GetURIs returns the URIs used in the download denoted by gid.
// The response is a slice of URIs.
func (c *Client) GetURIs(gid string) ([]URI, error) {
	var reply []URI
	err := c.rpcClient.Call(aria2proto.GetURIs, c.getArgs(gid), &reply)

	return reply, err
}

// GetFiles returns the file list of the download denoted by gid.
// The response is a slice of Files.
func (c *Client) GetFiles(gid string) ([]File, error) {
	var reply []File
	err := c.rpcClient.Call(aria2proto.GetFiles, c.getArgs(gid), &reply)

	return reply, err
}

// GetPeers returns a list of peers of the download denoted by gid.
// This method is for BitTorrent only.
// The response is a slice of Peers.
func (c *Client) GetPeers(gid string) ([]Peer, error) {
	var reply []Peer
	err := c.rpcClient.Call(aria2proto.GetPeers, c.getArgs(gid), &reply)

	return reply, err
}

// GetServers returns currently connected HTTP(S)/FTP/SFTP servers of the download denoted by gid.
// Returns a slice of FileServers.
func (c *Client) GetServers(gid string) ([]FileServers, error) {
	var reply []FileServers
	err := c.rpcClient.Call(aria2proto.GetServers, c.getArgs(gid), &reply)

	return reply, err
}

// TellActive returns a slice of active downloads represented by their Status.
// keys does the same as in the TellStatus() method.
func (c *Client) TellActive(keys ...string) ([]Status, error) {
	var reply []Status
	err := c.rpcClient.Call(aria2proto.TellActive, c.getArgs(keys), &reply)

	return reply, err
}

// TellWaiting returns a slice of waiting downloads including paused ones represented by their Status.
//
// offset is an integer and specifies the offset from the download waiting at the front.
// num is an integer and specifies the max. number of downloads to be returned.
//
// If offset is a positive integer, this method returns downloads in the range of [offset, offset + num).
// offset can be a negative integer. offset == -1 points last download in the waiting queue and offset == -2 points to
// the download before the last download, and so on. The returned statuses are in reversed order then.
//
// If specified, the returned Statuses only contain the keys passed to the method.
func (c *Client) TellWaiting(offset int, num uint, keys ...string) ([]Status, error) {
	var reply []Status
	err := c.rpcClient.Call(aria2proto.TellWaiting, c.getArgs(offset, num, keys), &reply)

	return reply, err
}

// TellStopped returns a slice of stopped downloads represented by their Status.
//
// offset is an integer and specifies the offset from the download waiting at the front.
// num is an integer and specifies the max. number of downloads to be returned.
//
// If offset is a positive integer, this method returns downloads in the range of [offset, offset + num).
// offset can be a negative integer. offset == -1 points last download in the waiting queue and offset == -2 points to
// the download before the last download, and so on. The returned statuses are in reversed order then.
//
// If specified, the returned Statuses only contain the keys passed to the method.
func (c *Client) TellStopped(offset int, num uint, keys ...string) ([]Status, error) {
	var reply []Status
	err := c.rpcClient.Call(aria2proto.TellStopped, c.getArgs(offset, num, keys), &reply)

	return reply, err
}

// PositionSetBehaviour determines how a position is to be interpreted
type PositionSetBehaviour string

const (
	// SetPositionStart sets the position relative to the start
	SetPositionStart PositionSetBehaviour = "POS_SET"

	// SetPositionEnd sets the position relative to the end
	SetPositionEnd PositionSetBehaviour = "POS_END"

	// SetPositionRelative sets the position relative to the current position
	SetPositionRelative PositionSetBehaviour = "POS_CUR"
)

// ChangePosition changes the position of the download denoted by gid in the queue.
//
// If how is SetPositionStart, it moves the download to a position relative to the beginning of the queue.
// If how is SetPositionRelative, it moves the download to a position relative to the current position.
// If how is SetPositionEnd, it moves the download to a position relative to the end of the queue.
// If the destination position is less than 0 or beyond the end of the queue,
// it moves the download to the beginning or the end of the queue respectively.
//
// The response is an integer denoting the resulting position.
func (c *Client) ChangePosition(gid string, pos int, how PositionSetBehaviour) (int, error) {
	args := c.getArgs(gid, pos)
	if how != "" {
		args = append(args, how)
	}

	var reply int
	err := c.rpcClient.Call(aria2proto.ChangePosition, args, &reply)

	return reply, err
}

// ChangeURIAt removes the URIs in delUris from and appends the URIs in addUris to download denoted by gid.
// A download can contain multiple files and URIs are attached to each file.
// fileIndex is used to select which file to remove/attach given URIs. fileIndex is 1-based.
// position is used to specify where URIs are inserted in the existing waiting URI list. position is 0-based.
//
// This method first executes the removal and then the addition.
// position is the position after URIs are removed, not the position when this method is called.
// When removing an URI, if the same URIs exist in download, only one of them is removed for each URI in delUris.
//
// Returns two integers.
// The first integer is the number of URIs deleted.
// The second integer is the number of URIs added.
func (c *Client) ChangeURIAt(gid string, fileIndex uint, delURIs []string, addURIs []string, position uint) (uint, uint, error) {
	args := c.getArgs(gid, fileIndex, delURIs, addURIs, position)

	var reply []uint
	err := c.rpcClient.Call(aria2proto.ChangeURI, args, &reply)

	return reply[0], reply[1], err
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
func (c *Client) ChangeURI(gid string, fileIndex uint, delURIs []string, addURIs []string) (uint, uint, error) {
	args := c.getArgs(gid, fileIndex, delURIs, addURIs)

	var reply []uint
	err := c.rpcClient.Call(aria2proto.ChangeURI, args, &reply)

	return reply[0], reply[1], err
}

// GetOptions returns Options of the download denoted by gid.
// Note that this method does not return options which have no default value and have not been set on the command-line,
// in configuration files or RPC methods.
func (c *Client) GetOptions(gid string) (Options, error) {
	var reply Options
	err := c.rpcClient.Call(aria2proto.GetOptions, c.getArgs(gid), &reply)

	return reply, err
}

// ChangeOptions changes options of the download denoted by gid dynamically.
//
// Except for following options, all options are available:
//   - DryRun
//   - MetalinkBaseURI
//   - ParameterizedURI
//   - Pause
//   - PieceLength
//   - RPCSaveUploadMetadata
//
// Except for the following options, changing the other options of active download makes it restart
// (restart itself is managed by aria2, and no user intervention is required):
//   - BtMaxPeers
//   - BtRequestPeerSpeedLimit
//   - BtRemoveUnselectedFile
//   - ForceSave
//   - MaxDownloadLimit
//   - MaxUploadLimit
func (c *Client) ChangeOptions(gid string, options Options) error {
	return c.rpcClient.Call(aria2proto.ChangeOptions, c.getArgs(gid, options), nil)
}

// GetGlobalOptions returns the global options.
// Note that this method does not return options which have no default value and have not been set on the command-line,
// in configuration files or RPC methods.
//
// Because global options are used as a template for the options of newly added downloads,
// the response contains keys returned by the GetOption() method.
func (c *Client) GetGlobalOptions() (Options, error) {
	var reply Options
	err := c.rpcClient.Call(aria2proto.GetGlobalOptions, c.getArgs(), &reply)

	return reply, err
}

// TODO global options

// ChangeGlobalOptions changes global options dynamically.
//
// The following global options are available:
//   - BtMaxOpenFiles
//   - DownloadResult
//   - KeepUnfinishedDownloadResult
//   - Log
//   - LogLevel
//   - MaxConcurrentDownloads
//   - MaxDownloadResult
//   - MaxOverallDownloadLimit
//   - MaxOverallUploadLimit
//   - OptimizeConcurrentDownloads
//   - SaveCookies
//   - SaveSession
//   - ServerStatOf
//
// Except for the following options, all other Options are available as well:
//   - Checksum
//   - IndexOut
//   - Out
//   - Pause
//   - SelectFile
//
// With the log option, you can dynamically start logging or change log file.
// To stop logging, specify an empty string as the parameter value.
// Note that log file is always opened in append mode.
func (c *Client) ChangeGlobalOptions(options Options) error {
	return c.rpcClient.Call(aria2proto.ChangeGlobalOptions, c.getArgs(options), nil)
}

// GetGlobalStats returns global statistics such as the overall download and upload speeds.
func (c *Client) GetGlobalStats() (Stats, error) {
	var reply Stats
	err := c.rpcClient.Call(aria2proto.GetGlobalStats, c.getArgs(), &reply)

	return reply, err
}

// PurgeDownloadResults purges completed/error/removed downloads to free memory
func (c *Client) PurgeDownloadResults() error {
	return c.rpcClient.Call(aria2proto.PurgeDownloadResults, c.getArgs(), nil)
}

// RemoveDownloadResult removes a completed/error/removed download denoted by gid from memory.
func (c *Client) RemoveDownloadResult(gid string) error {
	return c.rpcClient.Call(aria2proto.RemoveDownloadResult, c.getArgs(gid), nil)
}

// GetVersion returns the version of aria2 and the list of enabled features.
func (c *Client) GetVersion() (VersionInfo, error) {
	var reply VersionInfo
	err := c.rpcClient.Call(aria2proto.GetVersion, c.getArgs(), &reply)

	return reply, err
}

// GetSessionInfo returns session information.
func (c *Client) GetSessionInfo() (SessionInfo, error) {
	var reply SessionInfo
	err := c.rpcClient.Call(aria2proto.GetSessionInfo, c.getArgs(), &reply)

	return reply, err
}

// Shutdown shuts down aria2.
func (c *Client) Shutdown() error {
	return c.rpcClient.Call(aria2proto.Shutdown, c.getArgs(), nil)
}

// ForceShutdown shuts down aria2.
// Behaves like the Shutdown() method but doesn't perform any actions which take time,
// such as contacting BitTorrent trackers to unregister downloads first.
func (c *Client) ForceShutdown() error {
	return c.rpcClient.Call(aria2proto.ForceShutdown, c.getArgs(), nil)
}

// SaveSession saves the current session to a file specified by the SaveSession option.
func (c *Client) SaveSession() error {
	return c.rpcClient.Call(aria2proto.SaveSession, c.getArgs(), nil)
}

// MultiCall executes multiple method calls in one request.
// Returns a MethodResult for each MethodCall in order.
func (c *Client) MultiCall(methods ...*MethodCall) ([]MethodResult, error) {
	var rawResults []json.RawMessage
	err := c.rpcClient.Call(aria2proto.Multicall, c.getArgs(methods), &rawResults)

	results := make([]MethodResult, len(rawResults))

	for i, rawResult := range rawResults {
		var methodResult []byte

		var methodErr MethodCallError
		_ = json.Unmarshal(rawResult, &methodErr)

		if methodErr == (MethodCallError{}) {
			var resultArray [1]json.RawMessage
			_ = json.Unmarshal(rawResult, &resultArray)

			methodResult = resultArray[0]
		}

		results[i] = MethodResult{Result: methodResult, Error: &methodErr}
	}

	return results, err
}
