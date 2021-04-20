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

// Render Method to be used for rendering once the body is set.
func (c *Layout) Render(t *template.Template) (string, error) {
	return web.Render(t, c.TemplateName, c)
}

// Inject Method to be used to directly inject components as body.
// Sometimes a container could be used first
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

// SetBody Method should be used when we just want to render the passed body.
// Ex. when we already placed the markdown in a container.
func (c *Layout) SetBody(b []string) {
	c.Body = b
}

func (c *Layout) AppendToBody(markdown string) {
	c.Body = append(c.Body, markdown)
}
