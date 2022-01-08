package main

import (
	"embed"
	"fmt"
	"os"
	"text/template"

	"blocks/web/blogauthor"
	"blocks/web/card"
	"blocks/web/event"
	"blocks/web/layout"
	"blocks/webcontainers/row"
)

//go:embed templates/*.gohtml
var fs embed.FS

func main() {
	tmpl := template.New("views")
	tmpl, err := tmpl.ParseFS(fs, "templates/*.gohtml")
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	layout := layout.NewCo(layout.Content{
		Title: "This is title",
	})

	blogauthor.NewCo(blogauthor.Content{
		AvatarSrc: "https://bulma.io/images/placeholders/64x64.png",
		FullName:  "John Smith",
		Text:      "Lorem ...",
	}).Render(tmpl, layout)

	row := row.NewCo()

	card.NewCo(card.Content{
		ImageSrc:   "https://bulma.io/images/placeholders/128x128.png",
		ImageAlt:   "",
		Title:      "Card 1 Title",
		Text:       "Lorem",
		ButtonText: "Read More",
	}).Render(tmpl, row)

	card.NewCo(card.Content{
		ImageSrc:   "https://bulma.io/images/placeholders/128x128.png",
		ImageAlt:   "",
		Title:      "Card 2 Title",
		Text:       "Lorem",
		ButtonText: "Read More",
	}).Render(tmpl, row)

	card.NewCo(card.Content{
		ImageSrc:   "https://bulma.io/images/placeholders/128x128.png",
		ImageAlt:   "",
		Title:      "Card 3 Title",
		Text:       "Lorem",
		ButtonText: "Read More",
	}).Render(tmpl, row)

	row.Render(tmpl, layout)

	event1 := event.Event{
		ID:    "Event1",
		Title: "Title for Event 1",
		Text:  "This is text for Event 1",
	}

	event2 := event.Event{
		ID:    "Event2",
		Title: "Title for Event 2",
		Text:  "This is text for Event 2",
	}

	events := event.NewCo(event1, event2)
	events.Render(tmpl, layout)

	f, errCreate := os.Create("output.html")
	if errCreate != nil {
		fmt.Println("errCreate: ", errCreate)
	}
	defer f.Close()

	layout.Render(tmpl, f)
}
