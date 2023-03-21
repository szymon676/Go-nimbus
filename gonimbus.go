package gonimbus

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"

	"github.com/fatih/color"
	"github.com/gorilla/mux"
	"github.com/julienschmidt/httprouter"
)

type Gonimbus struct {
	router      *httprouter.Router
	middlewares []func(http.Handler) http.Handler
}

type HandlerFunc func(c Context)

func New() *Gonimbus {
	return &Gonimbus{
		router:      httprouter.New(),
		middlewares: []func(http.Handler) http.Handler{},
	}
}

var (
	ctx = context.Background()
)

func (g *Gonimbus) Serve(addr string) error {
	color.Cyan("-------------------------------------------------------------")
	color.Cyan("|	Server running on port http://localhost:" + addr + " 	    |")
	color.Cyan("|	Server running on port 127.0.0.1:" + addr + "		    |")
	color.Cyan("-------------------------------------------------------------")
	color.White("  		   Thanks for using Gonimbus!")
	handlerChain := g.applyMiddlewares(g.router)

	return http.ListenAndServe(":"+addr, handlerChain)
}

func (g *Gonimbus) Use(middleware func(http.Handler) http.Handler) {
	g.middlewares = append(g.middlewares, middleware)
}

func (g *Gonimbus) applyMiddlewares(handler http.Handler) http.Handler {
	for _, middleware := range g.middlewares {
		handler = middleware(handler)
	}
	return handler
}

func LogRequest(handler http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// Log the incoming request
		method := color.CyanString("%s", r.Method)
		color.Cyan("Incoming request %s on %s\n", method, r.URL.Path)
		// Call the handler function
		handler(w, r)
	}
}

func Cors(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Add the necessary headers to enable CORS
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, OPTIONS")
		w.Header().Set("Access-Control-Allow-Headers", "Content-Type, Authorization")

		// Call the next handler in the chain
		next.ServeHTTP(w, r)
	})
}

// BasicAuth is a middleware that adds basic authentication to the request.
func BasicAuth(username, password string) func(http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			user, pass, ok := r.BasicAuth()
			if !ok || user != username || pass != password {
				w.Header().Set("WWW-Authenticate", `Basic realm="Restricted"`)
				w.WriteHeader(http.StatusUnauthorized)
				w.Write([]byte("Unauthorized\n"))
				return
			}
			next.ServeHTTP(w, r)
		})
	}
}

func (g *Gonimbus) Get(path string, handle HandlerFunc) {
	g.router.GET(path, func(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
		handle(Context{w: w, r: r.WithContext(ctx)})
	})
}

// Post registers a POST request route with the provided path and handler function.
func (g *Gonimbus) Post(path string, handle HandlerFunc) {
	g.router.POST(path, func(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
		handle(Context{w: w, r: r.WithContext(ctx)})
	})
}

// Put registers a PUT request route with the provided path and handler function.
func (g *Gonimbus) Put(path string, handle HandlerFunc) {
	g.router.PUT(path, func(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
		handle(Context{w: w, r: r.WithContext(ctx)})
	})
}

// Delete registers a DELETE request route with the provided path and handler function.
func (g *Gonimbus) Delete(path string, handle HandlerFunc) {
	g.router.DELETE(path, func(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
		handle(Context{w: w, r: r.WithContext(ctx)})
	})
}

// Head registers a HEAD request route with the provided path and handler function.
func (g *Gonimbus) Head(path string, handle HandlerFunc) {
	g.router.HEAD(path, func(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
		handle(Context{w: w, r: r.WithContext(ctx)})
	})
}

// Patch registers a patch request route with the given path and handler function
func (g *Gonimbus) Patch(path string, handle HandlerFunc) {
	g.router.PATCH(path, func(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
		handle(Context{w: w, r: r.WithContext(ctx)})
	})
}

type Context struct {
	w http.ResponseWriter
	r *http.Request
}

// String writes the provided prompt as a string to the response writer.
func (c *Context) String(statuscode int, prompt string) {
	c.w.WriteHeader(statuscode)
	fmt.Fprintln(c.w, prompt)
}

// Return writes the provided values to the response writer as a string.
func (c *Context) Return(statuscode int, a ...any) {
	c.w.WriteHeader(statuscode)
	fmt.Fprint(c.w, a...)
}

// Redirect redirects the client to the provided link with the provided status code.
func (c *Context) Redirect(statuscode int, link string) {
	http.Redirect(c.w, c.r, link, statuscode)
}

func (c *Context) Param(v string) string {
	path := mux.Vars(c.r)
	result := path[v]
	return result
}

// BindJSON provides a JSON binding in request
func (c *Context) BindJSON(object interface{}) error {
	if c.r == nil || c.r.Body == nil {
		return errors.New("no request body")
	}
	defer c.r.Body.Close()
	decoder := json.NewDecoder(c.r.Body)
	if err := decoder.Decode(&object); err != nil {
		return err
	}
	return nil
}

// GetCookie retrieves the cookie with the specified name from the given request
// Returns the cookie or an error if the cookie is not found
func (c *Context) GetCookie(name string) (*http.Cookie, error) {
	cookie, err := c.r.Cookie(name)
	if err != nil {
		return nil, err
	}
	return cookie, nil
}

// SetCookie sets the given cookie in the response writer's headers
func (c *Context) SetCookie(cookie *http.Cookie) {
	http.SetCookie(c.w, cookie)
}
