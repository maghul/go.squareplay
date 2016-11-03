package main

import (
	"bufio"
	"bytes"
	"fmt"
	"io"
	"net/http"
	"strings"

	"github.com/maghul/go.raopd"
	"github.com/maghul/go.slf"
)

type logAdmin struct {
	out io.Writer
	wh  http.Handler
}

func LogAdminHandler(out io.Writer, wh http.Handler) http.Handler {
	lh := &logAdmin{out, wh}
	return lh
}

func (lh *logAdmin) ServeHTTP(resp http.ResponseWriter, req *http.Request) {
	fmt.Fprintln(lh.out, "ServeHTTP: url=", req.URL, ", client=", req.RemoteAddr)
	lh.wh.ServeHTTP(resp, req)
}

type logSetting struct {
	name        string
	description string
	setting     string
	settings    string
	setter      func(ls *logSetting)
}

var logSettings = []*logSetting{
	&logSetting{"sequencetrace", "Trace Packet sequencing", "off", "off|on", setRaopTraceLogging},
	&logSetting{"volumetrace", "Trace Volume handling", "off", "off|on", setRaopTraceLogging},
}

func setRaopLogging(ls *logSetting) {
	err := slf.SetLevel(ls.name, ls.setting)
	if err != nil {
		slog.Info.Println("Could not set logging: ", err)
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
			slog.Info.Println("initLogHandler: Error getting ", r.URL, ", err=", err)
			w.WriteHeader(404)
			fmt.Fprintln(w, "initLogHandler: Error getting ", r.URL, ", err=", err)
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
<h2>Loggers</h2>
<table>
`)

		for _, logger := range slf.Loggers() {
			name := logger.Name()
			v := form.Get(name)
			if v != "" {
				lvl, err := slf.FindLevel(v)
				if err != nil {
					logger.SetLevel(lvl)
				}
			}
			fmt.Fprintf(bw, "<tr><td>%s</td><td>%s</td><td>%s</td></tr>\n",
				makeSelect(name, logger.Level().String(), "off|info|debug"), name, logger.Description())

		}

		bw.WriteString(`
</table>
<h2>Tracers</h2>
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
			fmt.Fprintf(bw, "<tr><td>%s</td><td>%s</td><td>%s</td></tr>\n", ls.makeSelect(), ls.name, ls.description)

		}
		bw.WriteString(`
</table>
<h2>Update</h2>
<input type="submit" value="Update"/>
</form>
</body>
</html>
`)
		bw.Flush()
	})
}
