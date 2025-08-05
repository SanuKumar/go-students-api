package main

import (
	"context"
	"log"
	"log/slog"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/sanukumar/go-students-api/internal/config"
	"github.com/sanukumar/go-students-api/internal/http/handlers/student"
)

func main() {
	// load config
	cfg := config.MustLoad()

	// set custom logger -> if needed

	// database setup

	// setup router
	router := http.NewServeMux() // assign type automatically

	router.HandleFunc("POST /api/students", student.New()) // request object pointer

	// setup server

	server := http.Server{
		Addr:    cfg.Addr,
		Handler: router,
	}

	slog.Info("Server started", slog.String("address", cfg.Addr))
	// fmt.Printf("Server started %s", cfg.Addr)

	// ============== // graceful shutdown

	// channels --> they are uesd for synchronization
	done := make(chan os.Signal, 1) // can store only 1 value

	// signal
	signal.Notify(done, os.Interrupt, syscall.SIGINT, syscall.SIGTERM)

	// create go routine
	go func() {
		err := server.ListenAndServe()
		if err != nil {
			log.Fatal("failed to start server")
		}
	}()

	<-done

	slog.Info("Shutting down the server...")

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second) // create a context with timeout
	defer cancel()

	err := server.Shutdown(ctx) // graceful shutdown
	if err != nil {
		slog.Error("failed to shutdown server..", slog.String("error", err.Error()))
	}

	// if err := server.Shutdown(ctx); err != nil {
	// 	slog.Error("failed to shutdown server..", slog.String("error", err.Error()))
	// }

	slog.Info("Server shutdown successfully...")

}
