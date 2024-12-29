package main

import (
	"context"
	"fmt"
	"net/http"
	"os"
	"os/signal"
	"rapidkv-db/handlers"
	"rapidkv-db/models"
	"syscall"
	"time"

	"github.com/gin-gonic/gin"
)

const LOG_FILES_DIR = "./data"

func main() {
	models.Init()
	router := gin.Default()
	router.POST("/put", handlers.PutHandler)
	router.GET("/get", handlers.GetHandler)
	// router.POST("/merge", handlers.MergeHandler)
	// router.GET("/health", handlers.HealthHandler)
	// listen and serve on 0.0.0.0:9090
	server := &http.Server{
		Addr:    ":9090",
		Handler: router,
	}
	go func() {
		if err := server.ListenAndServe(); err != nil {
			fmt.Println("Server stopped:", err)
		}
	}()

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	if err := server.Shutdown(ctx); err != nil {
		fmt.Println("Failed to gracefully shutdown:", err)
	}
}
