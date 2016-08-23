package main

import (
	"io"
	"os"
)

func getKeyfile() io.Reader {
	f, err := os.Open("/tmp/airport.key")
	if err != nil {
		panic("No key found...")
	}
	return f
}
