/*
Color support for linux shell
*/
package klog

import (
	"reflect"
	"strings"
)

const ansiReset = "\x1b[0m"

type ansi struct{}

func (c *ansi) getMethod(name string) (reflect.Method, bool) {
	methodName := strings.Title(name)
	return reflect.TypeOf(c).MethodByName(methodName)
}

func (c *ansi) print(color, s string) string {
	return color + s + ansiReset
}

func (c *ansi) Red(s string) string {
	return c.print("\x1b[31m", s)
}

func (c *ansi) Green(s string) string {
	return c.print("\x1b[32m", s)
}

func (c *ansi) Yellow(s string) string {
	return c.print("\x1b[33m", s)
}

func (c *ansi) Blue(s string) string {
	return c.print("\x1b[34m", s)
}

func (c *ansi) Magenta(s string) string {
	return c.print("\x1b[35m", s)
}

func (c *ansi) Cyan(s string) string {
	return c.print("\x1b[36m", s)
}
