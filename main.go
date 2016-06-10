// An HTTP server to control
// keyboard inputs / presses
// with an HTTP API.
package main

import (
	"fmt"
	"github.com/gorilla/mux"
	"net"
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

// Get our local IP address:
// so that the remote controller
// can connect to this address + port
// to access us
func getLocalIP() string {
	addrs, err := net.InterfaceAddrs()
	if err != nil {
		return ""
	}
	for _, address := range addrs {
		if ipnet, ok := address.(*net.IPNet); ok && !ipnet.IP.IsLoopback() {
			if ipnet.IP.To4() != nil {
				return ipnet.IP.String()
			}
		}
	}

	return ""
}

// Simple helper function to print the
// address from which you can RC us
func welcome(ip string, port string) {
	if ip == "" {
		fmt.Print("Unable to get local IP address. Are we connected to a network?")
	} else {
		fmt.Printf("Aye, here we are: connect to http://%s:%s and enjoy!", ip, port)
	}
}

// Have fun!
func main() {
	localIp := getLocalIP()
	port := getPort()

	welcome(localIp, port)

	r := mux.NewRouter()
	r.HandleFunc("/press/{keys}", h.PressHandler)

	http.Handle("/", r)
	http.ListenAndServe(":"+port, nil)
}
