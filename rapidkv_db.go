package main

import (
	"fmt"
	"rapidkv-db/internal/storage"
	"rapidkv-db/internal/storage/models"
	"time"
)

const LOG_FILES_DIR = "./data"

var initial_file_name = "log_1"

func main() {
	keyDir := make(map[string]models.KeyDir)

	var key string
	var value string
	fmt.Print("Enter key value")
	fmt.Scan(&key, &value)
	var in_memory_storage models.KeyDir
	in_memory_storage.File_id = initial_file_name
	in_memory_storage.Timestamp = time.Now().Unix()

	offset, recordSize, err := storage.AppendToLog(initial_file_name, key, value)
	if err != nil {
		fmt.Println("Error:", err)
	} else {
		fmt.Print("Data appended. Offset:", offset, " , recordSize:", recordSize)
	}
	in_memory_storage.RecordSize = recordSize
	in_memory_storage.RecordPosition = offset
	fmt.Println("In memory storage", in_memory_storage)

	keyDir[key] = in_memory_storage
	fmt.Println("Key Directory", keyDir)
	var v string
	var e error
	v, e = storage.ReadFromLog(keyDir[key].File_id, key, keyDir[key].RecordPosition)
	if e != nil {
	}
	fmt.Println(v, "value")
}
