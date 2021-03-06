// Package main is an entry point of the application.
package main

import (
	"flag"
	"log"
	"net/http"

	"github.com/colegion/goal/internal/skeleton/assets/handlers"

	"github.com/colegion/contrib/routers/denco"
	"github.com/colegion/contrib/servers/grace"
	"github.com/goaltools/iniflag"
)

var addr = flag.String("http.addr", ":9000", "HTTP address the app must listen on")

// main parses configuration files making the values available
// to flags of the app, builds a handler using routes and handlers
// of the automatically generated package, and starts a new server.
func main() {
	// Parse configuration files and flags.
	err := iniflag.Parse("config/app.ini")
	assertNil(err)

	// Initialize and build routes.
	h, err := denco.Build(handlers.Init())
	assertNil(err)

	// Allocate and run a new HTTP server.
	s := &http.Server{
		Addr:    *addr,
		Handler: h,
	}
	err = grace.Serve(s)
	assertNil(err)
}

// assertNil gets an error as an argument
// and makes sure it is nil.
// If not, it terminates the program.
func assertNil(err error) {
	if err != nil {
		log.Fatal(err)
	}
}
