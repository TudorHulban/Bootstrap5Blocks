package layout

import (
	"blocks/web"
	"strings"
	"text/template"
)

type Content struct {
	Title    string
	Body     []string
	Markdown string
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

func (c *Layout) Inject(t *template.Template, blocks ...web.IWeb) error {
	for _, block := range blocks {
		markdown, err := block.Render(t)
		if err != nil {
			return err
		}

		c.Body = append(c.Body, markdown)
	}

	return nil
}

// Markdown Method produces accumulated markdown for component.
func (c *Layout) GetMarkdown() string {
	return strings.Join(c.Body, "")
}
