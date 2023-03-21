package main

import (
	gonimbus "github.com/szymon676/go-nimbus"
)

func main() {
	g := gonimbus.New()

	// if you want auth for all routes use "use"
	g.Use(gonimbus.BasicAuth("admin", "password"))

	// define a route for "auth" path after auth you get response hello
	g.Get("/auth", func(c gonimbus.Context) {
		c.String(200, "hello")
	})

	g.Serve("4000")
}
