package image

import (
	"blocks/web"
	"io"
	"strings"
	"text/template"
)

type Content struct {
	ImageSrc             web.URL
	ImageAlt             string
	AdditionalCSSClasses []string
}

// Image Component
type Image struct {
	TemplateName string
	CSSClass     string

	Content
}

var _ web.IWeb = (*Image)(nil)

func NewImage(c Content) *Image {
	return &Image{
		TemplateName: "image.gohtml",
		CSSClass:     "img-fluid",

		Content: c,
	}
}

func NewImageThumbnail(c Content) *Image {
	return &Image{
		TemplateName: "image.gohtml",
		CSSClass:     "img-thumbnail",

		Content: c,
	}
}

// Render Method should be used after markdown is set.
func (c *Image) Render(t *template.Template, w io.Writer) (int, error) {
	c.prepareCSS()

	return web.Render(t, c, w)
}

func (c *Image) prepareCSS() {
	c.AdditionalCSSClasses = append(c.AdditionalCSSClasses, c.CSSClass)

	c.CSSClass = strings.Join(c.AdditionalCSSClasses, " ")
}

func (c Image) GetTemplateName() string {
	return c.TemplateName
}
