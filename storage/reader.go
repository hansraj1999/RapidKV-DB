package storage

import (
	"encoding/binary"
	"fmt"
	"io"
	"os"
	"rapidkv-db/models"
)

// ReadFromLog reads a key-value pair from the log file starting at a given offset.
func GetValueFromDB(searchKey string) (string, error) {
	_value, exists := models.GetDataFromMemory(searchKey)
	var crc uint32
	var timestamp int64
	var keySize, valueSize int32

	if !exists {
		return "", fmt.Errorf("key not found from memory")
	}
	// Open the log file
	logFilePath := fmt.Sprintf("%s/%s", LOG_FILES_DIR, _value.FileName)
	logFile, err := os.OpenFile(logFilePath, os.O_RDONLY, 0666)
	if err != nil {
		fmt.Print("Failed to open log file: ", err)
		return "", fmt.Errorf("failed to open log file '%s': %w", logFilePath, err)
	}
	defer logFile.Close()

	// Seek to the specified offset
	if _, err = logFile.Seek(_value.RecordPosition, io.SeekStart); err != nil {
		return "", fmt.Errorf("failed to seek to offset: %w", err)
	}
	// Read metadata: CRC, timestamp, keySize, valueSize
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
	fmt.Print("key: ", string(key), " searchKey: ", searchKey)
	if string(key) == searchKey {
		return string(value), nil
	}

	return "", fmt.Errorf("key mismatch: expected '%s', found '%s'", searchKey, string(key))

}

// readBinary simplifies binary.Read calls
func readBinary(file *os.File, data interface{}) error {
	return binary.Read(file, binary.LittleEndian, data)
}
