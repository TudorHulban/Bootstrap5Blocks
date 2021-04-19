package container

import (
	"blocks/web"
	"blocks/webcontainers"
	"bytes"
	"errors"
	"strings"
	"text/template"
)

// Container Fluid Component
type Container struct {
	TemplateName string

	Markdown string
	Content  []string
}

var _ webcontainers.IWebContainer = (*Container)(nil)

func NewCo() *Container {
	return &Container{
		TemplateName: "container_fluid.gohtml",
	}
}

func (c *Container) SetContent(content []string) {
	c.Content = content
}

func (c *Container) Render(t *template.Template) (string, error) {
	tmpl := t.Lookup(c.TemplateName)
	if tmpl == nil {
		return "", errors.New("lookup did not work")
	}

	var buf bytes.Buffer

	err := tmpl.ExecuteTemplate(&buf, c.TemplateName, c)
	if err != nil {
		return "", err
	}

	return buf.String(), nil
}

func (c *Container) Inject(t *template.Template, blocks ...web.IWeb) error {
	for _, block := range blocks {
		markdown, err := block.Render(t)
		if err != nil {
			return err
		}

		c.Content = append(c.Content, markdown)
	}

	c.SetMarkdown()
	return nil
}

// Markdown Method produces accumulated markdown for component.
func (c *Container) SetMarkdown() {
	c.Markdown = strings.Join(c.Content, "")
}
