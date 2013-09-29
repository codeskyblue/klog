// utils.go
package klog

import (
	"os"
	"runtime"
)

func isTermOutput() (result bool) {
	switch runtime.GOOS {
	case "linux":
		fi, _ := os.Stdout.Stat()
		return fi.Mode()&os.ModeCharDevice == os.ModeCharDevice
	case "windows":
		return os.Stdout.Fd() < 0xff
	}
	return false
}
