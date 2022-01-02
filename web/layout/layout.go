package layout

import (
	"blocks/web"
	"io"
	"strings"
	"text/template"
)

type Content struct {
	Title string
	Body  []string
}

// Layout Component
type Layout struct {
	TemplateName string
	Markdown     string
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
func (c *Layout) Render(t *template.Template, w io.Writer) (int, error) {
	return web.Render(t, c, w)
}

// Markdown Method produces accumulated markdown for component.
func (c *Layout) HTML() string {
	return strings.Join(c.Body, "")
}

// SetBody Method should be used when we just want to render the passed body.
// Ex. when we already placed the markdown in a container.
func (c *Layout) SetBody(b []string) {
	c.Body = b
}

func (c *Layout) Write(markdown []byte) (int, error) {
	c.Body = append(c.Body, string(markdown))

	return len(c.Body), nil
}

func (c Layout) GetTemplateName() string {
	return c.TemplateName
}
