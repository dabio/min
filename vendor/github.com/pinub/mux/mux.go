// Package mux is a high performance HTTP request router, also called
// multiplexer or just mux.
//
// Example:
//
//  package main
//
//  import (
//  	"fmt"
//  	"log"
//  	"net/http"
//
//  	"github.com/pinub/mux"
//  )
//
//  func index(w http.ResponseWriter, r *http.Request) {
//  	fmt.Fprint(w, "Welcome!\n")
//  }
//
//  func hello(w http.ResponseWriter, r *http.Request) {
//  	fmt.Fprint(w, "Hello\n")
//  }
//
//  func main() {
//  	m := mux.New()
//  	m.Get("/", index)
//  	m.Get("/hello", hello)
//
//  	log.Fatal(http.ListenAndServe(":8080", m))
//  }
//
// The Muxer matches incoming requests by the method and path and delegates
// to that assiciated function.
// Currently GET, POST, PUT, PATCH, DELETE and OPTIONS are supported methods.
//
// Named parameters are not supported.
//
// Path: /foo/bar
//
// Requests:
//  /foo/bar        matches the function
//  /foo/bar/       doesn't match, but redirects to /foo/bar
//  /foo/foo        doesn't match
//  /foo            doesn't match
package mux

import "net/http"

// Router is a http.Handler used to dispatch request to different handler
// functions with routes.
type Router struct {
	routes map[string]http.Handler

	// Enables automatic redirection if the requested path doesn't match but
	// a handler with a path without the trailing slash exists. Default: true
	RedirectTrailingSlash bool

	// Enables 'Method Not Allowed' responses when a handler for the the
	// path, but not the requested method exists. Default: true
	HandleMethodNotAllowed bool

	// Custom http.HandlerFunction which is called when no handler was found
	// for the requested route. Defaults: http.NotFound.
	NotFound http.HandlerFunc
}

// New initializes the Router.
// All configurable options are enabled by default.
func New() *Router {
	return &Router{
		RedirectTrailingSlash:  true,
		HandleMethodNotAllowed: true,
	}
}

// Make sure Router conforms with http.Handler interface.
var _ http.Handler = New()

// Get registers a new request handle for the GET method and the given path.
func (r *Router) Get(path string, h http.HandlerFunc) {
	r.add(http.MethodGet, path, h)
}

// Post registers a new request handle for the POST method and the given path.
func (r *Router) Post(path string, h http.HandlerFunc) {
	r.add(http.MethodPost, path, h)
}

// Put registers a new request handle for the PUT method and the given path.
func (r *Router) Put(path string, h http.HandlerFunc) {
	r.add(http.MethodPut, path, h)
}

// Delete registers a new request handle for the DELETE method and the given path.
func (r *Router) Delete(path string, h http.HandlerFunc) {
	r.add(http.MethodDelete, path, h)
}

// Options registers a new request handle for the OPTIONS method and the given path.
func (r *Router) Options(path string, h http.HandlerFunc) {
	r.add(http.MethodOptions, path, h)
}

// Patch registers a new request handle for the PATCH method and the given path.
func (r *Router) Patch(path string, h http.HandlerFunc) {
	r.add(http.MethodPatch, path, h)
}

func (r *Router) add(method string, path string, h http.HandlerFunc) {
	if path[0] != '/' {
		panic("Path must begin with '/' in path '" + path + "'")
	}

	if r.routes == nil {
		r.routes = make(map[string]http.Handler)
	}
	r.routes[method+path] = h
}

// ServeHTTP makes this router implement the http.Handler interface.
func (r *Router) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	path := req.URL.Path
	method := req.Method

	if h, ok := r.routes[method+path]; ok {
		h.ServeHTTP(w, req)
		return
	}

	// redirect for paths ending with a '/'
	if r.RedirectTrailingSlash {
		n := len(path)
		if n > 1 && path[n-1] == '/' {
			if _, ok := r.routes[method+path[:n-1]]; ok {
				code := http.StatusMovedPermanently
				if method != http.MethodGet {
					code = http.StatusPermanentRedirect
				}

				http.Redirect(w, req, path[:n-1], code)
			}
		}
	}

	// method not allowed
	if r.HandleMethodNotAllowed {
		if allowed := r.allowed(path); len(allowed) > 0 {
			w.Header().Set("Allow", allowed)
			if method != http.MethodOptions {
				http.Error(w,
					http.StatusText(http.StatusMethodNotAllowed),
					http.StatusMethodNotAllowed,
				)
			}

			return
		}
	}

	if r.NotFound != nil {
		r.NotFound.ServeHTTP(w, req)
	} else {
		http.NotFound(w, req)
	}
}

func (r *Router) allowed(path string) (allowed string) {
	methods := []string{
		http.MethodGet,
		http.MethodPost,
		http.MethodPut,
		http.MethodDelete,
		http.MethodPatch,
		http.MethodOptions,
	}

	for _, method := range methods {
		if _, ok := r.routes[method+path]; ok {
			if len(allowed) == 0 {
				allowed = method
			} else {
				allowed += ", " + method
			}
		}
	}

	return
}
