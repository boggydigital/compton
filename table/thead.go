package table

import (
	_ "embed"
	"github.com/boggydigital/compton"
)

var (
	//go:embed "markup/thead.html"
	markupTHead []byte
)

func NewHead() compton.Element {
	return compton.NewElement(markupTHead)
}
