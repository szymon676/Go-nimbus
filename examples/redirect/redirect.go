package main

import (
	"net/http"

	gonimbus "github.com/szymon676/Go-nimbus"
)

func main() {
	// Create a new instance of the Gonimbus server
	g := gonimbus.New()

	// Define a GET route for the root path "/"
	g.Get("/", func(w http.ResponseWriter, r *http.Request) {
		// Redirect the client to the URL "http://localhost:3000" with status code 301 (Moved Permanently)
		g.Redirect(w, r, "http://localhost:3000", http.StatusMovedPermanently)
	})

	// Start the server on port 1000
	g.Serve("1000")
}
