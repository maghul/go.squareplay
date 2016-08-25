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
	"github.com/natefinch/lumberjack"
)

var airplayers *raopd.SinkCollection

var ilog *logger
var dlog *logger
var logfilename = ""

func main() {
	var port int
	var profile int
	flag.IntVar(&port, "w", 6111, "The server port for the proxy")
	flag.IntVar(&profile, "pprof", 0, "Set to a port to enable profiling")
	flag.StringVar(&logfilename, "l", "", "Set the logfile name, if omitted log to stderr")
	flag.Parse()

	if logfilename != "" {
		ljl := &lumberjack.Logger{
			Filename:   logfilename,
			MaxSize:    500, // megabytes
			MaxBackups: 3,
			MaxAge:     28, //days
		}
		ljl.Rotate()
		ilog = makeLogger("squareplay", ljl)
	} else {
		ilog = makeLogger("", os.Stderr)
	}

	ilog.Println("Starting SquarePlay Proxy 0.0.1(beta)")

	if profile > 0 {
		go http.ListenAndServe(fmt.Sprintf(":%d", profile), nil)
	}

	var err error
	airplayers, err = raopd.NewSinkCollection()
	if err != nil {
		panic(err)
	}

	initSignals()
	startServer(port)

	slog.Info.Println("Stopping Squareplay...")
	airplayers.Close()
}

func shutdownAll() {
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
			ilog.Println("Received Signal: ", s)
			if s == os.Interrupt || s == os.Kill || s == syscall.SIGTERM {
				shutdownAll()
			}
		}
	}()
	startServer(port)

	ilog.Println("Closing...")
	airplayers.Close()
}

func shutdownAll() {
	ln.Close()
	airplayers.Close()
	time.Sleep(100 * time.Millisecond)
	os.Exit(0)
}

/*
SquarePlay:
An AirPlay to squeezeserver proxy. Will start a web-server which can server audio, coverart and
metadata for airplay sessions.
It also support configuation and dynamic service handling.

The web URIs supported are found at http://<thisserver>/doc


*/
