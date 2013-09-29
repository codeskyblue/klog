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
	"[INFO_]",
	"[WARN_]",
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
	out     io.Writer
	level   Level
	logging *log.Logger
	color   *ansi
	flags   int
}

// default level is debug
func NewLogger(out io.Writer, prefix string) *Logger {
	if out == nil {
		out = os.Stdout
	}
	colorEnable := isTermOutput() //&& runtime.GOOS != "windows"
	return &Logger{
		level:   LInfo,
		logging: log.New(out, prefix, log.Ldate|log.Ltime),
		color:   &ansi{colorEnable},
		flags:   Fstdflag,
	}
}

func (l *Logger) SetFlags(flag int) {
	l.flags = flag
}

func (l *Logger) SetLevel(level Level) {
	l.level = level
}

// just an empty function.
// FIXME:defer k.Flush() // When call Fatal(..), program will call Flush too(so never mind if you forgot).
//func (l *Logger) Flush(){
//}

func (l *Logger) write(level Level, format string, a ...interface{}) {
	if level < l.level {
		return
	}
	var levelName string = levels[int(level)]
	var preStr string

	if l.flags&Fcolor != 0 {
		var colorName string = colors[int(level)]
		method, exists := l.color.getMethod(colorName)
		if exists {
			levelName = method.Func.Call([]reflect.Value{
				reflect.ValueOf(l.color),
				reflect.ValueOf(levelName)},
			)[0].String()
		}
	}
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

	sep := " "
	if format == "" {
		l.logging.Println(preStr + sep + fmt.Sprint(a...))
	} else {
		l.logging.Println(preStr + sep + fmt.Sprintf(format, a...))
	}
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
	l.Fatalf("", v...)
}

// will also call os.Exit(1)
func (l *Logger) Fatalf(format string, v ...interface{}) {
	l.write(LFatal, format, v...)
	os.Exit(1)
}
