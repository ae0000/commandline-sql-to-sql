package main

import (
	"fmt"
	"net/http"

	"github.com/gocraft/web"
	"github.com/unrolled/render"
)

// Context for holding data between the request and the template
type Context struct {
	RawSQL       string
	ConvertedSQL string
}

var rend = render.New(render.Options{
	Extensions: []string{".html"},
	Directory:  "templates",
	Layout:     "index",
})

func main() {
	// Setup router
	rootRouter := web.New(Context{})

	rootRouter.Middleware(web.LoggerMiddleware)
	rootRouter.Middleware(web.StaticMiddleware("assets"))

	// Routes
	rootRouter.Get("/", (*Context).convertSQL)
	rootRouter.Post("/", (*Context).convertSQL)

	// Serve
	port := "3333"
	fmt.Printf("\x1b[32;1m --------- ConvertSQL [listening on port %s]\x1b[0m", port)
	http.ListenAndServe("localhost:"+port, rootRouter)
}
