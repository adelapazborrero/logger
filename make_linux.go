//go:build linux
// +build linux

package main

import (
	"github.com/adelapazborrero/logger/keys"
)

func init() {
	keyLogger.CurrentLogger = keys.NewLinuxLogger()
}
