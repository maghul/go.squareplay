package main

import (
	"log"
	"net/http"
	"time"
)

var serverMux *http.ServeMux

func main() {
	serverMux = http.NewServeMux()

	initUsage(serverMux)

	startPlayer("Hejsan", "123")

	s := &http.Server{
		Addr:           ":8082",
		Handler:        serverMux,
		ReadTimeout:    10 * time.Second,
		WriteTimeout:   10 * time.Second,
		MaxHeaderBytes: 1 << 20,
	}

	log.Fatal(s.ListenAndServe())
}

/*
SquarePlay:
An AirPlay to squeezeserver proxy. Will start a web-server which can server audio, coverart and
metadata for airplay sessions.
It also support configuation and dynamic service handling.

The web URIs supported are

*/
