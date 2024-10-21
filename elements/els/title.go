package els

import (
	_ "embed"
	"github.com/boggydigital/compton"
	"golang.org/x/net/html/atom"
)

var (
	//go:embed "markup/title.html"
	markupTitle []byte
)

func Title(txt string) compton.Element {
	title := compton.NewElement(atom.Title, markupTitle)
	title.Append(Text(txt))
	return title
}
