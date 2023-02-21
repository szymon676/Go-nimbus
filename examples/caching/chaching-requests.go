package main

// caching requests can help you improve performance
import (
	"net/http"
	"time"

	gonimbus "github.com/szymon676/Go-nimbus"
)

func main() {
	g := gonimbus.New()

	// Add caching middleware to all requests.
	cacheOpts := &gonimbus.CacheOptions{
		MaxAge:         time.Duration(10) * time.Minute,
		Public:         true,
		MustRevalidate: true,
	}

	g.Use(gonimbus.Cache(cacheOpts))
	g.Get("/", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "text/plain")
		w.Header().Set("Cache-Control", "public, max-age=600, must-revalidate")
		w.Header().Set("Expires", time.Now().Add(time.Duration(10)*time.Minute).UTC().Format(http.TimeFormat))
		w.Write([]byte("Hello, World!"))
	})

	// Start the server.
	g.Serve("3000")
}
