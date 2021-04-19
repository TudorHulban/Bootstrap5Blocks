package card

import (
	"blocks/web"
	"text/template"
)

type Content struct {
	Image web.URL
	Body  string
}

// Card Component
type Card struct {
	TemplateName string

	Content
}

var _ web.IWeb = (*Card)(nil)

func NewCo(c Content) *Card {
	return &Card{
		TemplateName: "layout.gohtml",

		Content: c,
	}
}

func (c *Card) Render(t *template.Template) (string, error) {
	return web.Render(t, c.TemplateName, c)
}
