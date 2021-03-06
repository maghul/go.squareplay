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

	initDefault(serverMux)
	initUsage(serverMux)
	initControl(serverMux)
	initLogHandler(serverMux)

	mux := &LogHandler{LogAdminHandler(os.Stderr, serverMux)}

	srv := &http.Server{
		Addr:           fmt.Sprintf(":%d", port),
		Handler:        mux,
		ReadTimeout:    10 * time.Second,
		WriteTimeout:   0, // TODO: This should only be set for chunked persistent connection / notifications
		MaxHeaderBytes: 1 << 20,
		ErrorLog:       stdlog.New(slog.Info, "HTTP:", 0),
	}

	slog.Info.Println("serving Squareplay HTTP at port:", port)
	var err error
	ln, err = net.Listen("tcp", srv.Addr)
	if err != nil {
		panic(err)
	}
	srv.Serve(ln)
}

func initDefault(mux *http.ServeMux) {
	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		url := r.URL.Path[1:]
		if url == "" {
			w.Header().Add("Location", "/index.html")
			w.WriteHeader(302)
		} else if len(url) > 17 {
			purl := string(url[0:17])
			players := allSqueezePlayers.snapshot()
			player, ok := players[purl]
			if ok {
				sp := (player).(*SqueezePlayer)

				sp.playerHandler.ServeHTTP(w, r)
			} else {
				w.WriteHeader(503)
			}
		}
	})
}

type LogHandler struct {
	h http.Handler
}

func (lh *LogHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	slog.Debug.Println("request: ", r.URL)
	lh.h.ServeHTTP(w, r)
}
