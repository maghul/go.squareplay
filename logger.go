package main

import (
	"fmt"
	"io"
	"os"
	"runtime"
	"strings"
	"sync"
	"time"

	"github.com/maghul/go.slf"
	"github.com/natefinch/lumberjack"
)

type logOutput struct {
	wr  io.Writer
	mtx sync.Mutex
}

var slog = slf.GetLogger("squareplay")

var prev time.Time

const LOGCALLER = false
const timestampFormat = "15:04:05"

func initLogging() {
	var lo = makeSquareplayLoggerOutput()
	if logfilename != "" {
		ljl := &lumberjack.Logger{
			Filename:   logfilename,
			MaxSize:    500, // megabytes
			MaxBackups: 3,
			MaxAge:     28, //days
		}
		ljl.Rotate()

		lo.wr = ljl
	} else {
		lo.wr = os.Stderr
	}

	o := slf.Output(lo)
	slog.SetOutputLogger(o)
	slf.GetLogger("raopd").SetParent(slog)
	slog.SetLevel(slf.Info)
}

func (lo *logOutput) Print(ref string, lvl slf.Level, d ...interface{}) {
	if lo == nil {
		return
	}
	msg := fmt.Sprint(d...)

	lo.WriteMessage(ref, []byte(msg))
}

func (lo *logOutput) Printf(ref string, lvl slf.Level, args string, d ...interface{}) {
	if lo == nil {
		return
	}
	msg := fmt.Sprintf(args, d...)

	lo.WriteMessage(ref, []byte(msg))
}

func (lo *logOutput) WriteMessage(ref string, p []byte) {
	if lo == nil {
		return
	}

	lo.mtx.Lock()
	defer lo.mtx.Unlock()

	w := lo.wr
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
		now.Format(timestampFormat), now.Nanosecond()/1000000, ref)
	prefix = strings.Replace(prefix, ":", "", -1)
	for _, l := range strings.Split(string(p), "\n") {
		fmt.Fprintln(w, prefix, l, errloc)
		//			fmt.Println(w, prefix, l)
	}
}

func makeSquareplayLoggerOutput() *logOutput {
	mdo := &logOutput{}
	return mdo
}

func (lo *logOutput) String() string {
	return "Squareplay Logger"
}
