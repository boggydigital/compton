package els

import (
	_ "embed"
	"github.com/boggydigital/compton"
	"golang.org/x/net/html/atom"
)

var (
	//go:embed "markup/div.html"
	markupDiv []byte
)

func Div() compton.Element {
	return compton.NewElement(atom.Div, markupDiv)
}

func DivText(txt string) compton.Element {
	div := Div()
	div.Append(Text(txt))
	return div
}
