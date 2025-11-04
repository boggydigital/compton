package compton

import (
	"github.com/boggydigital/compton/consts/align"
	"github.com/boggydigital/compton/consts/color"
	"github.com/boggydigital/compton/consts/direction"
	"github.com/boggydigital/compton/consts/size"
	"github.com/boggydigital/compton/consts/wrap"
)

type FormattedBadge struct {
	Title string
	Icon  Symbol
	Color color.Color
}

func badgeIcon(r Registrar, icon Symbol, c color.Color) Element {
	return SvgUse(r, icon).ForegroundColor(c)
}

func badgeText(r Registrar, text string, c color.Color) *FspanElement {
	return Fspan(r, text).
		Width(size.Unset).
		BackgroundColor(color.Transparent).
		ForegroundColor(c).
		Padding(size.Unset)
}

func Badges(r Registrar, badges ...FormattedBadge) Element {

	badgesRow := FlexItems(r, direction.Row).
		FlexWrap(wrap.Wrap).
		RowGap(size.XSmall).
		ColumnGap(size.Small).
		FontSize(size.XXSmall).
		JustifyContent(align.Start).
		AlignItems(align.Center).
		AlignContent(align.Center).
		Width(size.Unset)

	for _, fb := range badges {

		if fb.Icon != NoSymbol && fb.Title != "" {
			badgeRow := FlexItems(r, direction.Row).
				FlexWrap(wrap.NoWrap).
				ColumnGap(size.Small).
				Width(size.Unset)
			badgeRow.Append(badgeIcon(r, fb.Icon, fb.Color), badgeText(r, fb.Title, fb.Color))
			badgesRow.Append(badgeRow)
		} else if fb.Icon != NoSymbol {
			badgesRow.Append(badgeIcon(r, fb.Icon, fb.Color))
		} else if fb.Title != "" {
			badgesRow.Append(badgeText(r, fb.Title, fb.Color))
		}

	}

	return badgesRow
}
