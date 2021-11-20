package logger

import (
	"context"
	"fmt"
	"io"
	"os"
	"runtime"
	"strings"
	"sync/atomic"
	"time"
)

// Colors

type color string

const (
	reset color = "\033[0m"
	// red         color = "\033[31m"
	green color = "\033[32m"
	// yellow      color = "\033[33m"
	// blue        color = "\033[34m"
	// magenta     color = "\033[35m"
	cyan  color = "\033[36m"
	white color = "\033[37m"
	// blueBold    color = "\033[34;1m"
	magentaBold color = "\033[35;1m"
	redBold     color = "\033[31;1m"
	// yellowBold  color = "\033[33;1m"
)

type LogLevel uint32

const (
	InfoLevel LogLevel = iota
	DebugLevel
	WarnLevel
	ErrorLevel
	CriticalLevel
)

func (l LogLevel) String() string {
	switch l {
	case InfoLevel:
		return "info"
	case DebugLevel:
		return "debug"
	case WarnLevel:
		return "warn"
	case ErrorLevel:
		return "error"
	case CriticalLevel:
		return "critical"
	}
	return "unknown"
}

const (
	infoColor     = white
	debugColor    = green
	warnColor     = cyan
	errorColor    = redBold
	criticalColor = magentaBold
)

type Interface interface {
	LogMode(LogLevel) Interface
	Info(context.Context, string, ...interface{})
	Debug(context.Context, string, ...interface{})
	Warn(context.Context, string, ...interface{})
	Error(context.Context, string, ...interface{})
	Critical(context.Context, string, ...interface{})
}

type Logger struct {
	Level         LogLevel
	output        io.Writer
	hideCallstack bool
	depth         int
}

func (logger *Logger) level() LogLevel {
	return LogLevel(atomic.LoadUint32((*uint32)(&logger.Level)))
}

func (logger *Logger) SetLevel(level LogLevel) {
	atomic.StoreUint32((*uint32)(&logger.Level), uint32(level))
}

func (level LogLevel) Color() color {
	switch level {
	case InfoLevel:
		return infoColor
	case DebugLevel:
		return debugColor
	case WarnLevel:
		return warnColor
	case ErrorLevel:
		return errorColor
	case CriticalLevel:
		return criticalColor
	}

	return white
}

func (logger *Logger) formatOutput(ctx context.Context, level LogLevel, output string) string {
	now := time.Now().Format("2006-01-02 15:04:05")

	_, file, line, ok := runtime.Caller(logger.depth)
	if !ok {
		file = "???"
		line = 0
	}
	// short file name
	for i := len(file) - 1; i > 0; i-- {
		if file[i] == '/' {
			file = file[i+1:]
			break
		}
	}
	return fmt.Sprintf("%-25s -%s- %s (%s:%d)",
		now, strings.ToUpper(level.String()), output, file, line)
}

func NewLogger() *Logger {
	return &Logger{
		Level:  CriticalLevel,
		output: os.Stdout,
		depth:  3,
	}
}

func (logger *Logger) logf(ctx context.Context, level LogLevel, format string, args ...interface{}) {
	if logger.Level < level {
		return
	}

	fmt.Fprintf(logger.output, "%s %s\n %s", level.Color(), logger.formatOutput(ctx, level, fmt.Sprintf(format, args...)), reset)
}

func (logger *Logger) Info(ctx context.Context, format string, args ...interface{}) {
	logger.logf(ctx, InfoLevel, format, args...)
}

func (logger *Logger) Warn(ctx context.Context, format string, args ...interface{}) {
	logger.logf(ctx, WarnLevel, format, args...)
}

func (logger *Logger) Error(ctx context.Context, format string, args ...interface{}) {
	logger.logf(ctx, ErrorLevel, format, args...)
}

func (logger *Logger) Debug(ctx context.Context, format string, args ...interface{}) {
	logger.logf(ctx, DebugLevel, format, args...)
}

func (logger *Logger) Critical(ctx context.Context, format string, args ...interface{}) {
	logger.logf(ctx, CriticalLevel, format, args...)
}

var logger = NewLogger()

func Info(ctx context.Context, format string, args ...interface{}) {
	logger.Info(ctx, format, args...)
}

func Debug(ctx context.Context, format string, args ...interface{}) {
	logger.Debug(ctx, format, args...)
}

func Warn(ctx context.Context, format string, args ...interface{}) {
	logger.Warn(ctx, format, args...)
}

func Error(ctx context.Context, format string, args ...interface{}) {
	logger.Error(ctx, format, args...)
}

func Critical(ctx context.Context, format string, args ...interface{}) {
	logger.Critical(ctx, format, args...)
}

func (logger *Logger) HideCallstack() *Logger {
	logger.hideCallstack = true
	return logger
}
