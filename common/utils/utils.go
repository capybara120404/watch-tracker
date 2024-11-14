package utils

import (
	"fmt"
	"log"
	"os"
	"path/filepath"
)

func CreateLogger(logFile, nameOfLogger string) (*log.Logger, *os.File, error) {
	logFilePath := filepath.Join("common", "logs", logFile)

	err := os.MkdirAll(filepath.Dir(logFilePath), 0755)
	if err != nil {
		return nil, nil, fmt.Errorf("failed to create directories for log file %s: %v", logFilePath, err)
	}

	file, err := os.OpenFile(logFilePath, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		return nil, nil, fmt.Errorf("failed to open log file %s: %w", logFilePath, err)
	}

	logger := log.New(file, nameOfLogger, log.LstdFlags|log.Lshortfile)

	return logger, file, nil
}
