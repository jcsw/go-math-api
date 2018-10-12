package syslog

import (
	"log"
	"os"
)

func newSyslog() *log.Logger {
	return log.New(os.Stdout, "go-math-api ", log.LstdFlags)
}

var syslog = newSyslog()

// Logger return the logger
func Logger() *log.Logger {
	return syslog
}

// Debug - Logging in level DEBUG
func Debug(log string, v ...interface{}) {
	syslog.Printf("DEBUG "+log, v...)
}

// Info - Logging in level INFO
func Info(log string, v ...interface{}) {
	syslog.Printf("INFO  "+log, v...)
}

// Warn - Logging in level WARN
func Warn(log string, v ...interface{}) {
	syslog.Printf("WARN  "+log, v...)
}

// Error - Logging in level ERROR
func Error(log string, v ...interface{}) {
	syslog.Printf("ERROR "+log, v...)
}

// Fatal - Logging in level FATAL
func Fatal(log string, v ...interface{}) {
	syslog.Fatalf("FATAL  "+log, v...)
}
