# chi-kit

`chi-kit` is a helper library for the `go-chi/chi` router. It provides a set of tools and wrappers to simplify the setup of a web server, offering an experience similar to the Gin framework.

## Features

*   **Engine Initialization**: Quickly set up a Chi router with common middleware (Logger, Recoverer, RealIP, RequestID).
*   **Simplified Routing**: Wrapper functions for `GET`, `POST`, `PUT`, `DELETE`, etc.
*   **Standardized JSON Responses**: Helper functions like `SendSuccess` and `SendError` for consistent JSON output.
*   **Static File Serving**: Easily serve static files from a directory or embedded file system.
*   **Graceful Shutdown**: Integrated server startup and graceful shutdown logic.

## Installation

```shell
go get github.com/openmindw/chi-kit
```

## Usage

Here is a basic example of how to use `chi-kit`:

```go
package main

import (
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/openmindw/chi-kit"
)

func main() {
	// 1. Initialize the engine (true for debug mode)
	chikit.Engine(true)

	// 2. Basic route
	chikit.Get("/", func(w http.ResponseWriter, r *http.Request) {
		chikit.SendSuccess(w, map[string]string{
			"message": "Hello, chi-kit!",
		})
	})

	// 3. Route group
	chikit.Route("/api/v1", func(r chi.Router) {
		r.Get("/users/{id}", func(w http.ResponseWriter, r *http.Request) {
			id := chikit.URLParam(r, "id")
			chikit.SendSuccess(w, map[string]string{
				"id":   id,
				"name": "User " + id,
			})
		})
	})

	// 4. Start the server
	chikit.Server(":8080")
}
```