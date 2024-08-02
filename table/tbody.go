package table

import (
	_ "embed"
	"github.com/boggydigital/compton"
)

var (
	//go:embed "markup/tbody.html"
	markupTBody []byte
)

func NewBody() compton.Element {
	return compton.NewElement(markupTBody)
}
