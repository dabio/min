package main

import (
	"html/template"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestIndex(t *testing.T) {
	tests := []struct {
		method string
		url    string
		want   int
	}{
		{"GET", "/", 200},
		{"HEAD", "/", 200},
		{"POST", "/", 200},

		{"PATCH", "/", 405},
		{"PUT", "/", 405},

		{"GET", "/blah", http.StatusMovedPermanently},
		{"POST", "/sth", http.StatusMovedPermanently},
	}

	c := context{
		templates: template.Must(template.ParseGlob("./views/*.html")),
	}

	for _, tt := range tests {
		req, _ := http.NewRequest(tt.method, tt.url, nil)
		rr := httptest.NewRecorder()
		h := http.HandlerFunc(c.index)
		h.ServeHTTP(rr, req)
		if status := rr.Code; status != tt.want {
			t.Errorf("wrong status code: got %v want %v", status, tt.want)
		}
	}
}
