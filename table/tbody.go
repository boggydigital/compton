package table

import (
	_ "embed"
	"github.com/boggydigital/compton"
)

const tbodyContentToken = ".TBody"

var (
	//go:embed "markup/tbody.html"
	markupTBody []byte
)

func NewBody() compton.Component {
	return compton.NewContainer(markupTBody, tbodyContentToken)
}
