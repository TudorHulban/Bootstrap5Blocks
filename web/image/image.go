package image

import (
	"blocks/web"
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

func (c *Image) Render(t *template.Template) (string, error) {
	c.prepareCSS()

	return web.Render(t, c.TemplateName, c)
}

func (c *Image) prepareCSS() {
	c.AdditionalCSSClasses = append(c.AdditionalCSSClasses, c.CSSClass)

	c.CSSClass = strings.Join(c.AdditionalCSSClasses, " ")
}