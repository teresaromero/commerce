package main

import (
	"context"
	"ecommerce-api/internal/config"
	"ecommerce-api/internal/db"
	"ecommerce-api/internal/router"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"

	"github.com/gin-gonic/gin"
)

func main() {
	ctx := context.Background()
	cfg, err := config.LoadConfig()
	if err != nil {
		log.Fatalf("Error loading config: %v", err)
	}

	dbConn, err := db.Init(cfg.DbURL)
	if err != nil {
		log.Fatalf("Error initializing database: %v", err)
	}

	engine := gin.Default()

	// load routes
	router.Init(engine, cfg.JwtSecret)

	srv := &http.Server{
		Addr:    ":8080",
		Handler: engine.Handler(),
	}

	// Run the server in a goroutine so that it doesn't block
	go func() {
		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatalf("listen: %s\n", err)
		}
	}()

	// Wait for interrupt signal to gracefully shutdown the server with
	quit := make(chan os.Signal, 1)
	// kill (no params) by default sends syscall.SIGTERM
	// kill -2 is syscall.SIGINT
	// kill -9 is syscall.SIGKILL but can't be caught, so don't need add it
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit
	log.Println("Gracefully Shutdown Server ...")

	// TODO: add any cleanup tasks here, like closing database connections
	if err := dbConn.Close(); err != nil {
		log.Println("Error closing database connection:", err)
	}

	if err := srv.Shutdown(ctx); err != nil {
		log.Println("Server Shutdown:", err)
	}

}
