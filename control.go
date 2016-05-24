package main

import (
	"fmt"
	"net/http"
)

func initControl(mux *http.ServeMux) {

	mux.HandleFunc("/control/start", start)
	mux.HandleFunc("/control/notify", notify)
	mux.HandleFunc("/control/logger", handleLogger)
	mux.HandleFunc("/notifications.json", notifications)
}

func start(w http.ResponseWriter, r *http.Request) {
	id := r.Header.Get("Airplay-Session-Id")
	name := r.Header.Get("Airplay-Session-Name")

	_, ok := allSqueezePlayers[id]
	if ok {
		fmt.Fprintf(w, "Player %s[%s] already started\r\n", name, id)
		w.WriteHeader(404)
	} else {
		_, err := startPlayer(name, id)
		if err != nil {
			fmt.Fprintf(w, "Player %s[%s] could not be started: %v\r\n", name, id, err)
			w.WriteHeader(404)
		} else {
			fmt.Fprintf(w, "Player %s[%s] started\r\n", name, id)
		}
	}
}

func notify(w http.ResponseWriter, r *http.Request) {
	broadcastNotifications()
}

func handleLogger(w http.ResponseWriter, r *http.Request) {
	// This is to set logging levels.
}

var notificationChannel = make(chan string, 32)

func notifications(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("Transfer-Encoding", "chunked")
	w.Header().Add("Content-Type", "text/text")
	w.Header().Add("Content-Length", "should not be set")
	fmt.Fprintf(w, "Initial notification...\r\n")
	w.(http.Flusher).Flush()
	for {
		not := <-notificationChannel
		fmt.Fprintf(w, "%s\r\n", not)
		w.(http.Flusher).Flush()

	}
}

func broadcastNotifications() {
	// Broadcast all cached notifications.
}
