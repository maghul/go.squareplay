package main

import (
	"flag"
	"fmt"

	"github.com/maghul/go.raopd"
)

var apServiceRegistry *raopd.ServiceRegistry

func main() {
	fmt.Println("Starting SquarePlay Proxy 0.0.1(beta)")

	var port int
	flag.IntVar(&port, "w", 6111, "The server port for the proxy")
	flag.Parse()

	var err error
	apServiceRegistry, err = raopd.NewServiceRegistry()
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
