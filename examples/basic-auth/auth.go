package main

import (
	"net/http"

	gonimbus "github.com/szymon676/Go-nimbus"
)

func main() {
	g := gonimbus.New()
	// if you wan to add this for all routes use "use"
	g.Use(gonimbus.BasicAuth("admin", "password"))

	g.Get("/auth", func(w http.ResponseWriter, r *http.Request) {
		g.String(w, "hello")
	})

	g.Serve("4000")
}
