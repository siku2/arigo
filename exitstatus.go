package arigo

import errs "github.com/siku2/arigo/errors"

// Re-export the ExitStatus type so callers can still use arigo.ExitStatus
type ExitStatus = errs.ExitStatus

const (
	// Success indicates that all downloads were successful.
	Success = errs.Success

	// UnknownError indicates that an unknown error occurred.
	UnknownError = errs.UnknownError

	// Timeout indicates that a timeout occurred.
	Timeout = errs.Timeout

	// ResourceNotFound indicates that a resource was not found.
	ResourceNotFound = errs.ResourceNotFound

	// ResourceNotFoundReached indicates that aria2 saw the specified number of “resource not found” error.
	// See the --max-file-not-found option.
	ResourceNotFoundReached = errs.ResourceNotFoundReached

	// DownloadSpeedTooSlow indicates that a download aborted because download speed was too slow.
	// See --lowest-speed-limit option.
	DownloadSpeedTooSlow = errs.DownloadSpeedTooSlow

	// NetworkError indicates that a network problem occurred.
	NetworkError = errs.NetworkError

	// UnfinishedDownloads indicates that there were unfinished downloads.
	// This error is only reported if all finished downloads were successful and there were unfinished
	// downloads in a queue when aria2 exited by pressing Ctrl-C by an user or sending TERM or INT signal.
	UnfinishedDownloads = errs.UnfinishedDownloads

	// RemoteNoResume indicates that the remote server did not support resume when resume was required to complete download.
	RemoteNoResume = errs.RemoteNoResume

	// NotEnoughDiskSpace indicates that there was not enough disk space available.
	NotEnoughDiskSpace = errs.NotEnoughDiskSpace

	// PieceLengthMismatch indicates that the piece length was different from one in .aria2 control file.
	// See --allow-piece-length-change option.
	PieceLengthMismatch = errs.PieceLengthMismatch

	// SameFileBeingDownloaded indicates that aria2 was downloading same file at that moment.
	SameFileBeingDownloaded = errs.SameFileBeingDownloaded

	// SameInfoHashBeingDownloaded indicates that aria2 was downloading same info hash torrent at that moment.
	SameInfoHashBeingDownloaded = errs.SameInfoHashBeingDownloaded

	// FileAlreadyExists indicates that the file already existed. See --allow-overwrite option.
	FileAlreadyExists = errs.FileAlreadyExists

	// RenamingFailed indicates that renaming the file failed. See --auto-file-renaming option.
	RenamingFailed = errs.RenamingFailed

	// CouldNotOpenExistingFile indicates that aria2 could not open existing file.
	CouldNotOpenExistingFile = errs.CouldNotOpenExistingFile

	// CouldNotCreateNewFile indicates that aria2 could not create new file or truncate existing file.
	CouldNotCreateNewFile = errs.CouldNotCreateNewFile

	// FileIOError indicates that a file I/O error occurred.
	FileIOError = errs.FileIOError

	// CouldNotCreateDirectory indicates that aria2 could not create directory.
	CouldNotCreateDirectory = errs.CouldNotCreateDirectory

	// NameResolutionFailed indicates that the name resolution failed.
	NameResolutionFailed = errs.NameResolutionFailed

	// MetalinkParsingFailed indicates that aria2 could not parse Metalink document.
	MetalinkParsingFailed = errs.MetalinkParsingFailed

	// FTPCommandFailed indicates that the FTP command failed.
	FTPCommandFailed = errs.FTPCommandFailed

	// HTTPResponseHeaderBad indicates that the HTTP response header was bad or unexpected.
	HTTPResponseHeaderBad = errs.HTTPResponseHeaderBad

	// TooManyRedirects indicates that too many redirects occurred.
	TooManyRedirects = errs.TooManyRedirects

	// HTTPAuthorizationFailed indicates that HTTP authorization failed.
	HTTPAuthorizationFailed = errs.HTTPAuthorizationFailed

	// BencodedFileParseError indicates that aria2 could not parse bencoded file (usually “.torrent” file).
	BencodedFileParseError = errs.BencodedFileParseError

	// TorrentFileCorrupt indicates that the “.torrent” file was corrupted or missing information that aria2 needed.
	TorrentFileCorrupt = errs.TorrentFileCorrupt

	// MagnetURIBad indicates that the magnet URI was bad.
	MagnetURIBad = errs.MagnetURIBad

	// RemoteServerHandleRequestError indicates that the remote server was unable to handle the request due to a
	// temporary overloading or maintenance.
	RemoteServerHandleRequestError = errs.RemoteServerHandleRequestError

	// JSONRPCParseError indicates that aria2 could not parse JSON-RPC request.
	JSONRPCParseError = errs.JSONRPCParseError

	// Reserved is a reserved value. If you get this exit status then the library is probably out-of-date,
	// or the universe is breaking down.
	Reserved = errs.Reserved

	// ChecksumValidationFailed indicates that the checksum validation failed.
	ChecksumValidationFailed = errs.ChecksumValidationFailed
)
