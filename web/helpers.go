package web

import (
	"bytes"

	"text/template"

	"github.com/pkg/errors"
)

func Render(t *template.Template, compoTemplateName string, component IWeb) (string, error) {
	tmpl := t.Lookup(compoTemplateName)
	if tmpl == nil {
		return "", errors.Errorf("lookup over %v templates did not work", len(t.Templates()))
	}

	var buf bytes.Buffer

	if err := tmpl.ExecuteTemplate(&buf, compoTemplateName, component); err != nil {
		return "", err
	}

	return buf.String(), nil
}
