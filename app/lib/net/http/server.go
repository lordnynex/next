package http

import (
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"time"
)

// ListenAndServe serves the handler at specified port
// and shuts down server gracefully in a shutdown timeout.
func ListenAndServe(addr string, handler http.Handler, shutdownTimeout time.Duration) {
	server := startServer(handler, addr)
	shutdownServerGracefully(server, shutdownTimeout)
}

func startServer(handler http.Handler, addr string) *http.Server {
	log.Print("Server started on ", addr)

	server := &http.Server{
		Addr:    addr,
		Handler: handler,
	}

	go func() {
		if err := server.ListenAndServe(); err != nil {
			// Cannot panic, because this probably is an intentional close.
			log.Print(err)
		}
	}()

	return server
}

func shutdownServerGracefully(server *http.Server, shutdownTimeout time.Duration) {
	// Wait for interrupt signal to gracefully shutdown the server
	// with a specified timeout.
	quit := make(chan os.Signal)
	signal.Notify(quit, os.Interrupt)
	<-quit

	log.Print("Shutting down the server...")

	ctx, cancel := context.WithTimeout(context.Background(), shutdownTimeout)
	defer cancel()

	if err := server.Shutdown(ctx); err != nil {
		panic(err)
	}

	log.Print("Server gracefully stopped.")
}
