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

	c := NewCo()

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
	tmpl, errParse := tmpl.ParseFiles("../../templates/"+c.TemplateName, "../../templates/"+c1.TemplateName)
	require.Nil(t, errParse)

	err := c.Inject(tmpl, c1, c2)
	require.NoError(t, err, "Injecting did not work")

	_, errRender := c.Render(tmpl, os.Stdout)
	require.NoError(t, errRender, "Did not render correctly.")
}
