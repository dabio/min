package main

import (
	"html/template"
	"log"
	"net/http"
	"os"
	"runtime"
	"time"
)

type context struct {
	templates *template.Template
}

func (c *context) render(w http.ResponseWriter, tmpl string, data interface{}) {
	if err := c.templates.ExecuteTemplate(w, tmpl+".html", data); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}
func (c *context) index(w http.ResponseWriter, r *http.Request) {
	if r.URL.String() != "/" {
		http.Redirect(w, r, "/", http.StatusMovedPermanently)
		return
	}
	if r.Method == "HEAD" {
		return
	}

	var data interface{}

	switch r.Method {
	default:
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	case "GET":
	case "POST":

	}

	w.Header().Set("Content-Type", "text/html; charset=utf=8")
	c.render(w, "index", data)
}

func init() {
	runtime.GOMAXPROCS(runtime.NumCPU())
}

func main() {
	c := context{
		templates: template.Must(template.ParseGlob("./views/*.html")),
	}

	m := http.NewServeMux()
	m.Handle("/css/", http.FileServer(http.Dir("./static/")))
	m.HandleFunc("/", track(c.index))

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
