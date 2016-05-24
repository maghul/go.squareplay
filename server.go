package main

import (
	"fmt"
	stdlog "log"
	"net"
	"net/http"
	"os"
	"time"
)

var serverMux *http.ServeMux
var ln net.Listener

func startServer(port int) {
	serverMux = http.NewServeMux()

	initUsage(serverMux)
	initControl(serverMux)

	mux := LogHandler(os.Stderr, serverMux)

	srv := &http.Server{
		Addr:           fmt.Sprintf(":%d", port),
		Handler:        mux,
		ReadTimeout:    10 * time.Second,
		WriteTimeout:   0, // TODO: This should only be set for chunked persistent connection / notifications
		MaxHeaderBytes: 1 << 20,
	}

	log.Info.Println("serving at ", GetOutboundIP(), ":", port)
	//	log.Info.Println(srv.ListenAndServe())
	var err error
	ln, err = net.Listen("tcp", srv.Addr)
	if err != nil {
		panic(err)
	}
	srv.Serve(ln)

	log.Info.Println("Bye bye")
}
