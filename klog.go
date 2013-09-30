// beelog project klog
package klog

import (
	"fmt"
	"io"
	"log"
	"os"
	"reflect"
	"runtime"
	"strings"
)

type Level int

const (
	Fshortfile = 1 << iota // show filename:lineno
	Fcolor
	Fdevflag = Fshortfile | Fcolor // for develop use
	Fstdflag = Fcolor
)

const (
	LDebug Level = iota
	LInfo
	LWarning
	LError
	LFatal
)

var levels = []string{
	"[DEBUG]",
	"[INFO.]",
	"[WARN.]",
	"[ERROR]",
	"[FATAL]",
}

var colors = []string{
	"Cyan",
	"Green",
	"Magenta",
	"Yellow",
	"Red",
}

type Logger struct {
	out         io.Writer
	level       Level
	logging     *log.Logger
	color       *ansi
	flags       int
	colorEnable bool
}

// default level is debug
func NewLogger(out io.Writer, prefix string) *Logger {
	if out == nil {
		out = os.Stdout
	}
	return &Logger{
		level:       LInfo,
		logging:     log.New(out, prefix, log.Ldate|log.Ltime),
		color:       &ansi{},
		colorEnable: isTermOutput() && runtime.GOOS != "windows",
		flags:       Fstdflag,
	}
}

func (l *Logger) SetFlags(flag int) {
	l.flags = flag
}

func (l *Logger) SetLevel(level Level) {
	l.level = level
}

func (l *Logger) write(level Level, format string, a ...interface{}) {
	if level < l.level {
		return
	}
	var levelName string = levels[int(level)]
	var preStr string

	preStr += levelName
	if l.flags&Fshortfile != 0 {
		// Retrieve the stack infos
		_, file, line, ok := runtime.Caller(2)
		if !ok {
			file = "<unknown>"
			line = -1
		} else {
			file = file[strings.LastIndex(file, "/")+1:]
		}
		preStr = fmt.Sprintf("%s %s:%d", preStr, file, line)
	}

	var outstr, sep = "", " "
	if format == "" {
		outstr = preStr + sep + fmt.Sprint(a...)
	} else {
		outstr = preStr + sep + fmt.Sprintf(format, a...)
	}
	outstr = strings.TrimSuffix(outstr, "\n")

	if l.colorEnable && l.flags&Fcolor != 0 {
		var colorName string = colors[int(level)]
		method, exists := l.color.getMethod(colorName)
		if exists {
			outstr = method.Func.Call([]reflect.Value{
				reflect.ValueOf(l.color),
				reflect.ValueOf(outstr)},
			)[0].String()
		}
	}
	l.logging.Print(outstr)
}

func (l *Logger) Debug(v ...interface{}) {
	l.write(LDebug, "", v...)
}

func (l *Logger) Debugf(format string, v ...interface{}) {
	l.write(LDebug, format, v...)
}

func (l *Logger) Info(v ...interface{}) {
	l.write(LInfo, "", v...)
}
func (l *Logger) Infof(format string, v ...interface{}) {
	l.write(LInfo, format, v...)
}
func (l *Logger) Warn(v ...interface{}) {
	l.write(LWarning, "", v...)
}
func (l *Logger) Warnf(format string, v ...interface{}) {
	l.write(LWarning, format, v...)
}
func (l *Logger) Error(v ...interface{}) {
	l.write(LError, "", v...)
}
func (l *Logger) Errorf(format string, v ...interface{}) {
	l.write(LError, format, v...)
}

// will also call os.Exit(1)
func (l *Logger) Fatal(v ...interface{}) {
	l.write(LFatal, "", v...)
	os.Exit(1)
}

// will also call os.Exit(1)
func (l *Logger) Fatalf(format string, v ...interface{}) {
	l.write(LFatal, format, v...)
	os.Exit(1)
}
