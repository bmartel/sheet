package sheet

import (
	"html/template"
	"log"
	"strings"

	"github.com/mattn/go-zglob"
)

// LoadTemplates builds up a hierarchical template cache
func LoadTemplates(viewPath string, layoutPath string, includePath string) map[string]*template.Template {
	views := make(map[string]*template.Template)
	viewTmplPath := getTemplatePath(viewPath)
	includeTmplPath := getTemplatePath(includePath)
	layoutTmplPath := getTemplatePath(layoutPath)
	defaultLayout := layoutTmplPath + "/" + TemplateBaseLayout + TemplateExtension

	viewTemplates, err := zglob.Glob(viewTmplPath + "/**/*" + TemplateExtension)
	if err != nil {
		log.Fatal(err)
	}

	includeTemplates, err := zglob.Glob(includeTmplPath + "/*" + TemplateExtension)
	if err != nil {
		log.Fatal(err)
	}

	for _, tmpl := range viewTemplates {
		var specificIncludes []string
		var layoutTemplates []string

		viewKey := strings.Replace(tmpl, viewTmplPath+"/", "", 1)
		parentDir := strings.Split(viewKey, "/")

		if parentDir[0] != viewKey {
			// Look for specific includes based on the parent most view directory
			specificIncludes, _ = zglob.Glob(includeTmplPath + "/" + parentDir[0] + "**/*" + TemplateExtension)

			// Add custom layout file if it exists
			layoutTemplates, _ = zglob.Glob(layoutTmplPath + "/" + parentDir[0] + TemplateExtension)
		}

		if len(layoutTemplates) < 1 {
			layoutTemplates = append(layoutTemplates, defaultLayout)
		}

		includes := append(includeTemplates, specificIncludes...)
		files := append(includes, layoutTemplates...)
		templateFiles := append(files, tmpl)

		views[viewKey] = template.Must(template.ParseFiles(templateFiles...))
	}

	return views
}

func getTemplatePath(relPath string) string {
	return strings.TrimRight(TemplatePath, "/") + "/" + strings.TrimRight(relPath, "/")
}
