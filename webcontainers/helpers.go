package webcontainers

import (
	"blocks/web"
	"io"
	"text/template"
)

func Render(t *template.Template, c web.IWeb, w io.Writer) (int, error) {
	return web.Render(t, c, w)
}
