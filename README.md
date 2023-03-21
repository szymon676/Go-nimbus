# Go-nimbus

## Go-nimbus is a minimalistic web framework that you can use in your side projects.

# Usage
## Downloading

```go
   go get github.com/szymon676/Go-nimbus
   go mod init
   go mod tidy
```

## Code

### Here's an example of how to use Go-nimbus:

```go
package main

import (
	// import go-nimbus as a gonimbus
	gonimbus "github.com/szymon676/Go-nimbus"
)

func main() {
	// initialize gonimbus engine
	g := gonimbus.New()
	
	// (optional) use cors to enable frontend fetching data
	g.Use(gonimbus.Cors)
	
	// simple get endpoint that returns "hello"
	g.Get("/", func(c gonimbus.Context) {
		c.String("hello")
	})
	
	// serve your http on port 3000
	g.Serve("3000")
}

```

# Features

## Go-nimbus includes the following features:

- Logging requests (only when you want to)
- Basic authentication
- Cookies support
- Blazingly fast HTTP methods
- Easy to use
- Combination of many different Go frameworks so you can easily adapt to code

# Contributing

## To contribute to Go-nimbus:

- Navigate to the Go-nimbus project at https://github.com/szymon676/Go-nimbus.
- Click "Fork".
- Select the owner of your forked repo.
- Write the fork name and description.
- Click "Create fork".
- Clone your fork and start working on your features.

That's it! With these simple steps, you can start using and contributing to Go-nimbus.
