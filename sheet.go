package sheet

import (
	"html/template"
)

const viewPath = "views"
const layoutPath = "layouts"
const includePath = "includes"

// TemplateExtension is the default file extension checked when parsing templates
var TemplateExtension = ".tmpl"

// TemplatePath is the default root level project folder to look for view templates
var TemplatePath = "templates"

// TemplateBaseLayout is the default layout file which will be used for view templates
var TemplateBaseLayout = "base"

// ViewCollection exists to manage templates and hierarchical parsing for inheritance
type ViewCollection struct {
	templates map[string]*template.Template
}

// New Create ViewCollection from base view directory
func New() *ViewCollection {
	return &ViewCollection{
		templates: LoadTemplates(viewPath, layoutPath, includePath),
	}
}

// NewFromDir Create ViewCollection from specified directory paths
func NewFromDir(viewDirectory string, layoutDirectory string, includeDirectory string) *ViewCollection {
	return &ViewCollection{
		templates: LoadTemplates(viewDirectory, layoutDirectory, includeDirectory),
	}
}
