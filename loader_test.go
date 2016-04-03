package sheet_test

import (
	"testing"

	"github.com/bmartel/sheet"
)

func TestLoadAllTemplates(t *testing.T) {
	sheet.TemplatePath = "testdata"

	templates := sheet.LoadTemplates("views", "layouts", "includes")

	if len(templates) < 1 {
		t.Error("Error loading templates")
	}

	t.Log(templates)
}

func TestLoadAllTemplatesAlternativeExtension(t *testing.T) {
	sheet.TemplateExtension = ".html"
	sheet.TemplatePath = "testdata"

	templates := sheet.LoadTemplates("views", "layouts", "includes")

	if len(templates) < 1 {
		t.Error("Error loading templates")
	}

	t.Log(templates)
}
