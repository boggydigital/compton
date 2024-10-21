package els

import (
	_ "embed"
	"github.com/boggydigital/compton"
	"github.com/boggydigital/compton/consts/compton_atoms"
)

var (
	//go:embed "markup/doctype.html"
	markupDoctype []byte
)

func Doctype() compton.Element {
	return compton.NewElement(compton_atoms.Doctype, markupDoctype)
}
