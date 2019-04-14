package arigo

// Integer returned by aria2 for failed downloads.
// Please see https://aria2.github.io/manual/en/html/aria2c.html#exit-status
type ExitStatus uint8

const (
	// All downloads were successful.
	Success ExitStatus = iota
	// An unknown error occurred.
	UnknownError
	// Time out occurred.
	Timeout
	// A resource was not found.
	ResourceNotFound
	// aria2 saw the specified number of “resource not found” error. See --max-file-not-found option.
	ResourceNotFoundReached
	// A download aborted because download speed was too slow. See --lowest-speed-limit option.
	DownloadSpeedTooSlow
	// Network problem occurred.
	NetworkError
	// There were unfinished downloads.
	// This error is only reported if all finished downloads were successful and there were unfinished
	// downloads in a queue when aria2 exited by pressing Ctrl-C by an user or sending TERM or INT signal.
	UnfinishedDownloads
	// Remote server did not support resume when resume was required to complete download.
	RemoteNoResume
	// There was not enough disk space available.
	NotEnoughDiskSpace
	// Piece length was different from one in .aria2 control file. See --allow-piece-length-change option.
	PieceLengthMismatch
	// aria2 was downloading same file at that moment.
	SameFileBeingDownloaded
	// aria2 was downloading same info hash torrent at that moment.
	SameInfoHashBeingDownloaded
	// File already existed. See --allow-overwrite option.
	FileAlreadyExists
	// Renaming file failed. See --auto-file-renaming option.
	RenamingFailed
	// aria2 could not open existing file.
	CouldNotOpenExistingFile
	// aria2 could not create new file or truncate existing file.
	CouldNotCreateNewFile
	// File I/O error occurred.
	FileIOError
	// aria2 could not create directory.
	CouldNotCreateDirectory
	// Name resolution failed.
	NameResolutionFailed
	// aria2 could not parse Metalink document.
	MetalinkParsingFailed
	// FTP command failed.
	FTPCommandFailed
	// HTTP response header was bad or unexpected.
	HTTPResponseHeaderBad
	// Too many redirects occurred.
	TooManyRedirects
	// HTTP authorization failed.
	HttpAuthorizationFailed
	// aria2 could not parse bencoded file (usually “.torrent” file).
	BencodedFileParseError
	// “.torrent” file was corrupted or missing information that aria2 needed.
	TorrentFileCorrupt
	// Magnet URI was bad.
	MagnetURIBad
	// The remote server was unable to handle the request due to a temporary overloading or maintenance.
	RemoteServerHandleRequestError
	// aria2 could not parse JSON-RPC request.
	JSONRPCParseError
	// Reserved. Not used. If you get this error then the library is probably out-of-date,
	// or the universe is breaking down.
	Reserved
	// Checksum validation failed.
	ChecksumValidationFailed
)
