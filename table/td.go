package table

import (
	_ "embed"
	"github.com/boggydigital/compton"
)

const tdContentToken = ".Td"

var (
	//go:embed "markup/td.html"
	markupTd []byte
)

func NewTd() compton.Component {
	return compton.NewContainer(markupTd, tdContentToken)
}
