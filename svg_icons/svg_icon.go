package svg_icons

import (
	"github.com/boggydigital/compton"
	"github.com/boggydigital/compton/compton_atoms"
)

func NewIcon(s Symbol) compton.Element {
	return compton.NewElement(compton_atoms.SvgIcon, MarkupSymbols[s])
}
