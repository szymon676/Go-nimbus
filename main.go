package main

import (
	"net/http"

	"github.com/julienschmidt/httprouter"
)

type Gonimbus struct {
	router *httprouter.Router
}

func New() *Gonimbus {
	// New creates a new instance of Gonimbus
	// It initializes a new httprouter.Router and returns a new Gonimbus object that wraps the router
	return &Gonimbus{router: httprouter.New()}
}

func (g *Gonimbus) Serve(addr string) error {
	// Serve listens on the TCP network address addr and then serves incoming HTTP requests using g.router
	return http.ListenAndServe(":"+addr, g.router)
}

func (g *Gonimbus) Get(statusCode int, path string, handle http.HandlerFunc) {
	// Get adds a new route for HTTP GET method to the router with the specified path and handle
	// It calls the handle function and writes the HTTP status code before writing the response body
	g.router.GET(path, func(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
		w.WriteHeader(statusCode)
		handle(w, r)
	})
}

func (g *Gonimbus) Post(statuscode int, path string, handle http.HandlerFunc) {
	// Post adds a new route for HTTP POST method to the router with the specified path and handle
	// It calls the handle function and writes the HTTP status code before writing the response body
	g.router.POST(path, func(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
		w.WriteHeader(statuscode)
		handle(w, r)
	})
}

func (g *Gonimbus) Put(statuscode int, path string, handle http.HandlerFunc) {
	// Put adds a new route for HTTP PUT method to the router with the specified path and handle
	// It calls the handle function and writes the HTTP status code before writing the response body
	g.router.PUT(path, func(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
		w.WriteHeader(statuscode)
		handle(w, r)
	})
}

func (g *Gonimbus) Delete(statuscode int, path string, handle http.HandlerFunc) {
	// Delete adds a new route for HTTP DELETE method to the router with the specified path and handle
	// It calls the handle function and writes the HTTP status code before writing the response body
	g.router.DELETE(path, func(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
		w.WriteHeader(statuscode)
		handle(w, r)
	})
}

func (g *Gonimbus) Head(statuscode int, path string, handle http.HandlerFunc) {
	// Head adds a new route for HTTP HEAD method to the router with the specified path and handle
	// It calls the handle function and writes the HTTP status code before writing the response body
	g.router.HEAD(path, func(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
		w.WriteHeader(statuscode)
		handle(w, r)
	})
}

func (g *Gonimbus) String(prompt string, w http.ResponseWriter) {
	// String writes the prompt to the response body using the http.ResponseWriter
	w.Write([]byte(prompt))
}

func main() {
	// New creates a new instance of Gonimbus
	g := New()

	// Serve listens on the TCP network address "3000" and serves incoming HTTP requests using g.router
	err := g.Serve("3000")
	if err != nil {
		panic(err)
	}
}
