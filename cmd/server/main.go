package main

import (
	"context"
	"errors"
	"evildethow/httpserver/internal/handler/middleware"
	"evildethow/httpserver/internal/handler/ping"
	"log"
	"net/http"
	"os"
	"os/signal"
)

func main() {
	srv := &http.Server{
		Addr:    ":8080",
		Handler: middleware.Logger(http.DefaultServeMux),
	}

	idleConnsClosed := make(chan struct{})
	go func() {
		sigint := make(chan os.Signal, 1)
		signal.Notify(sigint, os.Interrupt)
		<-sigint

		// We received an interrupt signal, shut down.
		if err := srv.Shutdown(context.Background()); err != nil {
			// Error from closing listeners, or context timeout:
			log.Printf("HTTP server Shutdown: %v", err)
		}
		close(idleConnsClosed)
	}()

	// Handlers
	http.HandleFunc("/ping", ping.Handler)

	if err := srv.ListenAndServe(); !errors.Is(err, http.ErrServerClosed) {
		// Error starting or closing listener:
		log.Fatalf("HTTP server ListenAndServe: %v", err)
	}

	<-idleConnsClosed
}
