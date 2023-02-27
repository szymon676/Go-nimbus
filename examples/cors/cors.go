package main

// cors middlware can be use to send requests from frontend aplication
import (
	"net/http"

	// Import the Go-nimbus package
	gonimbus "github.com/szymon676/go-nimbus"
)

func main() {
	// Create a new instance of the Go-nimbus framework
	g := gonimbus.New()

	// Use CORS middleware to enable cross-origin resource sharing
	g.Use(gonimbus.Cors)

	// Define a route for the root path
	g.Get("/", func(w http.ResponseWriter, r *http.Request) {
		// Write a response to the client with the "hello" message
		g.String(w, "hello")
	})

	// Start the server and listen for incoming requests on port 3000
	g.Serve("3000")
}
