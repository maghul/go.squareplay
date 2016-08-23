package main

import (
	"fmt"
)

type host struct {
	ref                 string
	notificationChannel chan []byte
}

var allHosts = newSyncMap()

func getHost(ref string) *host {
	if ref == "" {
		panic("Bad channel refererence, reference is empty")
	}
	f := func() (interface{}, error) {
		return &host{ref, make(chan []byte, 32)}, nil
	}
	c, err := allHosts.Get(ref, f)
	if err != nil {
		panic(err)
	}
	return c.(*host)
}

func (h *host) ToString() string {
	return fmt.Sprintf("host:%s", h.ref)
}
