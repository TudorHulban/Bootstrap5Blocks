package container

import (
	"testing"
	"text/template"

	"blocks/web/card"

	"github.com/stretchr/testify/require"
)

const token = "xxx"

func TestContainer(t *testing.T) {
	tmpl := template.New("views")

	container := NewCo()

	c1 := card.NewCo(card.Content{
		ImageSrc:   "https://bulma.io/images/placeholders/128x128.png",
		ImageAlt:   token,
		Title:      "Card Title",
		Text:       "Lorem",
		ButtonText: "Read More",
	})

	// the adition of the files is done automatically when embedding
	tmpl, errParse := tmpl.ParseFiles("../../templates/"+container.TemplateName, "../../templates/"+c1.TemplateName)
	require.Nil(t, errParse)

	c1.Render(tmpl, container)
}
