// utils.go
package klog

import (
	"os"
	"reflect"
	"runtime"
	"strings"
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

const ansiReset = "\x1b[0m"

type ansi struct {
	Enable bool
}

func (c *ansi) getMethod(name string) (reflect.Method, bool) {
	methodName := strings.Title(name)
	return reflect.TypeOf(c).MethodByName(methodName)
}

func (c *ansi) print(color, s string) string {
	if c.Enable {
		return color + s + ansiReset
	}
	return s
}

func (c *ansi) Red(s string) string {
	return c.print("\x1b[31m", s)
}

func (c *ansi) Green(s string) string {
	return c.print("\x1b[32m", s)
}

func (c *ansi) Yellow(s string) string {
	return "\x1b[33m" + s + ansiReset
}

func (c *ansi) Blue(s string) string {
	return "\x1b[34m" + s + ansiReset
}

func (c *ansi) Magenta(s string) string {
	return "\x1b[35m" + s + ansiReset
}

func (c *ansi) Cyan(s string) string {
	return "\x1b[36m" + s + ansiReset
}
