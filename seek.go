package main

import (
	"net/http"
)

func (sp *SqueezePlayer) seek(w http.ResponseWriter, r *http.Request) {
	// TODO Implement a seek function. URL is http://.../<player>/time/<ms>
}
