package layout

import (
	"fmt"
	"testing"
	"text/template"

	"github.com/stretchr/testify/require"
)

const token = "xxx"

func TestLayout(t *testing.T) {
	tmpl := template.New("views")

	c := NewCo(Content{
		Title: token,
		Body:  "lorem",
	})

	tmpl, err := tmpl.ParseFiles("../../templates/" + c.TemplateName)
	require.Nil(t, err)

	s, errRender := c.Render(tmpl)
	require.Nil(t, errRender, "Did not render correctly.")
	require.Contains(t, s, token, "Does not contain token.")

	fmt.Println(s)
}
