package gonimbus

import (
	"bytes"
	"context"
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

var (
	ctx = context.Background()
)

func New() *Gonimbus {
	return &Gonimbus{
		router:      httprouter.New(),
		middlewares: []func(http.Handler) http.Handler{},
	}
}

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

func (g *Gonimbus) Get(path string, handle http.HandlerFunc) {
	g.router.GET(path, func(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
		handle(w, r.WithContext(ctx))
	})
}

func (g *Gonimbus) Post(path string, handle http.HandlerFunc) {
	g.router.POST(path, func(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
		handle(w, r.WithContext(ctx))
	})
}

func (g *Gonimbus) Put(path string, handle http.HandlerFunc) {
	g.router.PUT(path, func(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
		handle(w, r.WithContext(ctx))
	})
}

func (g *Gonimbus) Delete(path string, handle http.HandlerFunc) {
	g.router.DELETE(path, func(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
		handle(w, r.WithContext(ctx))
	})
}

func (g *Gonimbus) Head(path string, handle http.HandlerFunc) {
	g.router.HEAD(path, func(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
		handle(w, r.WithContext(ctx))
	})
}

func (g *Gonimbus) String(prompt string, w http.ResponseWriter) {
	w.Write([]byte(prompt))
}

func (g *Gonimbus) Redirect(link string, statuscode int, w http.ResponseWriter, r *http.Request) {
	http.Redirect(w, r, link, statuscode)
}

func (g *Gonimbus) Statuscode(statuscode int, w http.ResponseWriter) {
	w.WriteHeader(statuscode)
}
