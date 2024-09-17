package recipes

import (
	"github.com/boggydigital/compton"
	"github.com/boggydigital/compton/consts/align"
	"github.com/boggydigital/compton/consts/direction"
	"github.com/boggydigital/compton/elements/flex_items"
)

func JustifyCenter(r compton.Registrar, elements ...compton.Element) compton.Element {
	row := flex_items.
		FlexItems(r, direction.Row).
		JustifyContent(align.Center)
	row.Append(elements...)
	return row
}
