package svg_inline

import (
	"github.com/boggydigital/compton"
	"github.com/boggydigital/compton/consts/compton_atoms"
)

func SvgInline(s Symbol) compton.Element {
	return compton.NewElement(compton_atoms.SvgInlineIcon, MarkupSymbols[s])
}
