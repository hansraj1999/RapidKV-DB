package models

import (
	"fmt"
	"sync"
)

type KeyDir struct {
	RecordSize     int
	RecordPosition int64
	FileName       string
	Timestamp      int64
}

var once sync.Once
var InMemoryStorage map[string]KeyDir
var mu sync.Mutex // Mutex for synchronizing memory writes

// init initializes InMemoryStorage
func Init() {
	once.Do(func() {
		InMemoryStorage = make(map[string]KeyDir)
		fmt.Println("InMemoryStorage initialized")
	})
}

func AddToMemory(FileName, Key string, RecordSize int, RecordPosition int64, Timestamp int64) {
	mu.Lock()

	_KeyDir := KeyDir{
		RecordSize:     RecordSize,
		RecordPosition: RecordPosition,
		FileName:       FileName,
		Timestamp:      Timestamp,
	}
	InMemoryStorage[Key] = _KeyDir
	defer mu.Unlock()
	fmt.Print("Added to memory: ", InMemoryStorage[FileName], "\n")

}
func GetDataFromMemory(key string) (KeyDir, bool) {
	mu.Lock()

	value, exists := InMemoryStorage[key]
	fmt.Print("Getting from memory: ", value, "\n")
	defer mu.Unlock()
	if !exists {
		return value, exists
	}
	return value, true
}
