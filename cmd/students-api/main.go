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
	"github.com/sanukumar/go-students-api/internal/storage/sqlite"
)

func main() {
	// load config
	cfg := config.MustLoad()

	// set custom logger -> if needed

	// database setup
	storage, err := sqlite.New(cfg)
	if err != nil {
		log.Fatal(err)
	}

	slog.Info("storage initialized", slog.String("env", cfg.Env), slog.String("version", "1.0.0"))

	// setup router
	router := http.NewServeMux() // assign type automatically

	router.HandleFunc("POST /api/students", student.New(storage)) // request object pointer
	router.HandleFunc("GET /api/students/{id}", student.GetById(storage))
	router.HandleFunc("GET /api/students", student.GetList(storage))

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

	if err := server.Shutdown(ctx); // graceful shutdown
	err != nil {
		slog.Error("failed to shutdown server..", slog.String("error", err.Error()))
	}

	// if err := server.Shutdown(ctx); err != nil {
	// 	slog.Error("failed to shutdown server..", slog.String("error", err.Error()))
	// }

	slog.Info("Server shutdown successfully...")

}
