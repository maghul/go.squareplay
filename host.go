package main

import (
	"fmt"
	"sync"
)

type host struct {
	ref       string
	ncs       []chan []byte
	ncs_mutex sync.Mutex
}

var allHosts = newSyncMap()

func getHost(ref string) *host {
	if ref == "" {
		panic("Bad channel refererence, reference is empty")
	}
	f := func() (interface{}, error) {
		return &host{ref, make([]chan []byte, 0), sync.Mutex{}}, nil
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

func (h *host) addNotificationChannel(nc chan []byte) {
	h.ncs_mutex.Lock()
	defer h.ncs_mutex.Unlock()

	h.ncs = append(h.ncs, nc)
}

func (h *host) removeNotificationChannel(nc chan []byte) {
	h.ncs_mutex.Lock()
	defer h.ncs_mutex.Unlock()

	for ii, c := range h.ncs {
		if c == nc {
			h.ncs = append(h.ncs[0:ii], h.ncs[ii+1:]...)
			return
		}
	}
}

func (h *host) txNotification(b []byte) {
	h.ncs_mutex.Lock()
	defer h.ncs_mutex.Unlock()

	for _, c := range h.ncs {
		select {
		case c <- b:
		default:
			// TODO: Maybe we should drop the notification channel if doesn't receive
		}
	}
}
