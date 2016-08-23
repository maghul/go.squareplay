package main

import (
	"io"
	"net/http"
)

type logHandler struct {
	out io.Writer
	wh  http.Handler
}

func LogHandler(out io.Writer, wh http.Handler) http.Handler {
	lh := &logHandler{out, wh}
	return lh
}

func (lh *logHandler) ServeHTTP(resp http.ResponseWriter, req *http.Request) {
	log.Debug().Println("ServeHTTP: url=", req.URL, ", client=", req.RemoteAddr)
	lh.wh.ServeHTTP(resp, req)
}
