// Package events provides a webservice that manages the library's special events.
package events

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/tislais/go-packages/services/internal/ports"
)

var port = 42

// StartServer registers the handlers and initiates the web service.
// The service is started on the local machine with the port specified by
// .../lm/services/internal/ports#EventService
func StartServer() error {
	sm := http.NewServeMux()
	sm.Handle("/", new(eventHandler))
	return http.ListenAndServe(":"+strconv.Itoa(port), sm)
}

func init() {
	fmt.Println("server.go 1", port)
	port = ports.EventService
	fmt.Println("server.go 2", port)
}

func init() {
	fmt.Println("second init exists in serve.go")
}
