package main

import (
	"bufio"
	"fmt"
	"net/http"
	"strings"
)

type SqueezePlayer struct {
	name          string
	mac           string
	playerHandler *http.ServeMux
}

var allSqueezePlayers = make(map[string]*SqueezePlayer)

func (sp *SqueezePlayer) addHandlerFunc(suburl string, handler func(http.ResponseWriter, *http.Request)) {
	sp.playerHandler.HandleFunc(fmt.Sprintf("/%s/%s", sp.mac, suburl), handler)
}

func startPlayer(name, id string) (*SqueezePlayer, error) {
	name = strings.Replace(name, " ", "", -1)
	sp := &SqueezePlayer{name, id, nil}
	allSqueezePlayers[id] = sp
	sp.initPlayer(serverMux)
	return sp, nil
}

func (sp *SqueezePlayer) initPlayer(mux *http.ServeMux) {
	url := fmt.Sprintf("/%s/", sp.mac)
	sp.playerHandler = http.NewServeMux()
	fmt.Println("URL=", url)
	mux.Handle(url, sp.playerHandler)

	sp.addHandlerFunc("metadata.json", sp.metadata)
	sp.addHandlerFunc("cover.jpg", sp.cover)
	sp.addHandlerFunc("audio.pcm", sp.audio)
	sp.addHandlerFunc("audio.wav", sp.audio)
	sp.addHandlerFunc("control/", sp.control)
}

func (sp *SqueezePlayer) close() {
}

func (sp *SqueezePlayer) metadata(w http.ResponseWriter, r *http.Request) {
	bw := bufio.NewWriter(w)
	fmt.Fprintf(bw, "This is a test....  %s", r.URL)
	bw.Flush()
}
func (sp *SqueezePlayer) cover(w http.ResponseWriter, r *http.Request) {
	bw := bufio.NewWriter(w)
	fmt.Fprintf(bw, "This is a test....  %s", r.URL)
	bw.Flush()
}
func (sp *SqueezePlayer) audio(w http.ResponseWriter, r *http.Request) {
	bw := bufio.NewWriter(w)
	fmt.Fprintf(bw, "This is a test....  %s", r.URL)
	bw.Flush()
}
func (sp *SqueezePlayer) control(w http.ResponseWriter, r *http.Request) {
	bw := bufio.NewWriter(w)
	url := r.URL.String()
	fmt.Println("URL=", url)
	is := strings.LastIndex(url, "/")
	if is < 0 {
		fmt.Println("URL=", "bye")
		w.WriteHeader(400)
		return
	}
	url = url[is+1:]
	fmt.Println("URL=", url)
	bw.Flush()
}
