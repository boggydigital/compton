package table

import (
	_ "embed"
	"github.com/boggydigital/compton"
)

const theadContentToken = ".THead"

var (
	//go:embed "markup/thead.html"
	markupTHead []byte
)

func NewHead() compton.Element {
	return compton.NewContainer(markupTHead, theadContentToken)
}
