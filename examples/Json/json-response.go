package main

import (
	"net/http"

	gonimbus "github.com/szymon676/Go-nimbus"
)

// Define a struct to represent a user
type User struct {
	Name  string `json:"name"`
	Email string `json:"email"`
}

func main() {
	// Create a new instance of the gonimbus framework
	app := gonimbus.New()

	// Define a slice of User structs
	users := []User{
		{Name: "Alice", Email: "alice@example.com"},
		{Name: "Bob", Email: "bob@example.com"},
	}

	// Add a middleware function to handle CORS headers
	app.Use(gonimbus.Cors)

	// Define a route to return a JSON response of the users slice
	app.Get("/jsonexample", func(w http.ResponseWriter, r *http.Request) {
		// Use the JSON method to encode the users slice and write it to the response
		app.JSON(w, r, users)
	})

	// Start the server and listen on port 8080
	app.Serve("8080")
}
