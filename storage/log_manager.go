package storage

import (
	"fmt"
	"os"
	"time"
)

// LogFileManager manages the current active log file
type LogFileManager struct {
	CurrentLogFile string
}

func NewLogFileManager() *LogFileManager {
	return &LogFileManager{
		CurrentLogFile: "log_1.log", // start with default name log_1.log should we move to constantss?
	}
}

// RotateLogFile rotates the current log file when it exceeds MaxLogFileSize
func (lm *LogFileManager) RotateLogFile() error {
	newFileName := fmt.Sprintf("log_%d.log", time.Now().Unix())
	logFilePath := fmt.Sprintf("%s/%s", LOG_FILES_DIR, newFileName)

	// Create a new log file
	file, err := os.OpenFile(logFilePath, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0666)
	if err != nil {
		return fmt.Errorf("failed to rotate log file: %w", err)
	}
	defer file.Close()

	// Update current log file
	lm.CurrentLogFile = newFileName
	return nil
}

// GetCurrentLogFile returns the current active log file
func (lm *LogFileManager) GetCurrentLogFile() string {

	return lm.CurrentLogFile
}
