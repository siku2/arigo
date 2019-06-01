package aria2proto

const (
	PauseAll      = "aria2.pauseAll"
	ForcePauseAll = "aria2.forcePauseAll"
	UnpauseAll    = "aria2.unpauseAll"

	TellActive  = "aria2.tellActive"
	TellWaiting = "aria2.tellWaiting"
	TellStopped = "aria2.tellStopped"

	GetGlobalOptions    = "aria2.getGlobalOption"
	ChangeGlobalOptions = "aria2.changeGlobalOption"
	GetGlobalStats      = "aria2.getGlobalStat"

	PurgeDownloadResults = "aria2.purgeDownloadResult"

	GetVersion     = "aria2.getVersion"
	GetSessionInfo = "aria2.getSessionInfo"

	Shutdown      = "aria2.shutdown"
	ForceShutdown = "aria2.forceShutdown"

	SaveSession = "aria2.saveSession"
)
