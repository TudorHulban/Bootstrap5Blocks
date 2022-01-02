package row

import (
	"blocks/web"
	"blocks/webcontainers"
	"io"
	"strings"
	"text/template"
)

// Row Component
type Row struct {
	TemplateName string

	Markdown string
	Content  []string
}

var _ webcontainers.IWebContainer = (*Row)(nil)

func NewCo() *Row {
	return &Row{
		TemplateName: "row.gohtml",
	}
}

// SetContent Method to be used for injecting markdown (ex. from already rendered templates)
func (c *Row) SetContent(content []string) {
	c.Content = content
}

// Render Method should be used after markdown is set.
func (c *Row) Render(t *template.Template, w io.Writer) (int, error) {
	return web.Render(t, c, w)
}

// Inject Method takes several components which it renders and decorates.
// func (c *Row) Inject(t *template.Template, blocks ...web.IWeb) error {
// 	for i, block := range blocks {
// 		markdown, err := block.Render(t)
// 		if err != nil {
// 			return errors.WithMessagef(err, "at injection of component number %v", i)
// 		}

// 		c.Content = append(c.Content, markdown)
// 	}

// 	c.SetMarkdown()
// 	return nil
// }

// Markdown Method produces accumulated markdown for component.
func (c *Row) SetMarkdown() {
	c.Markdown = strings.Join(c.Content, "")
}

func (c Row) GetTemplateName() string {
	return c.TemplateName
}
