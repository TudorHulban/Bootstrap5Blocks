package main

import (
	"embed"
	"fmt"
	"os"
	"text/template"

	"blocks/web/blogauthor"
	"blocks/web/card"
	"blocks/web/layout"
	"blocks/webcontainers/row"
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

	// c := card.NewCo(card.Content{
	// 	ImageSrc:   "https://bulma.io/images/placeholders/128x128.png",
	// 	ImageAlt:   "Image Missing",
	// 	Title:      "Card Title",
	// 	Text:       "Lorem",
	// 	ButtonText: "Read More",
	// })

	author := blogauthor.NewCo(blogauthor.Content{
		AvatarSrc: "https://bulma.io/images/placeholders/64x64.png",
		FullName:  "John Smith",
		Text:      "Lorem ...",
	})

	l.Inject(tmpl, author)

	c1 := card.NewCo(card.Content{
		ImageSrc:   "https://bulma.io/images/placeholders/128x128.png",
		ImageAlt:   "",
		Title:      "Card 1 Title",
		Text:       "Lorem",
		ButtonText: "Read More",
	})

	c2 := card.NewCo(card.Content{
		ImageSrc:   "https://bulma.io/images/placeholders/128x128.png",
		ImageAlt:   "",
		Title:      "Card 2 Title",
		Text:       "Lorem",
		ButtonText: "Read More",
	})

	c3 := card.NewCo(card.Content{
		ImageSrc:   "https://bulma.io/images/placeholders/128x128.png",
		ImageAlt:   "",
		Title:      "Card 3 Title",
		Text:       "Lorem",
		ButtonText: "Read More",
	})

	r := row.NewCo()
	r.Inject(tmpl, c1, c2, c3)

	markdown, errRen := r.Render(tmpl)
	if errRen != nil {
		fmt.Println(errRen)
		os.Exit(3)
	}

	l.AppendToBody(markdown)
	l.Markdown = l.GetMarkdown()

	errExe := tmpl.ExecuteTemplate(f, l.TemplateName, l)
	if errExe != nil {
		fmt.Println(errExe)
		os.Exit(4)
	}
}
