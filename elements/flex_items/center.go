package flex_items

import (
	"github.com/boggydigital/compton"
	"github.com/boggydigital/compton/consts/align"
	"github.com/boggydigital/compton/consts/direction"
)

func Center(r compton.Registrar, elements ...compton.Element) compton.Element {
	row := FlexItems(r, direction.Row).
		JustifyContent(align.Center).
		AlignItems(align.Center)
	row.Append(elements...)
	return row
}
