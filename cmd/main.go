package main

import (
	"context"
	"fmt"
	"net/http"
	"os"
	"os/signal"
	"path/filepath"
	"syscall"
	"time"

	"github.com/baleegh-ud-din/ascend/config"
	"github.com/baleegh-ud-din/ascend/database"
	"github.com/baleegh-ud-din/ascend/jobs"
	"github.com/baleegh-ud-din/ascend/routes"
	"github.com/baleegh-ud-din/ascend/utils"
)

func main() {
	// Load Config
	cfg := config.Cfg

	// Initiate Logger
	logger := utils.NewLogger()

	// Connect Database
	database.Connect(cfg)
	defer database.Close()

	// Define Routes
	mux := http.NewServeMux()
	routes.SetUpRoutes(mux)

	// Serve HTML
	mux.HandleFunc("/", ServeUI)

	// Configure Server
	server := &http.Server{
		Addr:    fmt.Sprintf(":%s", cfg.AppPort),
		Handler: mux,
	}

	// Start Server
	go func() {
		msg := fmt.Sprintf("🚀 Starting Server on Port %s", server.Addr)
		logger.Info(msg)
		if err := server.ListenAndServe(); err != http.ErrServerClosed {
			msg = fmt.Sprintf("⚠️ Error Starting Server: %s", err)
			logger.Error(msg)
			os.Exit(1)
		}
	}()

	// Start Jobs
	jobs.StartJobs()
	defer jobs.StopJobs()

	// Graceful Shutdown
	GracefulShutdown(server)
}

func ServeUI(w http.ResponseWriter, r *http.Request) {
	path := r.URL.Path
	fullPath := filepath.Join("./static", path)

	// Check if the requested file exists
	if _, err := os.Stat(fullPath); os.IsNotExist(err) {
		// If the file does not exist, serve index.html for dynamic routes
		http.ServeFile(w, r, filepath.Join("./static", "index.html"))
		return
	}

	// Serve the static file if it exists
	http.ServeFile(w, r, fullPath)
}

func GracefulShutdown(server *http.Server) {
	logger := utils.NewLogger()
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit

	fmt.Println("")
	logger.Warning("⚠️  Server is shutting down...")

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	if err := server.Shutdown(ctx); err != nil {
		msg := fmt.Sprintf("⚠️  Server forced to shutdown: %v\n", err)
		logger.Error(msg)
	}

	logger.Info("🏠 Server exited properly")
}
