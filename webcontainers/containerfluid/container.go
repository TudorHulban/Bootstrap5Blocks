package container

import (
	"blocks/web"
	"blocks/webcontainers"
	"io"
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

// SetContent Method to be used for injecting markdown (ex. from already rendered templates)
func (c *Container) SetContent(content []string) {
	c.Content = content
}

// Render Method should be used after markdown is set.
func (c *Container) Render(t *template.Template, w io.Writer) (int, error) {
	return web.Render(t, c, w)
}

// Markdown Method produces accumulated markdown for component.
func (c *Container) SetMarkdown() {
	c.Markdown = strings.Join(c.Content, "")
}

func (c Container) GetTemplateName() string {
	return c.TemplateName
}

func (c Container) Write(markdown []byte) (int, error) {
	c.Content = append(c.Content, string(markdown))

	return len(c.Content), nil
}
