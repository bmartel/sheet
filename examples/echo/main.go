package main

import (
	"io"
	"net/http"

	"github.com/bmartel/sheet"
	"github.com/labstack/echo"
	"github.com/labstack/echo/engine/standard"
	"github.com/labstack/echo/middleware"
)

type view struct {
	views *sheet.ViewCollection
}

func (v *view) Render(w io.Writer, name string, data interface{}, c echo.Context) error {
	return v.views.Render(w, name, data)
}

// Product example
type Product struct {
	Name  string
	Price int64
}

func main() {
	e := echo.New()

	sheet.TemplateExtension = ".html"
	viewRenderer := &view{
		views: sheet.New(),
	}

	e.SetRenderer(viewRenderer)

	// Middleware
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	// Route handlers
	e.Get("/", func(c echo.Context) error {
		data := map[string]interface{}{
			"title": "Example Index",
		}

		return c.Render(http.StatusOK, "index.html", data)
	})

	e.Get("/about", func(c echo.Context) error {
		data := map[string]interface{}{
			"title": "About Page",
		}

		return c.Render(http.StatusOK, "about.html", data)
	})

	e.Get("/sample", func(c echo.Context) error {
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

		return c.Render(http.StatusOK, "product/list.html", data)
	})

	// Start server
	e.Run(standard.New(":8080"))
}
