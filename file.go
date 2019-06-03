package arigo

// File represents a single file downloaded by aria2.
// It is returned by the GetFiles() method.
type File struct {
	// Index of the file, starting at 1, in the same order as files appear in the multi-file torrent.
	Index  int    `json:"index,string"`
	Path   string `json:"path"`          // File path
	Length uint   `json:"length,string"` // File size in bytes

	// Completed length of this file in bytes.
	// Please note that it is possible that sum of completedLength is less than the completedLength returned
	// by the TellStatus() method. This is because completedLength in GetFiles() only includes completed pieces.
	// On the other hand, completedLength in TellStatus() also includes partially completed pieces.
	CompletedLength uint `json:"completedLength,string"`

	// true if this file is selected by the SelectFile option.
	// If SelectFile is not specified or this is single-file torrent or not a torrent download at all,
	// this value is always true. Otherwise false.
	Selected bool  `json:"selected,string"`
	URIs     []URI `json:"uris"` // Array of URIs for this file.
}
