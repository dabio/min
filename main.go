package main

import (
	"log"
	"net/http"
	"os"
	"runtime"
	"time"
)

func init() {
	runtime.GOMAXPROCS(runtime.NumCPU())
}

func main() {
	m := http.NewServeMux()
	m.Handle("/css/", http.FileServer(http.Dir("./static/")))
	m.HandleFunc("/", track(index))

	s := &http.Server{
		Addr:           ":" + os.Getenv("PORT"),
		Handler:        m,
		ReadTimeout:    5 * time.Second,
		WriteTimeout:   10 * time.Second,
		IdleTimeout:    120 * time.Second,
		MaxHeaderBytes: 1 << 20,
	}
	log.Fatal(s.ListenAndServe())
}

func index(w http.ResponseWriter, r *http.Request) {
	if r.URL.String() != "/" {
		http.Redirect(w, r, "/", http.StatusMovedPermanently)
		return
	}
}

func track(fn http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		fn(w, r)
	}
}
