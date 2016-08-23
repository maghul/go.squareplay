package main

import (
	"fmt"
	"io"
	"runtime"
	"strings"
	"time"
)

type logger struct {
	wr   io.Writer
	name string
}

var prev time.Time

const LOGCALLER = false
const timestampFormat = "15:04:05"

func makeLogger(name string, wr io.Writer) *logger {
	return &logger{wr, name}
}

func (wr *logger) Println(d ...interface{}) {
	if wr == nil {
		return
	}
	msg := fmt.Sprint(d...)

	wr.Write([]byte(msg))
}

func (wr *logger) Printf(fmts string, d ...interface{}) {
	if wr == nil {
		return
	}
	msg := fmt.Sprintf(fmts, d...)

	wr.Write([]byte(msg))
}

func (wr *logger) Write(p []byte) (n int, err error) {
	if wr == nil {
		return 0, nil
	}

	w := wr.wr
	now := time.Now()

	nd := prev.Day() != now.Day()
	prev = now
	if nd {
		fmt.Fprintln(w, "")
		fmt.Fprintln(w, "New Day", now.Format(time.RFC1123))
		fmt.Fprintln(w, "")
	}

	errloc := ""
	if LOGCALLER {
		_, file, line, ok := runtime.Caller(4)
		if ok {
			errloc = fmt.Sprintf(" in line %d of file %s. ", line, file)
		}
	}

	prefix := fmt.Sprintf("%s.%3.3d %s- ",
		now.Format(timestampFormat), now.Nanosecond()/1000000, wr.name)
	prefix = strings.Replace(prefix, ":", "", -1)
	for _, l := range strings.Split(string(p), "\n") {
		fmt.Fprintln(w, prefix, l, errloc)
		//			fmt.Println(w, prefix, l)
	}
	return len(p), nil
}
