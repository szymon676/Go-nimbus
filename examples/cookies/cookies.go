package main

import (
	"net/http"
	"time"

	gonimbus "github.com/szymon676/Go-nimbus"
)

func main() {
	// Create a new Go-nimbus engine
	engine := gonimbus.New()
	// Define a route for the root URL "/"
	engine.Get("/", func(w http.ResponseWriter, r *http.Request) {
		// Try to get the cookie with the name "my-cookie" from the request
		cookie, err := engine.GetCookie("my-cookie", r)
		if err == nil {
			// If the cookie exists, return its value
			engine.String("cookie: "+cookie.Value, w)
		} else {
			// If the cookie doesn't exist, create a new one with the name "mycookie"
			// and a value of the current time, and set it in the response
			expiration := time.Now().Add(24 * time.Hour)
			cookie := http.Cookie{Name: "mycookie", Value: time.Now().String(), Expires: expiration}
			engine.SetCookie(&cookie, w)
			engine.String("Cookie set successfully", w)
		}
	})
}
