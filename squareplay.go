package main

import (
	"flag"
	"fmt"
	"net/http"
	_ "net/http/pprof"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/maghul/go.raopd"
)

var airplayers *raopd.SinkCollection

//var ilog *logger
//var dlog *logger
var logfilename = ""
var keyfilename = ""

func raopdDebug(name string, value interface{}) {
	slog.Debug.Println("Setting RAOP debug: ", name)
	err := raopd.Debug(name, value)
	if err != nil {
		slog.Info.Println("Could not set RAOP Debug '", name, "': ", err)
	}
}

func main() {
	var port int
	var profile int
	flag.IntVar(&port, "p", 6111, "The server port for the proxy")
	flag.IntVar(&profile, "pprof", 0, "Set to a port to enable profiling")
	flag.StringVar(&logfilename, "l", "", "Set the logfile name, if omitted log to stderr")
	flag.StringVar(&keyfilename, "k", "", "Set the keyfile path")
	flag.Parse()

	initLogging()
	slog.Info.Println("Starting SquarePlay 0.0.3")

	if profile > 0 {
		go http.ListenAndServe(fmt.Sprintf(":%d", profile), nil)
	}

	if len(keyfilename) == 0 {
		slog.Info.Println("No keyfile specified. exiting")
		return
	}
	var err error
	slog.Info.Printf("Starting with keyfile='%s'", keyfilename)
	airplayers, err = raopd.NewSinkCollection(keyfilename)
	if err != nil {
		panic(err)
	}

	initSignals()
	startServer(port)

	slog.Info.Println("Stopping Squareplay...")
	airplayers.Close()
}

func shutdownAll() {
	slog.Debug.Println("shutdownAll...")
	ln.Close()
	airplayers.Close()
	time.Sleep(100 * time.Millisecond)
	os.Exit(0)
}

func initSignals() {
	sc := make(chan os.Signal)
	signal.Notify(sc)
	go func() {
		for {
			s := <-sc
			slog.Debug.Println("Received Signal: ", s)
			if s == os.Interrupt || s == os.Kill || s == syscall.SIGTERM {
				shutdownAll()
			}
		}
	}()
}

/*
SquarePlay:
An AirPlay to squeezeserver proxy. Will start a web-server which can server audio, coverart and
metadata for airplay sessions.
It also support configuation and dynamic service handling.

The web URIs supported are found at http://<thisserver>/doc


*/
