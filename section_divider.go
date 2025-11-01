package compton

import (
	"github.com/boggydigital/compton/consts/align"
	"github.com/boggydigital/compton/consts/color"
	"github.com/boggydigital/compton/consts/compton_atoms"
	"github.com/boggydigital/compton/consts/direction"
	"github.com/boggydigital/compton/consts/size"
)

func SectionDivider(r Registrar, fmtBadge FormattedBadge) Element {

	r.RegisterStyles(DefaultStyle, compton_atoms.StyleName(compton_atoms.SectionDivider))

	sectionDividerRow := FlexItems(r, direction.Row).
		Width(size.FullWidth).
		MaxWidth(size.MaxWidth).
		AlignItems(align.Center).
		JustifyItems(align.Center).
		ColumnGap(size.Small).
		BackgroundColor(color.RepHighlight).
		ForegroundColor(color.RepForeground).
		BorderRadius(size.Small)

	titleFspan := Fspan(r, "").
		Width(size.Unset).
		FontSize(size.Small).
		PaddingBlock(size.Small).
		LineHeight(size.Normal)
	titleFspan.Append(Badges(r, fmtBadge))

	sectionDividerRow.Append(titleFspan)
	sectionDividerRow.AddClass("section-divider")

	return FICenter(r, sectionDividerRow)
}
