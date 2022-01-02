package main

import (
	"embed"
	"fmt"
	"os"
	"text/template"

	"blocks/web/blogauthor"
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

	l := layout.NewCo(layout.Content{
		Title: "This is title",
	})

	blogauthor.NewCo(blogauthor.Content{
		AvatarSrc: "https://bulma.io/images/placeholders/64x64.png",
		FullName:  "John Smith",
		Text:      "Lorem ...",
	}).Render(tmpl, l)

	// c1 := card.NewCo(card.Content{
	// 	ImageSrc:   "https://bulma.io/images/placeholders/128x128.png",
	// 	ImageAlt:   "",
	// 	Title:      "Card 1 Title",
	// 	Text:       "Lorem",
	// 	ButtonText: "Read More",
	// })

	// c2 := card.NewCo(card.Content{
	// 	ImageSrc:   "https://bulma.io/images/placeholders/128x128.png",
	// 	ImageAlt:   "",
	// 	Title:      "Card 2 Title",
	// 	Text:       "Lorem",
	// 	ButtonText: "Read More",
	// })

	// c3 := card.NewCo(card.Content{
	// 	ImageSrc:   "https://bulma.io/images/placeholders/128x128.png",
	// 	ImageAlt:   "",
	// 	Title:      "Card 3 Title",
	// 	Text:       "Lorem",
	// 	ButtonText: "Read More",
	// })

	r := row.NewCo()
	// r.Inject(tmpl, c1, c2, c3)

	_, errRen := r.Render(tmpl, l)
	if errRen != nil {
		fmt.Println(errRen)
		os.Exit(3)
	}

	l.Markdown = l.HTML()

	f, errCreate := os.Create("output.html")
	if errCreate != nil {
		fmt.Println("errCreate: ", errCreate)
	}
	defer f.Close()

	l.Render(tmpl, f)

	// errExe := tmpl.ExecuteTemplate(f, l.TemplateName, l)
	// if errExe != nil {
	// 	fmt.Println(errExe)
	// 	os.Exit(4)
	// }
}
