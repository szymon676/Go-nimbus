
# Go-nimbus

Minimalistic web framework to use in your side projects 

## Usage:

#### Downloading:

```go
    go get github.com/szymon676/Go-nimbus
    go mod init 
    go mod tidy
```

#### Code:

```go
    package main

import (
	"net/http"

	// Import the Go-nimbus package
	gonimbus "github.com/szymon676/Go-nimbus"
)

func main() {
	// Create a new instance of the Go-nimbus framework
	g := gonimbus.New()

	// Use the CORS middleware to enable cross-origin resource sharing
	g.Use(gonimbus.Cors)

	// Define a route for the root path
	g.Get("/", func(w http.ResponseWriter, r *http.Request) {
		// Write a response to the client with the "hello" message
		g.String("hello", w)
	})

	// Start the server and listen for incoming requests on port 3000
	g.Serve("3000")
}
```
---
# Contributing:

## steps:
- Navigate to the Go-nimbus project at https://github.com/szymon676/Go-nimbus.
- Click Fork
- Select owner of forked repo
- Write fork name and description
- Click Create fork
- Then clone your fork and start working on your feauters
# Authors: 
## Szymon Gil - szymoslav
## Features

- Caching requests
- built in Cors
- blazingly fast http methods 
- easy to use 
- combination of many different Go frameworks so you can easly adapt to code

