package logger

import (
	"log"
	"os"
)

type Logger interface {
	Info(message string)
	Error(message string)
	Fatal(message string)
}

type DefaultLogger struct {
	localLog  *os.File
	globalLog *os.File
}

func NewDefaultLogger(localLogFilePath string, globalLogFilePath string) (*DefaultLogger, error) {
	localLog, err := os.OpenFile(localLogFilePath, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	if err != nil {
		return nil, err
	}

	globalLog, err := os.OpenFile(globalLogFilePath, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	if err != nil {
		localLog.Close()
		return nil, err
	}

	return &DefaultLogger{
		localLog:  localLog,
		globalLog: globalLog,
	}, nil
}

func (l *DefaultLogger) Info(message string) {
	log.New(l.localLog, "[INFO] ", log.LstdFlags).Println(message)
	log.New(l.globalLog, "[INFO] ", log.LstdFlags).Println(message)
}

func (l *DefaultLogger) Error(message string) {
	log.New(l.localLog, "[ERROR] ", log.LstdFlags).Println(message)
	log.New(l.globalLog, "[ERROR] ", log.LstdFlags).Println(message)
}

func (l *DefaultLogger) Fatal(message string) {
	log.New(l.localLog, "[FATAL] ", log.LstdFlags).Fatal(message)
	log.New(l.globalLog, "[FATAL] ", log.LstdFlags).Fatal(message)
}

func (l *DefaultLogger) Close() error {
	localErr := l.localLog.Close()
	globalErr := l.globalLog.Close()

	if localErr != nil {
		return localErr
	}
	return globalErr
}
