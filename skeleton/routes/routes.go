package routes

import (
	"net/http"

	"github.com/anonx/sunplate/skeleton/assets/handlers"

	r "github.com/anonx/sunplate/routing"
)

// The line below tells golang's generate command you want
// it to scan your controllers and generate handler functions
// from them using rules of sunplate toolkit.
// Please, do not delete it unless you know what you are doing.
//
//go:generate sunplate generate handlers --input ../controllers --output ../assets/handlers

// List is a slice of routes of the following form:
//	Route:
//		Pattern
//		Handlers:
//			Method: Handler
// If using a standard router just call Context.Build() to get http handler
// as the first argument and an error (or nil) as the second one.
var List = r.Routes{
	r.Get("/", handlers.App.Index),
	r.Get("/greet/:name", handlers.App.PostGreet),

	// Serve static files of ./static directory.
	r.Get("/static", http.FileServer(http.Dir("./static")).ServeHTTP),
}