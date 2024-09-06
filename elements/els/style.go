package els

import (
	_ "embed"
	"github.com/boggydigital/compton"
	"golang.org/x/net/html/atom"
)

var (
	//go:embed "markup/style.html"
	markupStyle []byte
)

func Style(styles []byte) compton.Element {
	style := compton.NewElement(atom.Style, markupStyle)
	style.Append(Text(string(styles)))
	return style
}
