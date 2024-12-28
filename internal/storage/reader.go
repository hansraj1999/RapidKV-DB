package storage

import (
	"encoding/binary"
	"fmt"
	"io"
	"os"
)

// ReadFromLog reads a key-value pair from the log file starting at a given offset.
func ReadFromLog(fileName string, searchKey string, offset int64) (string, error) {
	// Open the log file
	logFile, err := os.OpenFile(fmt.Sprintf("%s/%s", LOG_FILES_DIR, fileName), os.O_RDONLY, 0666)
	if err != nil {
		return "", fmt.Errorf("failed to open log file: %w", err)
	}
	defer logFile.Close()

	// Seek to the specified offset
	if _, err = logFile.Seek(offset, io.SeekStart); err != nil {
		return "", fmt.Errorf("failed to seek to offset: %w", err)
	}

	// Read metadata: CRC, timestamp, keySize, valueSize
	var crc uint32
	var timestamp int64
	var keySize, valueSize int32

	if err := readBinary(logFile, &crc); err != nil {
		return "", fmt.Errorf("failed to read CRC: %w", err)
	}
	if err := readBinary(logFile, &timestamp); err != nil {
		return "", fmt.Errorf("failed to read timestamp: %w", err)
	}
	if err := readBinary(logFile, &keySize); err != nil {
		return "", fmt.Errorf("failed to read key size: %w", err)
	}
	if err := readBinary(logFile, &valueSize); err != nil {
		return "", fmt.Errorf("failed to read value size: %w", err)
	}

	// Read the key
	key := make([]byte, keySize)
	if _, err := io.ReadFull(logFile, key); err != nil {
		return "", fmt.Errorf("failed to read key: %w", err)
	}

	// Read the value
	value := make([]byte, valueSize)
	if _, err := io.ReadFull(logFile, value); err != nil {
		return "", fmt.Errorf("failed to read value: %w", err)
	}

	// Check if the key matches the search key
	if string(key) == searchKey {
		return string(value), nil
	}

	return "", fmt.Errorf("key not found")
}

// readBinary simplifies binary.Read calls
func readBinary(file *os.File, data interface{}) error {
	return binary.Read(file, binary.LittleEndian, data)
}