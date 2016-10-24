package main

import (
	"bufio"
	"fmt"
	"net/http"
)

func initUsage(mux *http.ServeMux) {
	mux.HandleFunc("/index.html", func(w http.ResponseWriter, r *http.Request) {
		bw := bufio.NewWriter(w)
		bw.WriteString(`
<!DOCTYPE html>
<html>
<head>
<title>Squareplay server</title>
<style media=\"screen\" type=\"text/css\">
body {
    background-color: #d0e4fe;
}

h1 {
    color: white;
    background-color: blue;
    font-family: \"Helvetica\";
    font-size: 38px;
}

h2 {
    color: orange;
    background-color: blue;
    font-family: \"Helvetica\";
    font-size: 30px;
}

p {
    font-family: \"Helvetica\";
    font-size: 16px;
}
</style>
</head>
<body>
<h1>SquarePlay server</h1>
<a href="/html/doc.html">Documentation on useage can be found here</a></h1>
`)

		for _, playeri := range allSqueezePlayers.snapshot() {
			player := playeri.(*SqueezePlayer)
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

	mux.HandleFunc("/doc", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprint(w, `
<html>
<body>
<h2>URLS used by the SquarePlayer server</h2>
<table>
<tr>
<td><pre>http://&lt;server&gt;/</pre></td>
<td>Show current players and some info about them</td>
</tr><tr>
<td><pre>http://&lt;server&gt;/doc</pre></td>
<td>This documentation</td>
</tr><tr>
<td><pre>http://&lt;server&gt;/notifications.json</pre></td>
<td>Get notifications from the server, such as </td>
</tr><tr>
<td><pre>http://&lt;server&gt;/control/notify</pre></td>
<td></td>
</tr><tr>
<td><pre>http://&lt;server&gt;/control/start</pre></td>
<td>Start a SquarePlay service. This becomes a new service for AirPlay devices to connect to</td>
</tr><tr>
<td><pre>http://&lt;server&gt;/control/logger/raopd.audio?&lt;off|info|debug&gt;</pre></td>
<td>Set audio logger</td>
</tr><tr>
<td><pre>http://&lt;server&gt;/control/logger/raopd.dacp?&lt;off|info|debug&gt;</pre></td>
<td>Set DACP logger for remote control of iDevice/td>
</tr><tr>
<td><pre>http://&lt;server&gt;/control/logger/raopd.dmap?&lt;off|info|debug&gt;</pre></td>
<td>Set DMAP logger, /td>
</tr><tr>
<td><pre>http://&lt;server&gt;/control/logger/raopd.auth?&lt;off|info|debug&gt;</pre></td>
<td>Log authorizations</td>
</tr><tr>
<td><pre>http://&lt;server&gt;/control/logger/raopd.net?&lt;off|info|debug&gt;</pre></td>
<td>Log low level network activity</td>
</tr><tr>
<td><pre>http://&lt;server&gt;/control/logger/raopd.raop?&lt;off|info|debug&gt;</pre></td>
<td>Log RAOP commands</td>
</tr><tr>
<td><pre>http://&lt;server&gt;/control/logger/raopd.rtp?&lt;off|info|debug&gt;</pre></td>
<td>Log RTP activity</td>
</tr><tr>
<td><pre>http://&lt;server&gt;/control/logger/raopd.rtsp?&lt;off|info|debug&gt;</pre></td>
<td>Log RTSP activity/td>
</tr><tr>
<td><pre>http://&lt;server&gt;/control/logger/raopd.sequencer?&lt;off|info|debug&gt;</pre></td>
<td>Log sequencer activity</td>
</tr><tr>
<td><pre>http://&lt;server&gt;/control/logger/raopd.volume?&lt;off|info|debug&gt;</pre></td>
<td>Log volume handling</td>
</tr><tr>
<td><pre>http://&lt;server&gt;/control/logger/raopd.zeroconf?&lt;off|info|debug&gt;</pre></td>
<td>Log ZeroConf/Bonjour activity</td>
</tr><tr>
<td><pre>http://&lt;server&gt;/control/logger/*?&lt;off|info|debug&gt;</pre></td>
<td>Set level for all loggers</td>
</tr><tr>
<td><pre>http://&lt;server&gt;/control/logger/sequencetrace?&lt;false|true&gt;</pre></td>
<td>Enable sequencetrace log, this will create a separate logfile specifically for sequencer in /tmp/*.sequencetrace.log</td>
</tr><tr>
<td><pre>http://&lt;server&gt;/control/logger/volumetrace?&lt;false|true&gt;</pre></td>
<td>Enable volumetrace log, this will create a separate logfile specifically for volume handle in /tmp/*.volumetrace.log</td>
</tr><tr>
<td><pre>http://&lt;server&gt;/&lt;client&gt;/metadata.json</pre></td>
<td>Get the metadata of the current song</td>
</tr><tr>
<td><pre>http://&lt;server&gt;/&lt;client&gt;/cover.jpg</pre></td>
<td>The coverart of the current song</td>
</tr><tr>
<td><pre>http://&lt;server&gt;/&lt;client&gt;/audio.pcm</pre></td>
<td>Get the audio as PCM</td>
</tr><tr>
<td><pre>http://&lt;server&gt;/&lt;client&gt;/audio.wav</pre></td>
<td>Get the audio as PCM</td>
</tr><tr>
<td><pre>http://&lt;server&gt;/&lt;client&gt;/control/volume=<volume></pre></td>
<td>Set the volume of the iDevice. This will not actually change the volume on the device but is used to indicate the volume of the Squeezebox</td>
</tr><tr>
<td><pre>http://&lt;server&gt;/&lt;client&gt;/control/beginff</pre></td>
<td>begin fast forward</td>
</tr><tr>
<td><pre>http://&lt;server&gt;/&lt;client&gt;/control/beginrew</pre></td>
<td>begin rewind</td>
</tr><tr>
<td><pre>http://&lt;server&gt;/&lt;client&gt;/control/mutetoggle</pre></td>
<td>toggle mute status</td>
</tr><tr>
<td><pre>http://&lt;server&gt;/&lt;client&gt;/control/nextitem</pre></td>
<td>play next item in playlist</td>
</tr><tr>
<td><pre>http://&lt;server&gt;/&lt;client&gt;/control/previtem</pre></td>
<td>play previous item in playlist</td>
</tr><tr>
<td><pre>http://&lt;server&gt;/&lt;client&gt;/control/pause</pre></td>
<td>pause playback</td>
</tr><tr>
<td><pre>http://&lt;server&gt;/&lt;client&gt;/control/playpause</pre></td>
<td>toggle between play and pause</td>
</tr><tr>
<td><pre>http://&lt;server&gt;/&lt;client&gt;/control/play</pre></td>
<td>start playback</td>
</tr><tr>
<td><pre>http://&lt;server&gt;/&lt;client&gt;/control/stop</pre></td>
<td>stop playback</td>
</tr><tr>
<td><pre>http://&lt;server&gt;/&lt;client&gt;/control/playresume</pre></td>
<td>play after fast forward or rewind</td>
</tr><tr>
<td><pre>http://&lt;server&gt;/&lt;client&gt;/control/shuffle_songs</pre></td>
<td>shuffle playlist</td>
</tr><tr>
<td><pre>http://&lt;server&gt;/&lt;client&gt;/control/volumedown</pre></td>
<td>turn audio volume down</td>
</tr><tr>
<td><pre>http://&lt;server&gt;/&lt;client&gt;/control/volumeup</pre></td>
<td>turn audio volume up</td>
</tr>
</table>
<h2>Notifications</h2>
Notifications are sent on a persustent chunked HTTP response on a http://<server>/notifications.js
request. They are sequence with a sequence number in the "chunk_extension" header in each chunk.
<h4>There are three notification messages.</h4>
<ul>
<li>dmap.listingitem</li> A json hash literal containig metadata about the current track
<li>volume</li> A volume change from the iDevice. The value may be an absolute between 0..100 or +2 and -2
for a relative volume.
<li>progress</li>A JSON hash literal containing 
{
    current=<time>;
    length=<time>;
}
Where time is in milliseconds.

</body>
</html>
`)
	})
}
