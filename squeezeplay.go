package main

import (
	"flag"
	"fmt"

	"github.com/maghul/go.raopd"
)

var apServiceRegistry *raopd.ServiceRegistry

func main() {
	var port int
	var profile int
	flag.IntVar(&port, "w", 6111, "The server port for the proxy")
	flag.IntVar(&profile, "pprof", 0, "Set to a port to enable profiling")
	flag.Parse()

	log.Info().Println("Starting SquarePlay Proxy 0.0.1(beta)")

	if profile > 0 {
		go http.ListenAndServe(fmt.Sprintf(":%d", profile), nil)
	}

	var err error
	apServiceRegistry, err = raopd.NewServiceRegistry(getKeyfile())
	if err != nil {
		panic(err)
	}

	startServer(port)
}

/*
SquarePlay:
An AirPlay to squeezeserver proxy. Will start a web-server which can server audio, coverart and
metadata for airplay sessions.
It also support configuation and dynamic service handling.

The web URIs supported are found at http://<thisserver>/doc


*/
