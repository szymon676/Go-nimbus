package gonimbus

import (
	"bytes"
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"time"

	"github.com/fatih/color"
	"github.com/julienschmidt/httprouter"
)

type Gonimbus struct {
	router      *httprouter.Router
	middlewares []func(http.Handler) http.Handler
}

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
	color.Cyan("Server running on port http://localhost:" + addr)
	color.Red("Thanks for using Gonimbus")
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
		color.White("Incoming request %s on %s\n", method, r.URL.Path)
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

// CacheOptions defines the options for the caching middleware.
type CacheOptions struct {
	MaxAge         time.Duration // The maximum age of the cache.
	Public         bool          // Whether the cache is public or private.
	MustRevalidate bool          // Whether the cache must revalidate after expiration.
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

// Cache is a middleware that adds caching headers to the response.
func Cache(opts *CacheOptions) func(http.Handler) http.Handler {
	expireTime := time.Time{}
	if opts.MaxAge > 0 {
		expireTime = time.Now().Add(opts.MaxAge).UTC()
	}

	cacheControlBuf := bytes.NewBuffer(make([]byte, 0, 32))
	cacheControlBuf.WriteString("max-age=")
	fmt.Fprintf(cacheControlBuf, "%d", int64(opts.MaxAge.Seconds()))
	if opts.Public {
		cacheControlBuf.WriteString(", public")
	} else {
		cacheControlBuf.WriteString(", private")
	}
	if opts.MustRevalidate {
		cacheControlBuf.WriteString(", must-revalidate")
	}
	cacheControl := cacheControlBuf.String()

	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			// Set the cache headers.
			if opts.MaxAge > 0 {
				w.Header().Set("Cache-Control", cacheControl)
				w.Header().Set("Expires", expireTime.Format(http.TimeFormat))
			}

			// Call the next handler in the chain.
			next.ServeHTTP(w, r)
		})
	}
}

// Get registers a GET request route with the provided path and handler function.
func (g *Gonimbus) Get(path string, handle http.HandlerFunc) {
	// Define the route using the httprouter library.
	g.router.GET(path, func(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
		// Call the provided handler function, passing in the response writer and the request with a context.
		handle(w, r.WithContext(ctx))
	})
}

// Post registers a POST request route with the provided path and handler function.
func (g *Gonimbus) Post(path string, handle http.HandlerFunc) {
	// Define the route using the httprouter library.
	g.router.POST(path, func(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
		// Call the provided handler function, passing in the response writer and the request with a context.
		handle(w, r.WithContext(ctx))
	})
}

// Put registers a PUT request route with the provided path and handler function.
func (g *Gonimbus) Put(path string, handle http.HandlerFunc) {
	// Define the route using the httprouter library.
	g.router.PUT(path, func(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
		// Call the provided handler function, passing in the response writer and the request with a context.
		handle(w, r.WithContext(ctx))
	})
}

// Delete registers a DELETE request route with the provided path and handler function.
func (g *Gonimbus) Delete(path string, handle http.HandlerFunc) {
	// Define the route using the httprouter library.
	g.router.DELETE(path, func(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
		// Call the provided handler function, passing in the response writer and the request with a context.
		handle(w, r.WithContext(ctx))
	})
}

// Head registers a HEAD request route with the provided path and handler function.
func (g *Gonimbus) Head(path string, handle http.HandlerFunc) {
	// Define the route using the httprouter library.
	g.router.HEAD(path, func(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
		// Call the provided handler function, passing in the response writer and the request with a context.
		handle(w, r.WithContext(ctx))
	})
}

// Patch registers a patch request route with the given path and handler function
func (g *Gonimbus) Patch(path string, handle http.HandlerFunc) {
	// Define the route using the httprouter library.
	g.router.PATCH(path, func(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
		// Call the provided handler function, passing in the response writer and the request with a context.
		handle(w, r.WithContext(ctx))
	})
}

func (g *Gonimbus) Int(w http.ResponseWriter, prompt int) {
	fmt.Fprint(w, prompt)
}

// String writes the provided prompt as a string to the response writer.
func (g *Gonimbus) String(w http.ResponseWriter, prompt string) {
	fmt.Fprint(w, prompt)
}

// Redirect redirects the client to the provided link with the provided status code.
func (g *Gonimbus) Redirect(w http.ResponseWriter, r *http.Request, link string, statuscode int) {
	http.Redirect(w, r, link, statuscode)
}

// Return writes the provided values to the response writer as a string.
func (g *Gonimbus) Return(w http.ResponseWriter, a ...interface{}) {
	fmt.Fprint(w, a...)
}

// Statuscode sets the provided status code to the response writer.
func (g *Gonimbus) Statuscode(w http.ResponseWriter, statuscode int) {
	// write status code
	w.WriteHeader(statuscode)
}

// JSON writes the provided data to the response writer as a JSON object.
func (g *Gonimbus) JSON(w http.ResponseWriter, r *http.Request, data interface{}) {
	// Set the content type of the response writer to JSON.
	w.Header().Set("Content-Type", "application/json")
	// Encode the provided data as a JSON object and write it to the response writer.
	json.NewEncoder(w).Encode(data)
}

// BindJSON provides a JSON binding in request
func (g *Gonimbus) BindJSON(r *http.Request, object interface{}) error {
	// check if request has data inside
	if r == nil || r.Body == nil {
		return errors.New("no request body")
	}
	// after finish binding close reader
	defer r.Body.Close()
	// decode response body
	decoder := json.NewDecoder(r.Body)
	// if err is not nil return error
	if err := decoder.Decode(&object); err != nil {
		return err
	}
	// if everything is ok return nil
	return nil
}

// GetCookie retrieves the cookie with the specified name from the given request
// Returns the cookie or an error if the cookie is not found
func (g *Gonimbus) GetCookie(r *http.Request, name string) (*http.Cookie, error) {
	// Use the Request's Cookie method to retrieve the cookie with the specified name
	cookie, err := r.Cookie(name)
	// If the cookie is not found, return nil and the error
	if err != nil {
		return nil, err
	}
	// If the cookie is found, return the cookie and nil error
	return cookie, nil
}

// SetCookie sets the given cookie in the response writer's headers
func (g *Gonimbus) SetCookie(w http.ResponseWriter, cookie *http.Cookie) {
	// Use the ResponseWriter's SetCookie method to set the cookie in the headers
	http.SetCookie(w, cookie)
}
