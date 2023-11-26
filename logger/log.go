package logger

import (
	"context"
	"fmt"
	"log"
	"os"
	"time"
)

const (
	logFlag   = "LOG"
	infoFlag  = "INFO"
	errorFlag = "ERROR"
	debugFlag = "DEBUG"
	warnFlag  = "WARN"
	traceFlag = "TRACE"
)

const (
	colorReset   = "\033[0m"
	colorRed     = "\033[31m"
	colorGreen   = "\033[32m"
	colorYellow  = "\033[33m"
	colorMagenta = "\033[35m"
)

var (
	logFile   *os.File
	logger    *log.Logger
	logToFile bool
)

func InitLogger(isLogToFile bool, filePath string) error {
	logToFile = isLogToFile

	if isLogToFile {
		// open the log file
		var err error
		logFile, err = os.OpenFile(filePath, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
		if err != nil {
			return err
		}

		// initialize the logger to write to the log file
		logger = log.New(logFile, "", log.LstdFlags)
	} else {
		// initialize the logger to write to the console
		logger = log.New(os.Stdout, "", log.LstdFlags)
	}

	return nil
}

func Log(ctx interface{}, message string, data interface{}) {
	logMessage(getContext(ctx), logFlag, colorReset, message, data)
}

func Info(ctx interface{}, message string, data interface{}) {
	logMessage(getContext(ctx), infoFlag, colorGreen, message, data)
}

func Error(ctx interface{}, message string, data interface{}) {
	logMessage(getContext(ctx), errorFlag, colorRed, message, data)
}

func Warn(ctx interface{}, message string, data interface{}) {
	logMessage(getContext(ctx), warnFlag, colorYellow, message, data)
}

func Debug(ctx interface{}, message string, data interface{}) {
	logMessage(getContext(ctx), debugFlag, colorMagenta, message, data)
}

func Trace(ctx interface{}, message string, data interface{}) {
	logMessage(getContext(ctx), traceFlag, colorReset, message, data)
}

func logMessage(ctx context.Context, level string, color string, message string, data interface{}) {
	traceID := getTraceID(ctx)
	spanID := getSpanID(ctx)
	logTime := time.Now().Format(time.RFC3339Nano)

	logEntry := fmt.Sprintf("Function: %s | Trace ID: %s | Span ID: %s | Level: %s | Message: %s | Data: %+v | Timestamp: %s",
		getFunctionName(), traceID, spanID, level, message, data, logTime)

	// Apply color to the log entry
	logEntry = fmt.Sprintf("%s%s%s", color, logEntry, colorReset)

	if logToFile {
		logger.Println(logEntry)
	} else {
		fmt.Println(logEntry)
	}
}
