package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strings"

	"github.com/maghul/go.slf"
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
	mux.HandleFunc("/control/stop", stop)
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
		slog.Debug.Println("Error decoding '", r.Body, "': err=", err)
		w.WriteHeader(500)
		fmt.Fprintln(w, "decodePlayers: Error decoding '", r.Body, "': err=", err)
	}
	return pa
}

func start(w http.ResponseWriter, r *http.Request) {
	host := getHost(r.Host)
	w.Header().Add("Content-Type", "text/text")
	fmt.Fprintf(w, "[\r\n")
	for _, pj := range decodePlayers(w, r) {
		slog.Info.Println("Starting player: ", pj, " from host", host)
		_, err := startPlayer(pj.Name, pj.Id, host)
		if err != nil {
			slog.Info.Printf("Player %s[%s] could not be started: %v", pj.Name, pj.Id, err)
			w.WriteHeader(404)
			fmt.Fprintf(w, "{ \"%s\": \"Failed\" }\r\n", pj.Id)
		} else {
			slog.Info.Printf("Player %s[%s] started", pj.Name, pj.Id)
			fmt.Fprintf(w, "{ \"%s\": \"OK\" }\r\n", pj.Id)
		}
	}
	fmt.Fprintf(w, "]\r\n")
}

func stop(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("Content-Type", "text/text")
	fmt.Fprintf(w, "[\r\n")
	for _, pj := range decodePlayers(w, r) {
		err := stopPlayer(pj.Name, pj.Id)
		if err != nil {
			w.WriteHeader(404)
			slog.Info.Printf("Player %s[%s] could not be stopped: %v", pj.Name, pj.Id, err)
			fmt.Fprintf(w, "{ \"%s\": \"Failed\" }\r\n", pj.Id)
		} else {
			slog.Info.Printf("Player %s[%s] stopped", pj.Name, pj.Id)
			fmt.Fprintf(w, "{ \"%s\": \"Bye\" }\r\n", pj.Id)
		}
	}
	fmt.Fprintf(w, "]\r\n")
}

func notify(w http.ResponseWriter, r *http.Request) {
	broadcastNotifications()
}

func handleLogger(w http.ResponseWriter, r *http.Request) {
	url := r.URL.String()
	is := strings.LastIndex(url, "/")
	if is < 0 {
		slog.Debug.Println("handleLogger: Failed to decode URL=", url)
		w.WriteHeader(400)
		fmt.Fprintln(w, "handleLogger: Failed to decode URL=", url)
		return
	}
	ss := strings.Split(url[is:], "?")
	err := slf.SetLevel(ss[0], ss[1])
	if err != nil {
		slog.Debug.Println("handleLogger: ", err)
		w.WriteHeader(400)
		fmt.Fprintln(w, "handleLogger: ", err)
		return
	}

}

func handleRaopDebug(w http.ResponseWriter, r *http.Request) {
	url := r.URL.String()
	is := strings.LastIndex(url, "/")
	if is < 0 {
		slog.Debug.Println("Failed to decode URL=", url)
		w.WriteHeader(400)
		fmt.Fprintln(w, "handleLogger: Failed to decode URL=", url)
		return
	}
	ss := strings.Split(url[is:], "?")
	logger := strings.ToLower(ss[0])
	level := strings.ToLower(ss[1])

	switch level {
	case "true":
		raopdDebug(logger, true)
	case "false":
		raopdDebug(logger, false)

	}

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
		slog.Debug.Println("Notification: ", string(not))
		w.Write(not)
		w.(http.Flusher).Flush()

	}
}

func broadcastNotifications() {
	// Broadcast all cached notifications.
}
