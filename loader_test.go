package sheet_test

import (
	"testing"

	"github.com/bmartel/sheet"
)

func TestLoadTemplates(t *testing.T) {
	sheet.TemplatePath = "testdata"

	templates := sheet.NewTemplateLoader("views", "layouts", "includes").Load()
	expected := []string{"index.tmpl", "page/sample.tmpl"}

	t.Log(templates)

	if len(templates) < 1 {
		t.Error("Error loading templates")
		t.Log(expected)
	}

	if len(expected) != len(templates) {
		t.Error("Template files list does not match the expected")
		t.Log(expected)
		return
	}

	i := 0
	for k := range templates {
		if k != expected[i] {
			t.Error("Expected template: " + expected[i] + " got: " + k)
		}
		i++
	}
}

func TestLoadTemplatesAlternativeExtension(t *testing.T) {
	sheet.TemplateExtension = ".html"
	sheet.TemplatePath = "testdata"

	templates := sheet.NewTemplateLoader("views", "layouts", "includes").Load()
	expected := []string{"index.html", "page/sample.html"}

	t.Log(templates)

	if len(templates) < 1 {
		t.Error("Error loading templates")
		t.Log(expected)
	}

	if len(expected) != len(templates) {
		t.Error("Template files list does not match the expected")
		t.Log(expected)
		return
	}

	i := 0
	for k := range templates {
		if k != expected[i] {
			t.Error("Expected template: " + expected[i] + " got: " + k)
		}
		i++
	}
}

func TestMergeTemplates(t *testing.T) {
	sheet.TemplateExtension = ".tmpl"
	expected := []string{"testdata/layouts/base.tmpl", "testdata/views/index.tmpl", "testdata/includes/sidebar.tmpl"}
	expectedViewKey := "index.tmpl"

	loader := sheet.NewTemplateLoader("views", "layouts", "includes")

	viewKey, templateFiles := loader.Merge(
		"testdata/views/index.tmpl",
		[]string{"testdata/includes/sidebar.tmpl"},
	)

	if viewKey != expectedViewKey {
		t.Error("The incorrect view key was extracted. Expected: " + expectedViewKey + " got: " + viewKey)
	}

	t.Log(templateFiles)

	if len(expected) != len(templateFiles) {
		t.Error("Merged template files list does not match the expected")
		t.Log(expected)
		return
	}

	for k, v := range templateFiles {
		if v != expected[k] {
			t.Error("Expected template: " + expected[k] + " got: " + v)
		}
	}
}

func TestLoadViewTemplates(t *testing.T) {
	sheet.TemplatePath = "testdata"
	loader := sheet.NewTemplateLoader("views", "layouts", "includes")

	expected := []string{"testdata/views/index.tmpl", "testdata/views/page/sample.tmpl"}
	views := loader.LoadViews()

	t.Log(views)

	if len(expected) != len(views) {
		t.Error("Merged template files list does not match the expected")
		t.Log(expected)
		return
	}

	for k, v := range views {
		if v != expected[k] {
			t.Error("Expected template: " + expected[k] + " got: " + v)
		}
	}
}

func TestLoadIncludeTemplates(t *testing.T) {
	sheet.TemplatePath = "testdata"
	loader := sheet.NewTemplateLoader("views", "layouts", "includes")

	expected := []string{"testdata/includes/sidebar.tmpl"}
	includes := loader.LoadIncludes()

	t.Log(includes)

	if len(expected) != len(includes) {
		t.Error("Merged template files list does not match the expected")
		t.Log(expected)
		return
	}

	for k, v := range includes {
		if v != expected[k] {
			t.Error("Expected template: " + expected[k] + " got: " + v)
		}
	}
}
