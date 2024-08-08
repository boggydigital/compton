package elements

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
