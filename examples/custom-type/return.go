package main

import (
	"net/http"
	// import gonimbus
	gonimbus "github.com/szymon676/Go-nimbus"
)

func main() {
	// initialize gonimbus engine
	g := gonimbus.New()
	// setup route for /
	g.Get("/", func(w http.ResponseWriter, r *http.Request) {
		// return status code
		g.Statuscode(w, 123)
		// return custom type
		g.Return(w, 123231, 1231, 31, 31, 31)
		// you can also return arrays,slices,maps,string, floats inside this function
	})
	// run your server on port 1337
	g.Serve("1337")

}
