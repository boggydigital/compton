package svg_inline

import (
	"github.com/boggydigital/compton"
	"github.com/boggydigital/compton/compton_atoms"
)

func New(s Symbol) compton.Element {
	return compton.NewElement(compton_atoms.SvgInlineIcon, MarkupSymbols[s])
}
