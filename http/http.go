// This package provides utilities to
// serve the frontend / API of touchy.
package http

import (
	"encoding/json"
	"github.com/gorilla/mux"
	"net/http"
	"strings"

	"github.com/odino/touchy/keyboard"
)

// A very barebone response to
// send out to clients calling
// the api.
type response struct {
	Keys   string `json:"keys"`
	Status string `json:"status"`
}

// Handler used to simulate keybord
// presses.
// Keys should be a "+" separated list
// of keys you'd like to press, like:
// - "a"
// - "ctrl+shift+s"
func PressHandler(w http.ResponseWriter, r *http.Request) {
	k := mux.Vars(r)["keys"]
	keyboard.Press(strings.Split(k, "+"))

	w.Header().Add("Content-Type", "application/json")
	j, _ := json.Marshal(response{k, "ok"})
	w.Write(j)
}

// Handler to serve any static content under /static:
// - /static/hello.txt --> static/hello.txt
func StaticHandler(w http.ResponseWriter, r *http.Request) {
	asset := mux.Vars(r)["asset"]
	data, err := Asset("static/" + asset)

	if err != nil {
		panic(err)
	}

	w.Write(data)
}

// Serve the homepage, loading index.html that
// will boot the frontend.
func HomeHandler(w http.ResponseWriter, r *http.Request) {
	data, err := Asset("static/index.html")

	if err != nil {
		panic(err)
	}

	w.Write(data)
}

// Custom 404 handler
func NotFoundHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("Ante", "up") // LOL can't  get it out of my head (https://www.youtube.com/watch?v=wUnfFC7wUrs)
	w.WriteHeader(404)
}
