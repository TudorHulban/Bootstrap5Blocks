package main

import (
	"embed"
	"fmt"
	"os"
	"text/template"

	"blocks/web/card"
	"blocks/web/layout"
)

//go:embed templates/*.gohtml
var f embed.FS

func main() {
	tmpl := template.New("views")

	tmpl, err := tmpl.ParseFS(f, "templates/*.gohtml")
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	f, errCreate := os.Create("output.html")
	if errCreate != nil {
		fmt.Println("errCreate: ", errCreate)
	}
	defer f.Close()

	l := layout.NewCo(layout.Content{
		Title: "This is title",
	})

	c := card.NewCo(card.Content{
		ImageSrc:   "https://bulma.io/images/placeholders/128x128.png",
		ImageAlt:   "Image Missing",
		Title:      "Card Title",
		Text:       "Lorem",
		ButtonText: "Read More",
	})

	l.Inject(tmpl, c)
	l.Markdown = l.GetMarkdown()

	errExe := tmpl.ExecuteTemplate(f, l.TemplateName, l)
	if errExe != nil {
		fmt.Println(errExe)
		os.Exit(3)
	}
}
