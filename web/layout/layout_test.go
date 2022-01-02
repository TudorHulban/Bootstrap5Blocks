package layout

import (
	"os"
	"testing"
	"text/template"

	"github.com/stretchr/testify/require"
)

const token = "xxx"

func TestLayout(t *testing.T) {
	tmpl := template.New("views")

	layout := NewCo(Content{
		Title: token,
		Body:  []string{"lorem"},
	})

	tmpl, err := tmpl.ParseFiles("../../templates/" + layout.TemplateName)
	require.NoError(t, err)

	layout.Render(tmpl, os.Stdout)
}
