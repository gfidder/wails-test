package log

import (
	"log"
	"os"
)

type aggregatedLogger struct {
	infoLogger  *log.Logger
	warnLogger  *log.Logger
	errorLogger *log.Logger
}

var enabled bool
var al aggregatedLogger

func Enable(logFile *os.File) {
	enabled = true
	flags := log.LstdFlags | log.Lshortfile
	infoLogger := log.New(logFile, "INFO: ", flags)
	warnLogger := log.New(logFile, "WARN: ", flags)
	errorLogger := log.New(logFile, "ERROR: ", flags)

	al = aggregatedLogger{
		infoLogger:  infoLogger,
		warnLogger:  warnLogger,
		errorLogger: errorLogger,
	}
}

func Info(message string) {
	if !enabled {
		return
	}
	al.infoLogger.Println(message)
}

func Warn(message string) {
	if !enabled {
		return
	}
	al.warnLogger.Println(message)
}

func Error(message string) {
	if !enabled {
		return
	}
	al.errorLogger.Fatalln(message)
}
