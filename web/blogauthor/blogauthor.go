package blogauthor

import (
	"blocks/web"
	"text/template"
)

type Content struct {
	AvatarSrc web.URL
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

func (c *BlogAuthor) Render(t *template.Template) (string, error) {
	return web.Render(t, c.TemplateName, c)
}
