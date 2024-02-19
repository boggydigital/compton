package main

import (
	"embed"
	"html/template"
	"os"
)

var (
	tmpl *template.Template
	//go:embed "templates/*.gohtml"
	templates embed.FS
)

func main() {

	tmpl = template.Must(
		template.
			New("").
			ParseFS(templates, "templates/*.gohtml"))

	outFile, err := os.Create("colors_tester.html")
	if err != nil {
		panic(err)
	}
	defer outFile.Close()

	if err := tmpl.ExecuteTemplate(outFile, "colors-tester", nil); err != nil {
		panic(err)
	}
}
