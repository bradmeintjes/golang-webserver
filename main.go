package main

import (
	"go/webserver/homepage"
	"go/webserver/internal/server"
	"log"
	"net/http"
	"os"
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
