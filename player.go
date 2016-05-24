package main

import (
	"bufio"
	"bytes"
	"context"
	"fmt"
	"io"
	"net"
	"net/http"
	"os"
	"strings"
	"time"

	"github.com/maghul/go.raopd"
)

type SqueezePlayer struct {
	playerHandler *http.ServeMux
	serviceInfo   *raopd.ServiceInfo
	pipeWriter    io.Writer
	pipeReader    io.Reader
	apService     *raopd.ServiceRef
	coverData     []byte
	metaData      string
}

var allSqueezePlayers = make(map[string]*SqueezePlayer)

func (sp *SqueezePlayer) addHandlerFunc(suburl string, handler func(http.ResponseWriter, *http.Request)) {
	sp.playerHandler.HandleFunc(fmt.Sprintf("/%s/%s", sp.Id(), suburl), handler)
}

func startPlayer(name, id string) (*SqueezePlayer, error) {
	name = strings.Replace(name, " ", "", -1)
	hwaddr, err := net.ParseMAC(id)
	if err != nil {
		return nil, err
	}
	si := &raopd.ServiceInfo{
		SupportsAbsoluteVolume: false,
		SupportsRelativeVolume: false,
		SupportsCoverArt:       true,
		SupportsMetaData:       "JSON",
		Name:                   name,
		HardwareAddress:        hwaddr,
		Port:                   0,
	}

	sp := &SqueezePlayer{nil, si, nil, nil, nil, nil, ""}
	allSqueezePlayers[id] = sp

	sp.apService, err = apServiceRegistry.RegisterService(sp)
	if err != nil {
		return nil, err
	}

	sp.initPlayer(serverMux)
	return sp, nil
}

func (sp *SqueezePlayer) Id() string {
	return sp.serviceInfo.HardwareAddress.String()
}

func (sp *SqueezePlayer) Name() string {
	return sp.serviceInfo.Name
}

func (sp *SqueezePlayer) initPlayer(mux *http.ServeMux) {
	url := fmt.Sprintf("/%s/", sp.Id())
	sp.playerHandler = http.NewServeMux()
	fmt.Println("URL=", url)
	mux.Handle(url, sp.playerHandler)

	sp.addHandlerFunc("metadata.json", sp.metadata)
	sp.addHandlerFunc("cover.jpg", sp.cover)
	sp.addHandlerFunc("audio.pcm", sp.audio)
	sp.addHandlerFunc("audio.wav", sp.audio)
	sp.addHandlerFunc("control/", sp.control)
	sp.addHandlerFunc("control/volume/", sp.volume)
}

func (sp *SqueezePlayer) close() {
}

func (sp *SqueezePlayer) metadata(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("Content-Type", "text/javascript")
	io.WriteString(w, sp.metaData)
}

func (sp *SqueezePlayer) cover(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("Content-Type", "image/jpeg")
	w.Write(sp.coverData)
}

func (sp *SqueezePlayer) audio(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(os.Stderr, "Starting audio\n")
	hj, ok := w.(http.Hijacker)
	if !ok {
		http.Error(w, "webserver doesn't support hijacking", http.StatusInternalServerError)
		return
	}

	// OK, hijack the connection and start transferring PCM.
	// TODO: Ensure we get at a sample boundary or there might be hell of
	//       a racket coming from the speakers...
	conn, bufrw, err := hj.Hijack()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	// Don't forget to close the connection:
	defer conn.Close()

	mimetype := "audio/x-pcm"

	bufrw.WriteString("HTTP/1.1 200 OK\r\n")
	//	bufrw.WriteString("Transfer-Encoding: binary\r\n")
	bufrw.WriteString("Content-Type: ")
	bufrw.WriteString(mimetype)
	bufrw.WriteString("\r\n")
	bufrw.WriteString("\r\n")
	bufrw.Flush()
	sp.pipeWriter = bufrw
	for {
		time.Sleep(100 * 1000000)
		//		bufrw.Flush()
	}
}

func (sp *SqueezePlayer) getTheCommand(w http.ResponseWriter, r *http.Request) string {
	bw := bufio.NewWriter(w)
	url := r.URL.String()
	fmt.Println("URL=", url)
	is := strings.LastIndex(url, "/")
	if is < 0 {
		fmt.Println("URL=", "bye")
		w.WriteHeader(400)
		return ""
	}
	url = url[is+1:]
	fmt.Println("URL=", url)
	bw.Flush()

	return url
}

func (sp *SqueezePlayer) control(w http.ResponseWriter, r *http.Request) {
	sp.apService.Command(sp.getTheCommand(w, r))
}

func (sp *SqueezePlayer) volume(w http.ResponseWriter, r *http.Request) {

	sp.apService.Volume(sp.getTheCommand(w, r))
}

func (sp *SqueezePlayer) notifyString(data string) {
	buf := bytes.NewBufferString("")
	buf.WriteRune('"')
	buf.WriteString(data)
	buf.WriteRune('"')
	sp.notify(buf.Bytes())
}

func (sp *SqueezePlayer) notify(data []byte) {
	buf := bytes.NewBufferString("{ \"")
	buf.WriteString(sp.Id())
	buf.WriteString("\":")
	buf.Write(data)
	buf.WriteString("}")
	// TODO: Send the client as part of the notification to avoid slushing bytes about
	notificationChannel <- buf.Bytes()
}

// --- raopd.Sink implementation

// Get the service info for the service.
func (sp *SqueezePlayer) ServiceInfo() *raopd.ServiceInfo {
	return sp.serviceInfo
}

// Get a writer for the audio stream. Only raw PCM with two channel
// 16-bit depth at 44100 samples/second is currently supported.
func (sp *SqueezePlayer) AudioWriter() io.Writer {
	return sp.pipeWriter
}

func (sp *SqueezePlayer) AudioWriterErr(err error) {
	fmt.Println("ERROR Writing audio: ", err)
	sp.pipeWriter = nil
}

func (sp *SqueezePlayer) SetCoverArt(mimetype string, content []byte) {
	fmt.Println("LoadCoverArt:", mimetype, " buffer size=", len(content))
	sp.coverData = content
	sp.notifyString("coverart")
}

func (sp *SqueezePlayer) SetMetadata(metadata []byte) {
	sp.metaData = metadata
	sp.notifyString(metadata)
}

// Set the volume of the output device. The volume value may be an absolute
// value from 0 - 100, or it may be up down values using UP=1000 and DOWN=-1000
func (sp *SqueezePlayer) SetVolume(volume float32) {
}

// Shows the progress of the track in milliseconds.
// pos is the current position, length is the total length of the current track
func (sp *SqueezePlayer) SetProgress(pos, length int) {
}

// Called when the stream is started.
func (sp *SqueezePlayer) Play() {
	sp.notifyString("play")
}

// Called when the stream is paused
func (sp *SqueezePlayer) Pause() {
	sp.notifyString("pause")
}

// Called when the connection to source is terminated
func (sp *SqueezePlayer) Close() {
	sp.notifyString("stop")
}
