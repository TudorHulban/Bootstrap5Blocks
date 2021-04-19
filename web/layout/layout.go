package layout

import (
	"blocks/web"
	"text/template"
)

type Content struct {
	Title string
	Body  string
}

// Layout Component
type Layout struct {
	TemplateName string

	Content
}

var _ web.IWeb = (*Layout)(nil)

func NewCo(c Content) *Layout {
	return &Layout{
		TemplateName: "layout.gohtml",

		Content: c,
	}
}

func (c *Layout) Render(t *template.Template) (string, error) {
	return web.Render(t, c.TemplateName, c)
}
