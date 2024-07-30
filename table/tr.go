package table

import (
	_ "embed"
	"github.com/boggydigital/compton"
)

const trContentToken = ".Tr"

var (
	//go:embed "markup/tr.html"
	markupTr []byte
)

func NewTr() compton.Component {
	return compton.NewContainer(markupTr, trContentToken)
}
