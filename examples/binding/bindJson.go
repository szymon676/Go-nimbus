package main

import (
	"fmt"
	"net/http"

	gonimbus "github.com/szymon676/go-nimbus"
)

type Person struct {
	Name    string `json:"name"`
	Surname string `json:"surname"`
	Age     int    `json:"age"`
}

func main() {
	// initialize gonimbus instance
	g := gonimbus.New()

	// handle get request on path /
	g.Get("/", func(w http.ResponseWriter, r *http.Request) {
		var p Person
		// bind json request to variable p
		if err := g.BindJSON(r, &p); err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
		// format string
		fstring := fmt.Sprint(p.Name, p.Surname, p.Age)
		// return that string
		g.String(w, fstring)
	})

	// serve your server on port 2137
	g.Serve("2137")
}
