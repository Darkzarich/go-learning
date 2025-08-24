package main

import (
	"log"
)

type LogLevel int

const (
	LogLevelError LogLevel = iota
	LogLevelWarning
	LogLevelInfo
)

type ExtendedLogger struct {
	*log.Logger
	logLevel LogLevel
}

func NewLogger() *ExtendedLogger {
	return &ExtendedLogger{
		logLevel: LogLevelInfo,
		Logger:   log.New(log.Writer(), "", log.LstdFlags),
	}
}

func (l *ExtendedLogger) SetLogLevel(level LogLevel) {
	l.logLevel = level
}

func (l *ExtendedLogger) Infoln(s string) {
	l.Println(LogLevelInfo, "INFO", s)
}

func (l *ExtendedLogger) Warnln(s string) {
	l.Println(LogLevelWarning, "WARN", s)
}

func (l *ExtendedLogger) Errorln(s string) {
	l.Println(LogLevelError, "ERROR", s)
}

func (l *ExtendedLogger) Println(level LogLevel, prefix string, s string) {
	if l.logLevel >= level {
		l.Printf("[%s] %s\n", prefix, s)
	}
}

func main() {
	logger := NewLogger()
	logger.SetLogLevel(LogLevelWarning)
	logger.Infoln("Should not be printed")
	logger.Warnln("Hello")
	logger.Errorln("World")
}
