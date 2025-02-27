package tars

import (
	"time"

	"github.com/jslyzt/tarsgo/tars/protocol/res/logf"
)

// RemoteTimeWriter writer for writing remote log.
type RemoteTimeWriter struct {
	logInfo       *logf.LogInfo
	logs          chan string
	logPtr        *logf.Log
	reportSuccPtr *PropertyReport
	reportFailPtr *PropertyReport
	hasPrefix     bool
}

// NewRemoteTimeWriter new and init RemoteTimeWriter
func NewRemoteTimeWriter() *RemoteTimeWriter {
	rw := new(RemoteTimeWriter)
	rw.logInfo = new(logf.LogInfo)
	rw.logInfo.SFormat = "%Y%m%d"
	rw.logInfo.SConcatStr = "_"

	logs := make(chan string, remoteLogQueueSize)
	rw.logs = logs
	rw.logPtr = new(logf.Log)
	comm := NewCommunicator()
	node := GetServerConfig().Log
	rw.EnableSufix(true)
	rw.EnablePrefix(true)
	rw.SetSeparator("|")
	rw.SetPrefix(true)

	comm.StringToProxy(node, rw.logPtr)
	go rw.Sync2remote()
	return rw
}

// Sync2remote syncs the log buffer to remote.
func (rw *RemoteTimeWriter) Sync2remote() {
	maxLen := remoteLogMaxNumOneTime
	interval := time.Second
	v := make([]string, 0, maxLen)
	for {
		select {
		case log := <-rw.logs:
			v = append(v, log)
			if len(v) >= maxLen {
				err := rw.sync2remote(v)
				if err != nil {
					TLOG.Error("sync to remote error")
					rw.reportFailPtr.Report(len(v))
				}
				rw.reportSuccPtr.Report(len(v))
				v = make([]string, 0, maxLen) //reset the slice after syncing log to remote
			}
		case <-time.After(interval):
			if len(v) > 0 {
				err := rw.sync2remote(v)
				if err != nil {
					TLOG.Error("sync to remote error")
					rw.reportFailPtr.Report(len(v))
				}
				rw.reportSuccPtr.Report(len(v))
				v = make([]string, 0, maxLen) //reset the slice after syncing log to remote
			}
		}
	}
}

func (rw *RemoteTimeWriter) sync2remote(s []string) error {
	err := rw.logPtr.LoggerbyInfo(rw.logInfo, s)
	return err
}

// InitServerInfo init the remote log server info.
func (rw *RemoteTimeWriter) InitServerInfo(app string, server string, filename string, setdivision string) {
	rw.logInfo.Appname = app
	rw.logInfo.Servername = server
	rw.logInfo.SFilename = filename
	rw.logInfo.Setdivision = setdivision

	serverInfo := app + "." + server + "." + filename
	failServerInfo := serverInfo + "_log_send_fail"
	failSum := NewSum()
	rw.reportFailPtr = CreatePropertyReport(failServerInfo, failSum)
	succServerInfo := serverInfo + "_log_send_succ"
	succSum := NewSum()
	rw.reportSuccPtr = CreatePropertyReport(succServerInfo, succSum)

}

// EnableSufix puts sufix after logs.
func (rw *RemoteTimeWriter) EnableSufix(hasSufix bool) {
	rw.logInfo.BHasSufix = hasSufix
}

// EnablePrefix puts prefix before logs.
func (rw *RemoteTimeWriter) EnablePrefix(hasAppNamePrefix bool) {
	rw.logInfo.BHasAppNamePrefix = hasAppNamePrefix
}

// SetFileNameConcatStr sets the filename concat string.
func (rw *RemoteTimeWriter) SetFileNameConcatStr(s string) {
	rw.logInfo.SConcatStr = s

}

// SetSeparator set seprator between logs.
func (rw *RemoteTimeWriter) SetSeparator(s string) {
	rw.logInfo.SSepar = s
}

// EnableSqarewrapper enables SquareBracket wrapper for the logs.
func (rw *RemoteTimeWriter) EnableSqarewrapper(hasSquareBracket bool) {
	rw.logInfo.BHasSquareBracket = hasSquareBracket
}

// SetLogType sets the log type.
func (rw *RemoteTimeWriter) SetLogType(logType string) {
	rw.logInfo.SLogType = logType

}

// InitFormat sets the log format.
func (rw *RemoteTimeWriter) InitFormat(s string) {
	rw.logInfo.SFormat = s
}

// NeedPrefix return if need prefix for the logger.
func (rw *RemoteTimeWriter) NeedPrefix() bool {
	return rw.hasPrefix
}

// SetPrefix set if need prefix for the logger.
func (rw *RemoteTimeWriter) SetPrefix(enable bool) {
	rw.hasPrefix = enable
}

// Write Writes the logs to the buffer.
func (rw *RemoteTimeWriter) Write(b []byte) {
	s := string(b[:])
	select {
	case rw.logs <- s:
	default:
		TLOG.Error("remote log chan is full")

	}
}
