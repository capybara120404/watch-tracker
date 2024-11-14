package utils

import (
	"fmt"
	"log"
	"os"
)

func CreateLogger(logFile, nameOfLogger string) (*log.Logger, *os.File, error) {
	err := os.MkdirAll("common/logs", 0755)
	if err != nil {
		return nil, nil, fmt.Errorf("failed to create directories for log file %s: %v", logFile, err)
	}

	file, err := os.OpenFile(logFile, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		return nil, nil, fmt.Errorf("failed to open log file %s: %w", logFile, err)
	}

	logger := log.New(file, nameOfLogger, log.LstdFlags|log.Lshortfile)

	return logger, file, nil
}
