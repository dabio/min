package main

import (
	"html/template"
	"log"
	"net/http"
	"os"
	"runtime"
	"time"

	"github.com/pinub/mux"
	"github.com/tdewolff/minify"
	"github.com/tdewolff/minify/js"
)

type context struct {
	templates *template.Template
}

func (c *context) render(w http.ResponseWriter, tmpl string, data interface{}) {
	w.Header().Set("Connection", "keep-alive")
	w.Header().Set("Content-Type", "text/html; charset=utf=8")

	if err := c.templates.ExecuteTemplate(w, tmpl+".html", data); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

func (c *context) get(w http.ResponseWriter, r *http.Request) {
	c.render(w, "index", nil)
}

func (c *context) post(w http.ResponseWriter, r *http.Request) {
	const mediatype = "text/javascript"

	m := minify.New()
	m.AddFunc(mediatype, js.Minify)

	s, err := m.String(mediatype, r.FormValue("content"))
	if err != nil {
		log.Panic(err)
	}

	data := struct {
		Input  string
		Output string
	}{
		r.FormValue("content"),
		s,
	}

	c.render(w, "index", data)
}

func init() {
	runtime.GOMAXPROCS(runtime.NumCPU())
}

func main() {
	c := context{
		templates: template.Must(template.ParseGlob("./views/*.html")),
	}

	m := mux.New()
	m.Get("/", track(c.get))
	m.Post("/", track(c.post))

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

func track(fn http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		defer func(start time.Time, r *http.Request) {
			if os.Getenv("ENV") != "production" {
				log.Printf("%s %s %s", r.Method, r.URL, time.Since(start))
			}
		}(time.Now(), r)

		fn(w, r)
	}
}
