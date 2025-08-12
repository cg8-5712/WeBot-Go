package utils

import (
	"fmt"
	"github.com/fatih/color"
	"log"
	"os"
	"strings"
	"sync"
	"time"
)

type LogLevel int

const (
	LevelDebug LogLevel = iota
	LevelInfo
	LevelWarn
	LevelError
)

var (
	mu           sync.Mutex
	currentLevel = LevelDebug
	logger       = log.New(os.Stdout, "", 0)
)

func SetLogLevel(level LogLevel) {
	mu.Lock()
	defer mu.Unlock()
	currentLevel = level
}

func ParseLogLevel(levelStr string) LogLevel {
	switch strings.ToLower(levelStr) {
	case "debug":
		return LevelDebug
	case "info":
		return LevelInfo
	case "warn", "warning":
		return LevelWarn
	case "error":
		return LevelError
	default:
		return LevelDebug // 默认debug
	}
}

func InitLogLevelFromConfig() {
	if Cfg == nil {
		fmt.Println("警告：日志模块初始化时配置为空，使用默认日志级别 DEBUG")
		SetLogLevel(LevelDebug)
		return
	}
	SetLogLevel(ParseLogLevel(Cfg.LogLevel))
}

func logWithColor(level LogLevel, colorAttr color.Attribute, prefix string, format string, args ...interface{}) {
	mu.Lock()
	defer mu.Unlock()

	if level < currentLevel {
		return
	}

	c := color.New(colorAttr).SprintFunc()
	timestamp := time.Now().Format("2006-01-02 15:04:05")
	msg := fmt.Sprintf(format, args...)
	logger.Output(3, fmt.Sprintf("%s [%s] %s", timestamp, c(prefix), msg))
}

func Debug(format string, args ...interface{}) {
	logWithColor(LevelDebug, color.FgCyan, "DEBUG", format, args...)
}

func Info(format string, args ...interface{}) {
	logWithColor(LevelInfo, color.FgGreen, "INFO", format, args...)
}

func Warn(format string, args ...interface{}) {
	logWithColor(LevelWarn, color.FgYellow, "WARN", format, args...)
}

func Error(format string, args ...interface{}) {
	logWithColor(LevelError, color.FgRed, "ERROR", format, args...)
}

func SetOutput(w interface{}) {
	mu.Lock()
	defer mu.Unlock()
	if logger != nil {
		if output, ok := w.(os.File); ok {
			logger.SetOutput(&output)
		}
	}
}
