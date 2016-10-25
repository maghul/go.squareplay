package main

//go:generate go-bindata -o html.go html/

import (
	"bufio"
	"fmt"
	"net/http"
	"sort"
	"strings"
)

func initUsage(mux *http.ServeMux) {
	mux.HandleFunc("/index.html", func(w http.ResponseWriter, r *http.Request) {
		bw := bufio.NewWriter(w)
		bw.WriteString(`
<!DOCTYPE html>
<html>
<head>
<title>Squareplay server</title>
<link type="text/css" media="screen" rel="stylesheet" href="/html/default.css">
</head>
<body>
<h1>SquarePlay server</h1>
<a href="/html/doc.html">Documentation on useage can be found here</a></h1>
`)

		for _, player := range makePlayerarray() {
			mac := player.Id()
			fmt.Fprintf(bw, "<h2>%s : %s</h2>\n", player.Name(), mac)
			fmt.Fprintf(bw, "<p>  AUDIO:<a href=\"%s/audio.pcm\">%s/audio.pcm</a></p>\n", mac, mac)
			fmt.Fprintf(bw, "<p>  COVER:<a href=\"%s/cover.jpg\">%s/cover.jpg</a></p>\n", mac, mac)
			//			fmt.Fprintf( bw, "<p>  RAOP: '%p'", cls->raop )
			//			if (cls->raop) {
			//				fmt.Fprintf( bw, "<p>     DACP: '%s'", dacp_state(cls->raop) );
			//			}
			//			fmt.Fprintf( bw, "<p>  Volume controls is %s</p>\n", cls->absolute_volume?"absolute":"relative" );
		}

		bw.WriteString(`
</body>
</html>
`)
		bw.Flush()
	})

	mux.HandleFunc("/html/", func(w http.ResponseWriter, r *http.Request) {
		url := r.URL.String()[1:]
		ilog.Println("HANDLING html URL '", url, "'")
		data, err := Asset(url)
		if err != nil {
			w.WriteHeader(404)
			return
		}
		w.Header().Add("Content-Type", getMimeType(url))
		_, err = w.Write(data)
		if err != nil {
			w.WriteHeader(404)
			return
		}
	})
}

func getMimeType(filename string) string {
	s := strings.Split(filename, ".")
	suffix := s[len(s)-1]
	switch suffix {
	case "html":
		return "text/html"
	case "css":
		return "text/css"
	default:
		return "text/ascii"
	}
}

type playerarray [](*SqueezePlayer)

func makePlayerarray() playerarray {
	players := allSqueezePlayers.snapshot()

	pa := make([](*SqueezePlayer), len(players))

	ii := 0
	for _, playeri := range players {
		player := playeri.(*SqueezePlayer)
		pa[ii] = player
		ii++
	}
	sort.Sort(playerarray(pa))
	return playerarray(pa)
}

func (pa playerarray) Len() int {
	return len(pa)
}

func (pa playerarray) Less(i, j int) bool {
	return strings.ToLower(pa[i].Name()) < strings.ToLower(pa[j].Name())
}

func (pa playerarray) Swap(i, j int) {
	sp := pa[i]
	pa[i] = pa[j]
	pa[j] = sp
}
