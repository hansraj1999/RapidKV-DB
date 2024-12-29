package storage

import (
	"encoding/binary"
	"fmt"
	"io"
	"os"
	"rapidkv-db/utils"
)

// AppendToLog appends a key-value pair to the log file and returns the offset and record size.
func AppendToLog(fileName, key, value string) (int64, int, int64, error) {
	// Calculate metadata
	crc := utils.CalculateCRC(key, value)
	timestamp := utils.GetTimestamp()
	keySize := int32(len(key))
	valueSize := int32(len(value))
	recordSize := CRCSize + TimestampSize + KeySizeSize + ValueSizeSize + len(key) + len(value)

	// Open log file
	logFile, err := os.OpenFile(fmt.Sprintf("%s/%s", LOG_FILES_DIR, fileName), os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0666)
	if err != nil {
		return 0, 0, 0, fmt.Errorf("failed to open log file: %w", err)
	}
	defer logFile.Close()

	// Get current file offset
	offset, err := logFile.Seek(0, io.SeekEnd)
	if err != nil {
		return 0, 0, 0, fmt.Errorf("failed to seek log file: %w", err)
	}

	// Write log entry
	if err := WriteLogEntry(logFile, crc, timestamp, keySize, valueSize, key, value); err != nil {
		return 0, 0, 0, fmt.Errorf("failed to write log entry: %w", err)
	}

	return offset, recordSize, timestamp, nil
}

// WriteLogEntry writes a log entry with metadata and data to the log file.
func WriteLogEntry(file *os.File, crc uint32, timestamp int64, keySize, valueSize int32, key, value string) error {
	writer := func(data interface{}) error {
		return binary.Write(file, binary.LittleEndian, data)
	}

	// Write metadata fields
	if err := writer(crc); err != nil {
		return fmt.Errorf("failed to write CRC: %w", err)
	}
	if err := writer(timestamp); err != nil {
		return fmt.Errorf("failed to write timestamp: %w", err)
	}
	if err := writer(keySize); err != nil {
		return fmt.Errorf("failed to write key size: %w", err)
	}
	if err := writer(valueSize); err != nil {
		return fmt.Errorf("failed to write value size: %w", err)
	}

	// Write key and value
	if _, err := file.Write([]byte(key)); err != nil {
		return fmt.Errorf("failed to write key: %w", err)
	}
	if _, err := file.Write([]byte(value)); err != nil {
		return fmt.Errorf("failed to write value: %w", err)
	}

	return nil
}
