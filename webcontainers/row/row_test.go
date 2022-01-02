package row

import (
	"os"
	"testing"
	"text/template"

	"blocks/web/card"

	"github.com/stretchr/testify/require"
)

const token = "xxx"

func TestContainer(t *testing.T) {
	tmpl := template.New("views")

	row := NewCo()

	c1 := card.NewCo(card.Content{
		ImageSrc:   "https://bulma.io/images/placeholders/128x128.png",
		ImageAlt:   token,
		Title:      "Card 1 Title",
		Text:       "Lorem",
		ButtonText: "Read More",
	})

	c2 := card.NewCo(card.Content{
		ImageSrc:   "https://bulma.io/images/placeholders/128x128.png",
		ImageAlt:   token,
		Title:      "Card 2 Title",
		Text:       "Lorem",
		ButtonText: "Read More",
	})

	// the adition of the files is done automatically when embedding
	tmpl, errParse := tmpl.ParseFiles("../../templates/"+row.TemplateName, "../../templates/"+c1.TemplateName)
	require.NoError(t, errParse)

	c1.Render(tmpl, row)
	c2.Render(tmpl, row)

	_, errRender := row.Render(tmpl, os.Stdout)
	require.NoError(t, errRender, "Did not render correctly.")
}
