package main

import (
	"net/http"

	gonimbus "github.com/szymon676/Go-nimbus"
)

func main() {
	// Create a new Go-nimbus engine
	engine := gonimbus.New()

	// Define a route for the root URL "/"
	engine.Get("/req", gonimbus.LogRequest(func(w http.ResponseWriter, r *http.Request) {
		// enjoy loggin request middleware :D
	}))

	//start gonimbus engine on port 3000
	engine.Serve("3000")
}
