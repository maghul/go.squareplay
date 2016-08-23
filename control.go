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

	fmt.Println("Starting client: name=", name, ", id=", id)

	host := getHost(r.Host)
	_, err := startPlayer(name, id, host)
	if err != nil {
		w.WriteHeader(404)
		fmt.Fprintf(w, "Player %s[%s] could not be started: %v\r\n", name, id, err)
	} else {
		fmt.Fprintf(w, "Player %s[%s] started\r\n", name, id)
	}
}

func notify(w http.ResponseWriter, r *http.Request) {
	broadcastNotifications()
}

func handleLogger(w http.ResponseWriter, r *http.Request) {
	// This is to set logging levels.
}

func notifications(w http.ResponseWriter, r *http.Request) {
	h := getHost(r.Host)

	w.Header().Add("Transfer-Encoding", "chunked")
	w.Header().Add("Content-Type", "text/text")
	//	w.Header().Add("Content-Length", "should not be set")
	fmt.Fprintf(w, "{ \"serverStatus\": \"OK\" }\r\n")
	w.(http.Flusher).Flush()

	for {
		fmt.Println("Waiting for notifications...")
		not := <-h.notificationChannel
		fmt.Println("SENDING NOTIFICATION: ", string(not))
		w.Write(not)
		w.(http.Flusher).Flush()

	}
}

func broadcastNotifications() {
	// Broadcast all cached notifications.
}
