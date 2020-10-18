package main

import (
	"godockermicroservice/server"
	"godockermicroservice/homepage"
	"fmt"
	"net/http"
	"log"
	"os"
)

var (
	GomsCertFile = os.Getenv("GOMS_CERT_FILE")
	GomsKeyFile = os.Getenv("GOMS_KEY_FILE")
	GomsServiceAddr = os.Getenv("GOMS_SERVICE_ADDR")
)

const message = "Hello"

func main() {
	logger := log.New(os.Stdout, "goms ", log.LstdFlags | log.Lshortfile)

	// dependency injection in go
	h := homepage.NewHandlers(logger)

	mux := http.NewServeMux()
	h.SetupRoutes(mux)

	srv := server.NewServer(mux, GomsServiceAddr)

	logger.Println("Server starting")
	err :=	srv.ListenAndServeTLS(GomsCertFile, GomsKeyFile)
	if err != nil {
		logger.Fatalf("server failed to start : %v", err)
	}
	fmt.Println(message)
}
