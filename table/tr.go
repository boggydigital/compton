package table

import (
	_ "embed"
	"github.com/boggydigital/compton"
)

var (
	//go:embed "markup/tr.html"
	markupTr []byte
)

func NewTr() compton.Element {
	return compton.NewElement(markupTr)
}
