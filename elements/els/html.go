package els

import (
	_ "embed"
	"github.com/boggydigital/compton"
	"golang.org/x/net/html/atom"
)

var (
	//go:embed "markup/html.html"
	markupHtml []byte
)

func Html(lang string) compton.Element {
	html := compton.NewElement(atom.Html, markupHtml)
	if lang != "" {
		html.SetAttribute("lang", lang)
	}
	return html
}
