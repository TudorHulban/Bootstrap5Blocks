package event

import (
	"blocks/web"
	"io"
	"text/template"
)

type Event struct {
	ID    string
	Title string
	Text  string
}

type Events struct {
	TemplateName string

	Content []Event
}

var _ web.IWeb = (*Events)(nil)

func NewCo(e ...Event) *Events {
	var content []Event
	content = append(content, e...)

	return &Events{
		TemplateName: "events.gohtml",

		Content: content,
	}
}

func (e *Events) Render(t *template.Template, w io.Writer) (int, error) {
	return web.Render(t, e, w)
}

func (e Events) GetTemplateName() string {
	return e.TemplateName
}
