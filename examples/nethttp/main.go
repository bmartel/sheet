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
	http.HandleFunc("/sample", handler.sample)
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

	h.views.HTML(w, "index.html", data)
}

func (h *handler) about(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/about" {
		h.notfound(w, r)
		return
	}

	data := map[string]interface{}{
		"title": "About Page",
	}

	h.views.HTML(w, "about.html", data)
}

// Product example
type Product struct {
	Name  string
	Price int64
}

func (h *handler) sample(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/sample" {
		h.notfound(w, r)
		return
	}

	data := map[string]interface{}{
		"title": "This is a sample product list",
		"products": []Product{
			{
				Name:  "Stuff",
				Price: 1234,
			},
			{
				Name:  "Things",
				Price: 4567,
			},
			{
				Name:  "Doodads",
				Price: 8912,
			},
			{
				Name:  "Test",
				Price: 170,
			},
		},
	}

	h.views.HTML(w, "product/list.html", data)
}

func (h *handler) notfound(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusNotFound)
}
