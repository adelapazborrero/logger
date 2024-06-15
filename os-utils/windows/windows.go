package windows

import (
	"bytes"
)

type WindowsLogger struct {
}

func NewWindowsLogger() *WindowsLogger {
	return &WindowsLogger{}
}

func (l *WindowsLogger) CaptureKeyboardKeys(logBuffer *bytes.Buffer) {
	// Not implemented
}
