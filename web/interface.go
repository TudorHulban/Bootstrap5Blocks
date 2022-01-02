package web

import (
	"io"
	"text/template"
)

// IWeb Groups web components using simple method.
type IWeb interface {
	GetTemplateName() string
	Render(t *template.Template, w io.Writer) (int, error)
}
