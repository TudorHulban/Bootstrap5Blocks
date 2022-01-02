package web

import (
	"bytes"
	"io"

	"text/template"

	"github.com/pkg/errors"
)

func Render(t *template.Template, co IWeb, w io.Writer) (int, error) {
	tmpl := t.Lookup(co.GetTemplateName())
	if tmpl == nil {
		return 0, errors.Errorf("lookup over %v templates did not work", len(t.Templates()))
	}

	var buf bytes.Buffer

	if err := tmpl.ExecuteTemplate(&buf, co.GetTemplateName(), co); err != nil {
		return 0, err
	}

	return w.Write(buf.Bytes())
}
