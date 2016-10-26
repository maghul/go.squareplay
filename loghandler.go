package main

import (
	"bufio"
	"bytes"
	"fmt"
	"io"
	"net/http"
	"strings"

	"github.com/maghul/go.raopd"
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
	fmt.Fprintln(lh.out, "ServeHTTP: url=", req.URL, ", client=", req.RemoteAddr)
	lh.wh.ServeHTTP(resp, req)
}

type logSetting struct {
	name     string
	setting  string
	settings string
	setter   func(ls *logSetting)
}

var logSettings = []*logSetting{
	&logSetting{"sequencetrace", "off", "off|on", setRaopTraceLogging},
	&logSetting{"volumetrace", "off", "off|on", setRaopTraceLogging},

	&logSetting{"raopd.dacp", "off", "off|info|debug", setRaopLogging},
	&logSetting{"raopd.dmap", "off", "off|info|debug", setRaopLogging},
	&logSetting{"raopd.auth", "off", "off|info|debug", setRaopLogging},
	&logSetting{"raopd.net", "off", "off|info|debug", setRaopLogging},
	&logSetting{"raopd.raop", "off", "off|info|debug", setRaopLogging},
	&logSetting{"raopd.rtp", "off", "off|info|debug", setRaopLogging},
	&logSetting{"raopd.rtsp", "off", "off|info|debug", setRaopLogging},
	&logSetting{"raopd.sequencer", "off", "off|info|debug", setRaopLogging},
	&logSetting{"raopd.volume", "off", "off|info|debug", setRaopLogging},
	&logSetting{"raopd.zeroconf", "off", "off|info|debug", setRaopLogging},
}

func setRaopLogging(ls *logSetting) {
	dlr := fmt.Sprint("log.debug/", ls.name)
	ilr := fmt.Sprint("log.info/", ls.name)
	switch ls.setting {
	case "off":
		raopd.Debug(ilr, nil)
		raopd.Debug(dlr, nil)
	case "info":
		raopd.Debug(ilr, ilog)
		raopd.Debug(dlr, nil)
	case "frbug":
		raopd.Debug(ilr, ilog)
		raopd.Debug(dlr, dlog)
	}
}

func setRaopTraceLogging(ls *logSetting) {
	switch ls.setting {
	case "off":
		raopd.Debug(ls.name, false)
	case "on":
		raopd.Debug(ls.name, true)
	}

}

func makeSelect(name, selected, options string) string {
	buf := bytes.NewBufferString("")
	fmt.Fprintf(buf, "<select name=\"%s\">\n", name)
	for _, option := range strings.Split(options, "|") {
		if option == selected {
			fmt.Fprintf(buf, "<option selected value=\"%s\">%s</option>\n", option, strings.Title(option))
		} else {
			fmt.Fprintf(buf, "<option value=\"%s\">%s</option>\n", option, strings.Title(option))
		}
	}
	fmt.Fprintln(buf, "</select>")
	return buf.String()
}

func (ls *logSetting) makeSelect() string {
	return makeSelect(ls.name, ls.setting, ls.settings)
}

func initLogHandler(mux *http.ServeMux) {
	mux.HandleFunc("/logging.html", func(w http.ResponseWriter, r *http.Request) {
		err := r.ParseForm()
		if err != nil {
			ilog.Println("Error getting ", r.URL, ", err=", err)
			w.WriteHeader(404)
			return
		}
		form := r.Form

		bw := bufio.NewWriter(w)
		bw.WriteString(`
<!DOCTYPE html>
<html>
<head>
<title>Squareplay server logging</title>
<link type="text/css" media="screen" rel="stylesheet" href="/html/default.css">
</head>
<body>
<h1>SquarePlay server logging</h1>
<form>
<table>
`)

		for _, ls := range logSettings {
			v := form.Get(ls.name)
			if v != "" {
				if ls.setting != v {
					ls.setting = v
					ls.setter(ls)
				}
			}
			fmt.Fprintf(bw, "<tr><td>%s</td><td>%s</td></tr>\n", ls.name, ls.makeSelect())

		}
		bw.WriteString(`
</table>
<input type="submit" value="Update"/>
</form>
</body>
</html>
`)
		bw.Flush()
	})
}
