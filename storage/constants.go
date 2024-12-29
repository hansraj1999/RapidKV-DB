package storage

const LOG_FILES_DIR = "./data"

const (
	CRCSize       = 4
	TimestampSize = 8
	KeySizeSize   = 4
	ValueSizeSize = 4
)

const (
	MaxLogFileSize = 100 * 1024 * 1024 // 100MB
)
