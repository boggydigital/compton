package table

import (
	_ "embed"
	"github.com/boggydigital/compton"
)

var (
	//go:embed "markup/td.html"
	markupTd []byte
)

func NewTd() compton.Element {
	return compton.NewElement(markupTd)
}
