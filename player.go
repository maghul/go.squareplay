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

	"github.com/maghul/go.raopd"
)

type SqueezePlayer struct {
	playerHandler *http.ServeMux
	serviceInfo   *raopd.SinkInfo
	pipeReader    io.Reader
	apService     *raopd.Source
	coverData     []byte
	metaData      string

	h    *host
	name string
}

var allSqueezePlayers = newSyncMap()

func (sp *SqueezePlayer) addHandlerFunc(suburl string, handler func(http.ResponseWriter, *http.Request)) {
	sp.playerHandler.HandleFunc(fmt.Sprintf("/%s/%s", sp.Id(), suburl), handler)
}

func startPlayer(name, id string, h *host) (*SqueezePlayer, error) {
	ff := func() (interface{}, error) {
		hwaddr, err := net.ParseMAC(id)
		if err != nil {
			return nil, err
		}
		si := &raopd.SinkInfo{
			SupportsCoverArt: true,
			SupportsMetaData: "JSON",
			Name:             name,
			HardwareAddress:  hwaddr,
			Port:             0,
		}

		sp := &SqueezePlayer{nil, si, nil, nil, nil, "", h, name}
		sp.initPlayer()

		sp.apService, err = airplayers.Register(sp)
		if err != nil {
			return nil, err
		}

		return sp, nil
	}
	spi, err := allSqueezePlayers.Get(id, ff)
	if err != nil {
		return nil, err
	}
	return spi.(*SqueezePlayer), err
}

func stopPlayer(name, id string) error {
	spi, err := allSqueezePlayers.Remove(id)
	if err != nil {
		return err
	}
	sp := spi.(*SqueezePlayer)

	sp.shutdown()

	return nil
}

func (sp *SqueezePlayer) Id() string {
	return sp.serviceInfo.HardwareAddress.String()
}

func (sp *SqueezePlayer) Name() string {
	return sp.serviceInfo.Name
}

func (sp *SqueezePlayer) initPlayer() {
	url := fmt.Sprintf("/%s/", sp.Id())
	sp.playerHandler = http.NewServeMux()
	dlog.Println("URL=", url)

	sp.addHandlerFunc("metadata.json", sp.metadata)
	sp.addHandlerFunc("cover.jpg", sp.cover)
	sp.addHandlerFunc("audio.pcm", sp.audio)
	sp.addHandlerFunc("audio.wav", sp.audio)
	sp.addHandlerFunc("time/", sp.seek)
	sp.addHandlerFunc("control/volume/", sp.volume)
	sp.addHandlerFunc("control/", sp.control)
}

func (sp *SqueezePlayer) close() {
}

func (sp *SqueezePlayer) shutdown() {
	airplayers.Unregister(sp)
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

	audioCtx := context.Background()
	dlog.Println("!!!!!!!!!!!!!!!!!!  Waiting for audio pipe to shut down", r.RemoteAddr)
	sp.apService.NewAudioStream(audioCtx, conn)
	<-audioCtx.Done()
	dlog.Println("!!!!!!!!!!!!!!!!!!  Audio pipe was shut down", r.RemoteAddr)
}

func (sp *SqueezePlayer) getTheCommand(w http.ResponseWriter, r *http.Request) string {
	bw := bufio.NewWriter(w)
	url := r.URL.String()
	dlog.Println("URL=", url)
	is := strings.LastIndex(url, "/")
	if is < 0 {
		dlog.Println("URL=", "bye")
		w.WriteHeader(400)
		return ""
	}
	url = url[is+1:]
	dlog.Println("URL=", url)
	bw.Flush()

	return url
}

func (sp *SqueezePlayer) control(w http.ResponseWriter, r *http.Request) {
	sp.apService.Command(sp.getTheCommand(w, r))
}

func (sp *SqueezePlayer) volume(w http.ResponseWriter, r *http.Request) {
	cmd := sp.getTheCommand(w, r)
	dlog.Println("volume command ", cmd)
	switch cmd {
	case "relative":
		dlog.Println("Setting volume to relative for ", sp)
		sp.apService.VolumeMode(false)
	case "absolute":
		dlog.Println("Setting volume to absolute for ", sp)
		sp.apService.VolumeMode(true)
	default:
		v, err := toRaopVolume(cmd)
		if err != nil {
			ilog.Println("Error converting volume ", cmd, ":", err)
			w.WriteHeader(400)
			return
		}
		sp.apService.Volume(v)
	}
}

func (sp *SqueezePlayer) notifyString(data string) {
	buf := bytes.NewBufferString("{ \"")
	buf.WriteString(sp.Id())
	buf.WriteString("\":")
	buf.WriteString(data)
	buf.WriteString("}")
	// TODO: Send the client as part of the notification to avoid slushing bytes about
	sp.h.txNotification(buf.Bytes())
}

func (sp *SqueezePlayer) notify(data []byte) {
	buf := bytes.NewBufferString("{ \"")
	buf.WriteString(sp.Id())
	buf.WriteString("\":")
	buf.Write(data)
	buf.WriteString("}")
	// TODO: Send the client as part of the notification to avoid slushing bytes about
	sp.h.txNotification(buf.Bytes())
}

// --- raopd.Sink implementation

// Get the service info for the service.
func (sp *SqueezePlayer) Info() *raopd.SinkInfo {
	return sp.serviceInfo
}

func (sp *SqueezePlayer) SetCoverArt(mimetype string, content []byte) {
	dlog.Println("LoadCoverArt:", mimetype, " buffer size=", len(content))
	sp.coverData = content
	sp.notifyString("\"coverart\"")
}

func (sp *SqueezePlayer) SetMetadata(metadata string) {
	sp.metaData = metadata
	sp.notifyString(metadata)
}

// Set the volume of the output device. The volume value may be an absolute
// value from 0 - 100, or it may be up down values using UP=1000 and DOWN=-1000
func (sp *SqueezePlayer) SetVolume(volume float32) {
	switch volume {
	case 1000:
		sp.notifyString("{ \"volume\": \"+2\" }")
	case -1000:
		sp.notifyString("{ \"volume\": \"-2\" }")
	default:
		sp.notifyString(fmt.Sprintf("{ \"volume\": %d }", int(ios2decVolume(volume))))
	}
}

// Shows the progress of the track in milliseconds.
// pos is the current position, length is the total length of the current track
func (sp *SqueezePlayer) SetProgress(pos, length int) {
	sp.notifyString(fmt.Sprintf("{ \"progress\": { \"current\": %d, \"length\": %d }}", pos, length))
}

// Called when a device has connected, specifically the DACP/Control connection.
func (sp *SqueezePlayer) Connected(name string) {
	sp.notifyString(fmt.Sprintf("{ \"source\": \"%s\" }", name))
}

// Called when the stream is started.
func (sp *SqueezePlayer) Play() {
	sp.notifyString("\"play\"")
}

// Called when the stream is paused
func (sp *SqueezePlayer) Pause() {
	sp.notifyString("\"pause\"")
}

// Called when the connection to source is terminated
func (sp *SqueezePlayer) Stopped() {
	sp.notifyString("\"stop\"")
}

// Called when the sink has been removed
func (sp *SqueezePlayer) Closed() {
}

// A name for the player
func (sp *SqueezePlayer) String() string {
	return sp.name
}
