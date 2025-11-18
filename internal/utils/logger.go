// Package utils provides shared helpers such as logger, file handlers, etc.
package utils

import (
	"fmt"
	"log"
	"os"
	"time"
)

// LogLevel represents the level of a log message.
type LogLevel int

const (
	INFO LogLevel = iota
	WARN
	ERROR
	SUCCESS
)

var (
	logFile *os.File   // active log file
	Logger  *log.Logger // writer for log file
)

// InitLogger creates/opens the main reconx.log file.
func InitLogger() {
	var err error

	logFile, err = os.OpenFile("reconx.log", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		fmt.Println("Failed to create log file:", err)
		return
	}

	// Logger uses no prefix; timestamps are handled manually.
	Logger = log.New(logFile, "", 0)
}

// Log prints a message to stdout + writes it to reconx.log.
func Log(level LogLevel, msg string) {
	timestamp := time.Now().Format("2006-01-02 15:04:05")

	var prefix string
	switch level {
	case INFO:
		prefix = "[INFO]"
	case WARN:
		prefix = "[WARN]"
	case ERROR:
		prefix = "[ERROR]"
	case SUCCESS:
		prefix = "[OK]"
	}

	fmt.Printf("%s %s %s\n", timestamp, prefix, msg)

	if Logger != nil {
		Logger.Printf("%s %s %s\n", timestamp, prefix, msg)
	}
}
