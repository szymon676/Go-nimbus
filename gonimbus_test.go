package gonimbus

// benchmarks
import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"
)

func BenchmarkGonimbus_Get(b *testing.B) {
	g := New()
	g.Get("/", func(w http.ResponseWriter, r *http.Request) {
		// Do nothing
	})

	req, err := http.NewRequest("GET", "/", nil)
	if err != nil {
		b.Fatal(err)
	}
	rr := httptest.NewRecorder()

	for i := 0; i < b.N; i++ {
		g.router.ServeHTTP(rr, req)
		rr.Body.Reset()
	}
}

func BenchmarkGonimbus_Post(b *testing.B) {
	g := New()
	g.Post("/", func(w http.ResponseWriter, r *http.Request) {
		// Do nothing
	})

	req, err := http.NewRequest("POST", "/", bytes.NewReader([]byte{}))
	if err != nil {
		b.Fatal(err)
	}
	rr := httptest.NewRecorder()

	for i := 0; i < b.N; i++ {
		g.router.ServeHTTP(rr, req)
		rr.Body.Reset()
	}
}

func BenchmarkGonimbus_Put(b *testing.B) {
	g := New()
	g.Put("/", func(w http.ResponseWriter, r *http.Request) {
		// Do nothing
	})

	req, err := http.NewRequest("PUT", "/", bytes.NewReader([]byte{}))
	if err != nil {
		b.Fatal(err)
	}
	rr := httptest.NewRecorder()

	for i := 0; i < b.N; i++ {
		g.router.ServeHTTP(rr, req)
		rr.Body.Reset()
	}
}

func BenchmarkGonimbus_Delete(b *testing.B) {
	g := New()
	g.Delete("/", func(w http.ResponseWriter, r *http.Request) {
		// Do nothing
	})

	req, err := http.NewRequest("DELETE", "/", nil)
	if err != nil {
		b.Fatal(err)
	}
	rr := httptest.NewRecorder()

	for i := 0; i < b.N; i++ {
		g.router.ServeHTTP(rr, req)
		rr.Body.Reset()
	}
}

func BenchmarkGonimbus_Head(b *testing.B) {
	g := New()
	g.Head("/", func(w http.ResponseWriter, r *http.Request) {
		// Do nothing
	})

	req, err := http.NewRequest("HEAD", "/", nil)
	if err != nil {
		b.Fatal(err)
	}
	rr := httptest.NewRecorder()

	for i := 0; i < b.N; i++ {
		g.router.ServeHTTP(rr, req)
		rr.Body.Reset()
	}
}

func BenchmarkGonimbus_PATCH(b *testing.B) {
	g := New()
	g.Patch("/", func(w http.ResponseWriter, r *http.Request) {
		// Do nothing
	})

	req, err := http.NewRequest("PATCH", "/", nil)
	if err != nil {
		b.Fatal(err)
	}
	rr := httptest.NewRecorder()

	for i := 0; i < b.N; i++ {
		g.router.ServeHTTP(rr, req)
		rr.Body.Reset()
	}
}

func BenchmarkGonimbus_String(b *testing.B) {
	g := New()
	rr := httptest.NewRecorder()

	for i := 0; i < b.N; i++ {
		g.String(rr, "hello world")
		rr.Body.Reset()
	}
}
func BenchmarkInt(b *testing.B) {
	g := &Gonimbus{}
	w := httptest.NewRecorder()

	for i := 0; i < b.N; i++ {
		g.Int(w, 12345)
	}
}
func BenchmarkReturn(b *testing.B) {
	g := &Gonimbus{}
	w := httptest.NewRecorder()

	for i := 0; i < b.N; i++ {
		g.Return(w, "hello", "world")
	}
}

func BenchmarkJSON(b *testing.B) {
	// Create a test HTTP response recorder
	w := httptest.NewRecorder()
	// Create a dummy request for the JSON function
	r, _ := http.NewRequest(http.MethodGet, "/", nil)
	// Create a dummy data object for the JSON function
	data := map[string]interface{}{
		"foo": "bar",
		"baz": 123,
	}
	// Create a Gonimbus instance
	g := &Gonimbus{}

	// Run the JSON function b.N times and record the execution time
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		// Reset the response recorder before each iteration
		w = httptest.NewRecorder()
		// Call the JSON function with the dummy request, response writer, and data
		g.JSON(w, r, data)
	}
	b.StopTimer()

	// Check the encoded JSON output for correctness (optional)
	var buf bytes.Buffer
	json.NewEncoder(&buf).Encode(data)
	if w.Body.String() != buf.String() {
		b.Errorf("encoded JSON output doesn't match expected output")
	}
}

func BenchmarkCache(b *testing.B) {
	g := New()

	// Create a CacheOptions struct with a TTL of 10 minutes
	cacheOpts := &CacheOptions{
		MaxAge:         time.Duration(10) * time.Minute,
		Public:         true,
		MustRevalidate: true,
	}

	// Use caching middleware with the CacheOptions struct
	g.Use(Cache(cacheOpts))

	// Define a route handler that will be cached
	g.Get("/", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "text/plain")
		w.Header().Set("Cache-Control", "public, max-age=600, must-revalidate")
		w.Header().Set("Expires", time.Now().Add(time.Duration(10)*time.Minute).UTC().Format(http.TimeFormat))
		w.Write([]byte("Hello, World!"))
	})

	// Create a request object for the cached route
	req, _ := http.NewRequest("GET", "/", nil)

	// Create a response recorder
	rr := httptest.NewRecorder()

	// Make the cached request repeatedly
	for i := 0; i < b.N; i++ {
		g.router.ServeHTTP(rr, req)
		rr.Body.Reset()
	}
}
