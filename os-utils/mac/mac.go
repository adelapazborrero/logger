package mac

import (
	"bytes"
)

type MacLogger struct {
}

func NewMacLogger() *MacLogger {
	return &MacLogger{}
}

func (l *MacLogger) CaptureKeyboardKeys(logBuffer *bytes.Buffer) {
	// Not implemented
}
