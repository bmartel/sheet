package main

import (
	"net/http"

	"github.com/bmartel/sheet"
)

type handler struct {
	views *sheet.ViewCollection
}

func main() {
	sheet.TemplateExtension = ".html"

	handler := &handler{
		views: sheet.New(),
	}

	http.HandleFunc("/", handler.index)
	http.HandleFunc("/about", handler.about)
	http.ListenAndServe(":8080", nil)
}

func (h *handler) index(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		h.notfound(w, r)
		return
	}

	data := map[string]interface{}{
		"title": "Example Index",
	}

	h.views.Render(w, "index.html", data)
}

func (h *handler) about(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/about" {
		h.notfound(w, r)
		return
	}

	data := map[string]interface{}{
		"title": "About Page",
	}

	h.views.Render(w, "about.html", data)
}

func (h *handler) notfound(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusNotFound)
}
