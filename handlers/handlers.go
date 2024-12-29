package handlers

import (
	"fmt"
	"net/http"
	"rapidkv-db/models"
	"rapidkv-db/storage"

	"github.com/gin-gonic/gin"
)

func PutHandler(c *gin.Context) {
	var payload struct {
		Key   string `json:"key"`
		Value string `json:"value"`
	}

	if err := c.ShouldBindJSON(&payload); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid JSON payload"})
		return
	}

	offset, recordSize, timestamp, fileName, err := storage.AppendToLog(payload.Key, payload.Value)
	fmt.Println(offset, recordSize, timestamp, "Got from AppendToLog")
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	models.AddToMemory(fileName, payload.Key, recordSize, offset, timestamp)
	c.JSON(http.StatusOK, gin.H{"message": "Key-Value pair added successfully"})
}

// GetHandler handles GET requests to retrieve a value by key
func GetHandler(c *gin.Context) {
	key := c.Query("key")
	if key == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Key query parameter is required"})
		return
	}

	value, err := storage.GetValueFromDB(key)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": err})
		return
	}

	c.JSON(http.StatusOK, gin.H{"key": key, "value": value})
}

// // MergeHandler triggers the merge process
// func MergeHandler(c *gin.Context) {
// 	go db.DB.Merge() // Run merge in a separate goroutine
// 	c.JSON(http.StatusOK, gin.H{"message": "Merge process started"})
// }

// // HealthHandler checks if the server is running
// func HealthHandler(c *gin.Context) {
// 	c.JSON(http.StatusOK, gin.H{"status": "ok"})
// }
