package utils

import (
	"hash/crc32"
)

func CalculateCRC(key string, value string) uint32 {
	return crc32.ChecksumIEEE([]byte(key + value))
}
