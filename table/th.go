package table

import (
	_ "embed"
	"github.com/boggydigital/compton"
)

const thContentToken = ".Th"

var (
	//go:embed "markup/th.html"
	markupTh []byte
)

func NewTh() compton.Element {
	return compton.NewContainer(markupTh, thContentToken)
}
