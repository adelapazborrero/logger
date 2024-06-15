//go:build darwin
// +build darwin

package main

import (
	"github.com/adelapazborrero/logger/keys"
)

func init() {
	keyLogger.CurrentLogger = keys.NewMacLogger()
}
