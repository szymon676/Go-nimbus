package main

import (
	"net/http"

	gonimbus "github.com/szymon676/Go-nimbus"
)

func main() {
	g := gonimbus.New()

	// if you want auth for all routes use "use"
	g.Use(gonimbus.BasicAuth("admin", "password"))

	// define a route for "auth" path after auth you get response hello
	g.Get("/auth", func(w http.ResponseWriter, r *http.Request) {
		g.String(w, "hello")
	})

	g.Serve("4000")
}
