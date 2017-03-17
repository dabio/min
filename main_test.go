package main

import (
	"html/template"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestIndex(t *testing.T) {
	c := context{
		templates: template.Must(template.ParseGlob("./views/*.html")),
	}

	tests := []struct {
		method string
		url    string
		h      http.HandlerFunc
		want   int
	}{
		{"GET", "/", c.get, 200},
		{"HEAD", "/", c.get, 200},
		{"POST", "/", c.post, 200},
	}

	for _, tt := range tests {
		req, _ := http.NewRequest(tt.method, tt.url, nil)
		rr := httptest.NewRecorder()
		tt.h.ServeHTTP(rr, req)
		if status := rr.Code; status != tt.want {
			t.Errorf("wrong status code: got %v want %v", status, tt.want)
		}
	}
}
