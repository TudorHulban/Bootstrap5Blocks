package blogauthor

import (
	"blocks/web"
	"io"
	"text/template"
)

type Content struct {
	AvatarSrc web.URL
	AvatarAlt string
	FullName  string
	Text      string
}

// BlogAuthor Component
type BlogAuthor struct {
	TemplateName string

	Content
}

var _ web.IWeb = (*BlogAuthor)(nil)

func NewCo(c Content) *BlogAuthor {
	return &BlogAuthor{
		TemplateName: "blog_author.gohtml",

		Content: c,
	}
}

func (c *BlogAuthor) Render(t *template.Template, w io.Writer) (int, error) {
	return web.Render(t, c, w)
}

func (c BlogAuthor) GetTemplateName() string {
	return c.TemplateName
}
