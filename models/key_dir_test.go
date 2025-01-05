package models

import (
	"strconv"
	"sync"
	"testing"
	"time"
)

func TestAddToMemory_Concurrency(t *testing.T) {
	// Initialize the InMemoryStorage
	Init()

	// Number of concurrent goroutines to run
	numGoroutines := 100
	var wg sync.WaitGroup

	// Launch multiple goroutines that will write to the memory concurrently
	for i := 0; i < numGoroutines; i++ {
		wg.Add(1)
		go func(i int) {
			defer wg.Done()

			// Simulate different data for each goroutine
			FileName := "File_" + strconv.Itoa(i) // Fix here
			Key := "Key_" + strconv.Itoa(i)       // Fix here
			RecordSize := i
			RecordPosition := int64(i)
			Timestamp := time.Now().Unix()

			AddToMemory(FileName, Key, RecordSize, RecordPosition, Timestamp)
		}(i)
	}

	// Wait for all goroutines to finish
	wg.Wait()

	// Verify the memory contains the expected keys
	for i := 0; i < numGoroutines; i++ {
		Key := "Key_" + strconv.Itoa(i) // Fix here
		if _, exists := GetDataFromMemory(Key); !exists {
			t.Errorf("Expected key %s to be in memory, but it was not found", Key)
		}
	}
}
