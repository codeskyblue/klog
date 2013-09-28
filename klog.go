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
	LDebug Level = iota
	LInfo
	LWarning
	LError
	LFatal
)

var levels = []string{
	"DEBUG",
	"INFO",
	"WARN",
	"ERROR",
	"FATAL",
}

var colors = []string{
	"Green",
	"Red",
	"Red",
	"Red",
	"Red",
}

type Logger struct {
	out   io.Writer
	level Level
	slog  *log.Logger
	color *ansi
}

func NewLogger(out io.Writer, prefix string) *Logger {
	if out == nil {
		out = os.Stdout
	}
	return &Logger{
		level: LInfo,
		slog:  log.New(out, prefix, log.Ldate|log.Ltime),
		color: &ansi{isTermOutput() && runtime.GOOS != "windows"},
	}
}

func (l *Logger) SetLevel(level Level) {
	l.level = level
}

func (l *Logger) write(level Level, format string, a ...interface{}) {
	if level < l.level {
		return
	}
	var s string = levels[int(level)]
	var colorName string = colors[int(level)]
	// Retrieve the stack infos
	_, file, line, ok := runtime.Caller(2)
	if !ok {
		file = "<unknown>"
		line = -1
	} else {
		file = file[strings.LastIndex(file, "/")+1:]
	}

	method, exists := l.color.getMethod(colorName)
	var levelStr string
	if !exists {
		levelStr = s
	} else {
		levelStr = method.Func.Call([]reflect.Value{
			reflect.ValueOf(l.color),
			reflect.ValueOf(s)},
		)[0].String()
	}

	l.slog.Printf(fmt.Sprintf("[%s] %s:%d  %s\n", levelStr, file, line, format), a...)
	if level == LFatal {
		os.Exit(1)
	}

}

func (l *Logger) Debugf(format string, v ...interface{}) {
	l.write(LDebug, format, v...)
}

/*
func Debug(v ...interface{}) {
	logPrint(LDebug, "%v", v)
}

func Info(v ...interface{}) {
	logPrint(LInfo, "%v", v)
}

func Warn(v ...interface{}) {
	logPrint(LWarning, "%v", v)
}

func Error(v ...interface{}) {
	logPrint(LError, "%v", v)
}
*/
