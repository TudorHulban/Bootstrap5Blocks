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
	templateName string

	Markdown []string
}

var _ webcontainers.IWebContainer = (*Container)(nil)

func NewCo() *Container {
	return &Container{
		templateName: "container.gohtml",
	}
}

func (c *Container) SetMarkdown(markdown []string) {
	c.Markdown = markdown
}

func (c *Container) Render(t *template.Template) (string, error) {
	tmpl := t.Lookup(c.templateName)
	if tmpl == nil {
		return "", errors.New("lookup did not work")
	}

	var buf bytes.Buffer

	err := tmpl.ExecuteTemplate(&buf, c.templateName, c)
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

		c.Markdown = append(c.Markdown, markdown)
	}

	return nil
}

// Markdown Method produces accumulated markdown for component.
func (c *Container) GetMarkdown() string {
	return strings.Join(c.Markdown, "")
}
