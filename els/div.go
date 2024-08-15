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

func NewDiv() compton.Element {
	return compton.NewElement(atom.Div, markupDiv)
}

func NewDivText(txt string) compton.Element {
	div := NewDiv()
	div.Append(NewText(txt))
	return div
}
