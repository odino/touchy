package touch

import (
	"encoding/json"
	"github.com/gorilla/mux"
	"net/http"
	"strings"

	"github.com/odino/touchy/keyboard"
)

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
