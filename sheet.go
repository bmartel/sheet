package sheet

import (
	"fmt"
	"html/template"
	"io"
	"net/http"

	"github.com/oxtoacart/bpool"
)

const viewPath = "views"
const layoutPath = "layouts"
const includePath = "includes"

// TemplateExtension is the default file extension checked when parsing templates
var TemplateExtension = ".tmpl"

// TemplatePath is the default root level project folder to look for view templates
var TemplatePath = "templates"

// DefaultLayout is the default layout file which will be used for view templates
var DefaultLayout = "base"

var pool *bpool.BufferPool

func init() {
	pool = bpool.NewBufferPool(48)
}

// New Create ViewCollection from base view directory
func New() *ViewCollection {
	return &ViewCollection{
		templates: NewTemplateLoader(viewPath, layoutPath, includePath).Load(),
	}
}

// NewFromDir Create ViewCollection from specified directory paths
func NewFromDir(viewDirectory string, layoutDirectory string, includeDirectory string) *ViewCollection {
	return &ViewCollection{
		templates: NewTemplateLoader(viewDirectory, layoutDirectory, includeDirectory).Load(),
	}
}

// ViewCollection exists to manage templates and hierarchical parsing for inheritance purposes
type ViewCollection struct {
	templates map[string]*template.Template
}

// HTML render the template with content type header set
func (v *ViewCollection) HTML(w http.ResponseWriter, name string, data interface{}) error {
	w.Header().Set("Content-Type", "text/html; charset=utf-8")

	return v.Render(w, name, data)
}

// Render out the given template with data
func (v *ViewCollection) Render(w io.Writer, name string, data interface{}) error {
	tmpl, ok := v.templates[name]
	if !ok {
		return fmt.Errorf("The template %s does not exist.", name)
	}

	// Create a buffered writer to check for template errors so they can be reported without
	// completely halting system operation
	buf := pool.Get()
	defer pool.Put(buf)

	err := tmpl.Execute(buf, data)
	if err != nil {
		// Probably should find a better way to display this, and maybe some debug trace view
		// in development would be useful
		return err
	}

	// Write the buffered output to the io.Writer
	buf.WriteTo(w)

	return nil
}
