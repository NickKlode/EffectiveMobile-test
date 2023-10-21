package main

import (
	"context"
	httpserver "emobletest/internal/http-server"
	"emobletest/internal/lib/logger"
	"emobletest/internal/storage/postgres"
	"fmt"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/joho/godotenv"
	"golang.org/x/exp/slog"
)

func main() {

	log := setupLogger()

	if err := godotenv.Load(); err != nil {
		log.Error("error while loading .env file", logger.Err(err))
		os.Exit(1)

	}

	connStr := fmt.Sprintf("user=%s password=%s host=%s port=%s dbname=%s sslmode=%s",
		os.Getenv("POSTGRES_USER"),
		os.Getenv("POSTGRES_PASSWORD"),
		os.Getenv("POSTGRES_HOST"),
		os.Getenv("POSTGRES_PORT"),
		os.Getenv("POSTGRES_DB"),
		os.Getenv("POSTGRES_SSL_MODE"),
	)
	db, err := postgres.New(connStr)
	if err != nil {
		log.Error("error while connecting to db", logger.Err(err))
		os.Exit(1)
	}
	api := httpserver.New(db, log)

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGTERM, syscall.SIGINT, os.Interrupt)

	srv := &http.Server{
		Addr:         os.Getenv("ADDRESS"),
		Handler:      api.Router(),
		ReadTimeout:  10 * time.Second,
		WriteTimeout: 10 * time.Second,
		IdleTimeout:  30 * time.Second,
	}
	go func() {
		if err := srv.ListenAndServe(); err != nil {
			log.Error("error while starting server", logger.Err(err))
			os.Exit(1)
		}
	}()
	log.Info("server started")
	<-quit

	log.Info("stopping server")

	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()

	if err := srv.Shutdown(ctx); err != nil {
		log.Error("failed to stop server", logger.Err(err))
		os.Exit(1)
	}
	log.Info("server stopped")
}

func setupLogger() *slog.Logger {
	log := slog.New(
		slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{Level: slog.LevelDebug}),
	)
	return log
}
