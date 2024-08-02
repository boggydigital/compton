package table

import (
	_ "embed"
	"github.com/boggydigital/compton"
)

var (
	//go:embed "markup/th.html"
	markupTh []byte
)

func NewTh() compton.Element {
	return compton.NewElement(markupTh)
}
