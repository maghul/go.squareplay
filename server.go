package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"time"
)

var serverMux *http.ServeMux

func startServer(port int) {
	serverMux = http.NewServeMux()

	initUsage(serverMux)
	initControl(serverMux)

	mux := LogHandler(os.Stderr, serverMux)

	s := &http.Server{
		Addr:           fmt.Sprintf(":%d", port),
		Handler:        mux,
		ReadTimeout:    10 * time.Second,
		WriteTimeout:   0, // TODO: This should only be set for chunked persistent connection / notifications
		MaxHeaderBytes: 1 << 20,
	}

	log.Fatal(s.ListenAndServe())
}
