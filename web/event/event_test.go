package event

import (
	"os"
	"testing"
	"text/template"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestEvents(t *testing.T) {
	tmpl := template.New("views")

	event1 := Event{
		ID:    "Event1",
		Title: "Title for Event 1",
		Text:  "This is text for Event 1",
	}

	event2 := Event{
		ID:    "Event2",
		Title: "Title for Event 2",
		Text:  "This is text for Event 2",
	}

	events := NewCo(event1, event2)

	tmpl, err := tmpl.ParseFiles("../../templates/" + events.TemplateName)
	require.NoError(t, err)

	_, errRe := events.Render(tmpl, os.Stdout)
	assert.NoError(t, errRe)
}
