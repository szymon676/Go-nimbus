package gonimbus

// benchmarks
import (
	"bytes"
	"net/http"
	"net/http/httptest"
	"testing"
)

func BenchmarkGonimbus_Get(b *testing.B) {
	g := New()
	g.Get("/", func(c Context) {
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
	g.Post("/", func(c Context) {
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
	g.Put("/", func(c Context) {
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
	g.Delete("/", func(c Context) {
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
	g.Head("/", func(c Context) {
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
	g.Patch("/", func(c Context) {
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
	rr := httptest.NewRecorder()
	c := Context{w: rr}
	for i := 0; i < b.N; i++ {
		c.String(200, "hello world")
		rr.Body.Reset()
	}
}
func BenchmarkReturn(b *testing.B) {
	rr := httptest.NewRecorder()
	c := Context{w: rr}
	for i := 0; i < b.N; i++ {
		c.Return(200, "hello", "world")
		rr.Body.Reset()
	}
}
