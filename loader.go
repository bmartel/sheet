package sheet

import (
	"html/template"
	"log"
	"strings"

	"github.com/mattn/go-zglob"
)

// NewTemplateLoader ...
func NewTemplateLoader(viewPath string, layoutPath string, includePath string) *TemplateLoader {
	viewTmplPath := getTemplatePath(viewPath)
	includeTmplPath := getTemplatePath(includePath)
	layoutTmplPath := getTemplatePath(layoutPath)
	defaultLayout := layoutTmplPath + "/" + TemplateBaseLayout + TemplateExtension

	return &TemplateLoader{
		viewPath:      viewTmplPath,
		includePath:   includeTmplPath,
		layoutPath:    layoutTmplPath,
		defaultLayout: defaultLayout,
	}
}

// TemplateLoader ...
type TemplateLoader struct {
	viewPath      string
	includePath   string
	layoutPath    string
	defaultLayout string
}

// Load builds up a hierarchical template cache
func (t *TemplateLoader) Load() map[string]*template.Template {
	views := make(map[string]*template.Template)

	viewTemplates := t.LoadViews()

	includeTemplates := t.LoadIncludes()

	for _, tmpl := range viewTemplates {
		viewKey, templateFiles := t.Merge(tmpl, includeTemplates)

		views[viewKey] = template.Must(template.ParseFiles(templateFiles...))
	}

	return views
}

// Merge arranges all the templates such that it is prepared for template.ParseFiles to execute
func (t *TemplateLoader) Merge(mainTmpl string, otherTemplates []string) (string, []string) {
	var specificIncludes []string
	var layoutTemplates []string

	viewKey := strings.Replace(mainTmpl, t.viewPath+"/", "", 1)
	parentDir := strings.Split(viewKey, "/")

	if parentDir[0] != viewKey {
		// Look for specific includes based on the parent most view directory
		specificIncludes, _ = zglob.Glob(t.includePath + "/" + parentDir[0] + "**/*" + TemplateExtension)

		// Add custom layout file if it exists
		layoutTemplates, _ = zglob.Glob(t.layoutPath + "/" + parentDir[0] + TemplateExtension)
	}

	if len(layoutTemplates) < 1 {
		layoutTemplates = append(layoutTemplates, t.defaultLayout)
	}

	includes := append(otherTemplates, specificIncludes...)
	files := append(includes, layoutTemplates...)
	return viewKey, append(files, mainTmpl)
}

// LoadViews ...
func (t *TemplateLoader) LoadViews() []string {
	views, err := zglob.Glob(t.viewPath + "/**/*" + TemplateExtension)
	if err != nil {
		log.Fatal(err)
	}

	return views
}

// LoadIncludes ...
func (t *TemplateLoader) LoadIncludes() []string {
	includes, err := zglob.Glob(t.includePath + "/*" + TemplateExtension)
	if err != nil {
		log.Fatal(err)
	}

	return includes
}

func getTemplatePath(relPath string) string {
	return strings.TrimRight(TemplatePath, "/") + "/" + strings.TrimRight(relPath, "/")
}
