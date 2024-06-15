package utils

import (
	"os"
	"runtime"
	"strings"
	"time"
)

func AVBehaviourCheck() {
	now := time.Now()
	time.Sleep(3 * time.Second)

	if time.Since(now) < 3*time.Second {
		os.Exit(1)
	}
}

func SandboxCheck() {
	sandboxIndicators := []string{"virtualbox", "vmware", "hyper-v"}
	for _, indicator := range sandboxIndicators {
		if strings.Contains(runtime.GOARCH, indicator) || strings.Contains(runtime.GOOS, indicator) {
			os.Exit(1)
		}
	}
}
