package compton

import (
	"github.com/boggydigital/compton/consts/align"
	"github.com/boggydigital/compton/consts/color"
	"github.com/boggydigital/compton/consts/direction"
	"github.com/boggydigital/compton/consts/size"
)

func SectionDivider(r Registrar, content ...Element) Element {

	sectionDividerRow := FlexItems(r, direction.Row).
		Width(size.FullWidth).
		MaxWidth(size.MaxWidth).
		AlignItems(align.Center).
		JustifyItems(align.Center).
		ColumnGap(size.Small).
		BackgroundColor(color.Highlight).
		BorderRadius(size.XSmall)

	titleFspan := Fspan(r, "").
		Width(size.Unset).
		FontSize(size.Small).
		PaddingBlock(size.Small)
	titleFspan.Append(content...)

	sectionDividerRow.Append(titleFspan)

	return FICenter(r, sectionDividerRow)
}
