package aria2proto

const (
	AddURI      = "aria2.addUri"
	AddTorrent  = "aria2.addTorrent"
	AddMetalink = "aria2.addMetalink"

	Remove      = "aria2.remove"
	ForceRemove = "aria2.forceRemove"

	Pause      = "aria2.pause"
	ForcePause = "aria2.forcePause"
	Unpause    = "aria2.unpause"

	TellStatus = "aria2.tellStatus"

	GetURIs    = "aria2.getUris"
	GetFiles   = "aria2.getFiles"
	GetPeers   = "aria2.getPeers"
	GetServers = "aria2.getServers"

	ChangePosition = "aria2.changePosition"
	ChangeURI      = "aria2.changeUri"

	GetOptions    = "aria2.getOption"
	ChangeOptions = "aria2.changeOption"

	RemoveDownloadResult = "aria2.removeDownloadResult"
)
