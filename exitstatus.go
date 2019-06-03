package arigo

//go:generate stringer -type=ExitStatus

// ExitStatus is an integer returned by aria2 for downloads which describes why a download exited.
// Please see https://aria2.github.io/manual/en/html/aria2c.html#exit-status
type ExitStatus uint8

const (
	// Success indicates that all downloads were successful.
	Success ExitStatus = iota

	// UnknownError indicates that an unknown error occurred.
	UnknownError

	// Timeout indicates that a timeout occurred.
	Timeout

	// ResourceNotFound indicates that a resource was not found.
	ResourceNotFound

	// ResourceNotFoundReached indicates that aria2 saw the specified number of “resource not found” error.
	// See the --max-file-not-found option.
	ResourceNotFoundReached

	// DownloadSpeedTooSlow indicates that a download aborted because download speed was too slow.
	// See --lowest-speed-limit option.
	DownloadSpeedTooSlow

	// NetworkError indicates that a network problem occurred.
	NetworkError

	// UnfinishedDownloads indicates that there were unfinished downloads.
	// This error is only reported if all finished downloads were successful and there were unfinished
	// downloads in a queue when aria2 exited by pressing Ctrl-C by an user or sending TERM or INT signal.
	UnfinishedDownloads

	// RemoteNoResume indicates that the remote server did not support resume when resume was required to complete download.
	RemoteNoResume

	// NotEnoughDiskSpace indicates that there was not enough disk space available.
	NotEnoughDiskSpace

	// PieceLengthMismatch indicates that the piece length was different from one in .aria2 control file.
	// See --allow-piece-length-change option.
	PieceLengthMismatch

	// SameFileBeingDownloaded indicates that aria2 was downloading same file at that moment.
	SameFileBeingDownloaded

	// SameInfoHashBeingDownloaded indicates that aria2 was downloading same info hash torrent at that moment.
	SameInfoHashBeingDownloaded

	// FileAlreadyExists indicates that the file already existed. See --allow-overwrite option.
	FileAlreadyExists

	//RenamingFailed indicates that renaming the file failed. See --auto-file-renaming option.
	RenamingFailed

	// CouldNotOpenExistingFile indicates that aria2 could not open existing file.
	CouldNotOpenExistingFile

	// CouldNotCreateNewFile indicates that aria2 could not create new file or truncate existing file.
	CouldNotCreateNewFile

	// FileIOError indicates that a file I/O error occurred.
	FileIOError

	// CouldNotCreateDirectory indicates that aria2 could not create directory.
	CouldNotCreateDirectory

	// NameResolutionFailed indicates that the name resolution failed.
	NameResolutionFailed

	// MetalinkParsingFailed indicates that aria2 could not parse Metalink document.
	MetalinkParsingFailed

	// FTPCommandFailed indicates that the FTP command failed.
	FTPCommandFailed

	// HTTPResponseHeaderBad indicates that the HTTP response header was bad or unexpected.
	HTTPResponseHeaderBad

	// TooManyRedirects indicates that too many redirects occurred.
	TooManyRedirects

	// HTTPAuthorizationFailed indicates that HTTP authorization failed.
	HTTPAuthorizationFailed

	// BencodedFileParseError indicates that aria2 could not parse bencoded file (usually “.torrent” file).
	BencodedFileParseError

	// TorrentFileCorrupt indicates that the “.torrent” file was corrupted or missing information that aria2 needed.
	TorrentFileCorrupt

	// MagnetURIBad indicates that the magnet URI was bad.
	MagnetURIBad

	// RemoteServerHandleRequestError indicates that the remote server was unable to handle the request due to a
	// temporary overloading or maintenance.
	RemoteServerHandleRequestError

	// JSONRPCParseError indicates that aria2 could not parse JSON-RPC request.
	JSONRPCParseError

	// Reserved is a reserved value. If you get this exit status then the library is probably out-of-date,
	// or the universe is breaking down.
	Reserved

	// ChecksumValidationFailed indicates that the checksum validation failed.
	ChecksumValidationFailed
)
