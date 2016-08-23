package main

import (
	"flag"
	"fmt"
	"net/http"
	_ "net/http/pprof"
	"os"

	"github.com/maghul/go.raopd"
	"github.com/natefinch/lumberjack"
)

var airplayers *raopd.SinkCollection

var ilog *logger
var dlog *logger
var logfilename = ""

func main() {
	// TODO: add option
	logfilename = "/var/log/squeezeboxserver/squareplay.log"
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

	var port int
	var profile int
	flag.IntVar(&port, "w", 6111, "The server port for the proxy")
	flag.IntVar(&profile, "pprof", 0, "Set to a port to enable profiling")
	flag.Parse()

	ilog.Println("Starting SquarePlay Proxy 0.0.1(beta)")

	if profile > 0 {
		go http.ListenAndServe(fmt.Sprintf(":%d", profile), nil)
	}

	var err error
	airplayers, err = raopd.NewSinkCollection()
	if err != nil {
		panic(err)
	}

	startServer(port)

	ilog.Println("Closing...")
	airplayers.Close()
}

func shutdownAll() {
	ln.Close()
}

/*
SquarePlay:
An AirPlay to squeezeserver proxy. Will start a web-server which can server audio, coverart and
metadata for airplay sessions.
It also support configuation and dynamic service handling.

The web URIs supported are found at http://<thisserver>/doc


*/
