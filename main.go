package main

import (
	"context"
	"go/webserver/homepage"
	"go/webserver/server"
	"log"
	"net/http"
	"os"
	"os/signal"
	"time"
)

var (
	serverAddress = os.Getenv("SERVER_ADDR")
	certFile      = os.Getenv("TLS_CERT")
	certKey       = os.Getenv("TLS_CERT_KEY")

	logger = log.New(os.Stdout, "appname", log.LstdFlags)
)

func main() {
	logger.Println("Server is starting...")

	router := http.NewServeMux()

	home := homepage.New(logger)
	home.SetupRoutes(router)

	svr := server.New(logger, router, serverAddress)
	serve(svr)
}

func serve(svr *http.Server) {
	done := make(chan bool)
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, os.Interrupt)

	go func() {
		<-quit
		logger.Println("Server is shutting down...")

		ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
		defer cancel()

		svr.SetKeepAlivesEnabled(false)
		if err := svr.Shutdown(ctx); err != nil {
			logger.Fatalf("Could not gracefully shutdown the server: %v\n", err)
		}
		close(done)
	}()

	isSecure := len(certFile) != 0 && len(certKey) != 0
	logger.Println("Server is ready to handle requests at", serverAddress)

	var err error
	if isSecure {
		err = svr.ListenAndServeTLS(certFile, certKey)
	} else {
		err = svr.ListenAndServe()
	}

	if err != nil && err != http.ErrServerClosed {
		logger.Fatalf("Could not listen on %s: %v\n", serverAddress, err)
	}

	<-done
	logger.Println("Server stopped")
}
