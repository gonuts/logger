package logger

import (
	"fmt"
	"io"
	"os"
)

type Level int

const (
	VERBOSE Level = -2
	DEBUG         = -1
	INFO          = 0
	WARNING       = 1
	ERROR         = 2
	FATAL         = 3
)

type Logger struct {
	name string
	lvl  Level
	w    io.Writer
}

func New(name string) *Logger {
	return &Logger{
		name: name,
		lvl:  INFO,
		w:    os.Stdout,
	}
}

func NewLogger(name string, lvl Level, w io.Writer) *Logger {
	return &Logger{
		name: name,
		lvl:  lvl,
		w:    w,
	}
}

func (msg *Logger) Name() string {
	return msg.name
}

func (msg *Logger) Level() Level {
	return msg.lvl
}

func (msg *Logger) Errorf(format string, args ...interface{}) (int, error) {
	if msg.lvl <= ERROR {
		return fmt.Fprintf(msg.w, msg.name+" ERROR    "+format, args...)
	}
	return 0, nil
}

func (msg *Logger) Warnf(format string, args ...interface{}) (int, error) {
	if msg.lvl <= WARNING {
		return fmt.Fprintf(msg.w, msg.name+" WARNING "+format, args...)
	}
	return 0, nil
}

func (msg *Logger) Infof(format string, args ...interface{}) (int, error) {
	if msg.lvl <= INFO {
		return fmt.Fprintf(msg.w, msg.name+" INFO    "+format, args...)
	}
	return 0, nil
}

func (msg *Logger) Debugf(format string, args ...interface{}) (int, error) {
	if msg.lvl <= DEBUG {
		return fmt.Fprintf(msg.w, msg.name+" DEBUG   "+format, args...)
	}
	return 0, nil
}

func (msg *Logger) Verbosef(format string, args ...interface{}) (int, error) {
	if msg.lvl <= VERBOSE {
		return fmt.Fprintf(msg.w, msg.name+" VERBOSE "+format, args...)
	}
	return 0, nil
}

// EOF
