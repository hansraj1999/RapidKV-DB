package main

import (
	"rapidkv-db/handlers"

	"github.com/gin-gonic/gin"
)

const LOG_FILES_DIR = "./data"

var initial_file_name = "log_1"

func main() {
	router := gin.Default()
	router.POST("/put", handlers.PutHandler)
	router.GET("/get", handlers.GetHandler)
	// router.POST("/merge", handlers.MergeHandler)
	// router.GET("/health", handlers.HealthHandler)
	// listen and serve on 0.0.0.0:9090
	router.Run(":9090")
}
