// An HTTP server to control
// keyboard inputs / presses
// with an HTTP API.
package main

import (
	"github.com/gorilla/mux"
	"net/http"
	"os"

	h "github.com/odino/touchy/http"
)

// Returns the port used by the server
// By default 8080 will be used, but
// you can override this with the
// HTTP_PORT env variable.
func getPort() string {
	port := "8080"
	os_port := os.Getenv("HTTP_PORT")

	if os_port != "" {
		port = os_port
	}

	return port
}

// Have fun!
func main() {
	r := mux.NewRouter()
	r.HandleFunc("/press/{keys}", h.PressHandler)

	http.Handle("/", r)
	http.ListenAndServe(":"+getPort(), nil)
}
