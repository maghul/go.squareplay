package main

import (
	"encoding/json"
	"fmt"
	"net/http"

	"golang.org/x/text/encoding/charmap"
)

type playerJson struct {
	Name string
	Id   string
}

var decoder = charmap.ISO8859_1.NewDecoder()

func initControl(mux *http.ServeMux) {

	mux.HandleFunc("/control/restart", restart)
	mux.HandleFunc("/control/start", start)
	mux.HandleFunc("/control/notify", notify)
	mux.HandleFunc("/control/logger", handleLogger)
	mux.HandleFunc("/notifications.json", notifications)
}

func restart(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Restarting squareplay server\r\n")
	shutdownAll()
}

func getHeader(r *http.Request, name string) string {

	iv := r.Header.Get(name)
	dv, err := decoder.String(iv)
	if err != nil {
		panic(err)
	}
	return dv
}

func (pj *playerJson) String() string {
	return fmt.Sprintf("player:[{'Name': '%s', 'Id': '%s'}]", pj.Name, pj.Id)
}

func decodePlayers(w http.ResponseWriter, r *http.Request) []*playerJson {
	pa := make([]*playerJson, 0)
	d := json.NewDecoder(r.Body)
	err := d.Decode(&pa)
	if err != nil {
		ilog.Println("Error decoding start request:", err)
		w.WriteHeader(500)
	}
	return pa
}

func start(w http.ResponseWriter, r *http.Request) {
	ilog.Println("start client")

	host := getHost(r.Host)
	for _, pj := range decodePlayers(w, r) {
		ilog.Println("Starting client: ", pj, " from host", host)
		_, err := startPlayer(pj.Name, pj.Id, host)
		if err != nil {
			w.WriteHeader(404)
			fmt.Fprintf(w, "Player %s[%s] could not be started: %v\r\n", pj.Name, pj.Id, err)
		} else {
			fmt.Fprintf(w, "Player %s[%s] started\r\n", pj.Name, pj.Id)
		}
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

	nc := make(chan []byte, 32)
	h.addNotificationChannel(nc)
	defer h.removeNotificationChannel(nc)

	for {
		not := <-nc
		dlog.Println("SENDING NOTIFICATION: ", string(not))
		w.Write(not)
		w.(http.Flusher).Flush()

	}
}

func broadcastNotifications() {
	// Broadcast all cached notifications.
}
