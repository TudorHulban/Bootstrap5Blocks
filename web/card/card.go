package card

import (
	"blocks/web"
	"io"
	"text/template"
)

type Content struct {
	ImageSrc   web.URL
	ImageAlt   string
	Title      string
	Text       string
	ButtonText string
	ButtonURL  web.URL
}

// Card Component
type Card struct {
	TemplateName string

	Content
}

var _ web.IWeb = (*Card)(nil)

func NewCo(c Content) *Card {
	return &Card{
		TemplateName: "card.gohtml",

		Content: c,
	}
}

func (c *Card) Render(t *template.Template, w io.Writer) (int, error) {
	return web.Render(t, c, w)
}

func (c Card) GetTemplateName() string {
	return c.TemplateName
}
