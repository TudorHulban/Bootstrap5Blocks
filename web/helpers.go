package web

import (
	"bytes"
	"errors"
	"text/template"
)

func Render(t *template.Template, compoTemplateName string, component IWeb) (string, error) {
	tmpl := t.Lookup(compoTemplateName)
	if tmpl == nil {
		return "", errors.New("lookup did not work")
	}

	var buf bytes.Buffer

	if err := tmpl.ExecuteTemplate(&buf, compoTemplateName, component); err != nil {
		return "", err
	}

	return buf.String(), nil
}
