package table

import (
	_ "embed"
	"github.com/boggydigital/compton"
	"golang.org/x/net/html/atom"
)

var (
	//go:embed "markup/td.html"
	markupTd []byte
)

func NewTd() compton.Element {
	return compton.NewElement(atom.Td, markupTd)
}
