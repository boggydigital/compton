package compton

import (
	"github.com/boggydigital/compton/consts/color"
	"github.com/boggydigital/compton/consts/font_weight"
	"github.com/boggydigital/compton/consts/size"
)

type FormattedBadge struct {
	Title      string
	Icon       Symbol
	Class      string
	Background color.Color
	Foreground color.Color
}

func Badge(r Registrar, text string, bgColor, fgColor color.Color) *FspanElement {
	return Fspan(r, text).
		Width(size.Unset).
		FontSize(size.XXSmall).
		FontWeight(font_weight.Normal).
		PaddingInline(size.Small).
		PaddingBlock(size.XXSmall).
		BorderRadius(size.XSmall).
		BackgroundColor(bgColor).
		ForegroundColor(fgColor)
}

func SmallBadge(r Registrar, text string, bgColor, fgColor color.Color) *FspanElement {
	return Fspan(r, text).
		Width(size.Unset).
		FontSize(size.XXXSmall).
		FontWeight(font_weight.Normal).
		PaddingInline(size.XSmall).
		PaddingBlock(size.XXSmall).
		BorderRadius(size.XXSmall).
		BackgroundColor(bgColor).
		ForegroundColor(fgColor)
}

func BadgeIcon(r Registrar, icon Symbol, bgColor, fgColor color.Color) *FspanElement {
	span := Badge(r, "", bgColor, fgColor)
	span.Append(SvgUse(r, icon))
	return span
}

func SmallBadgeIcon(r Registrar, icon Symbol, bgColor, fgColor color.Color) *FspanElement {
	span := SmallBadge(r, "", bgColor, fgColor)
	span.Append(SvgUse(r, icon))
	return span
}
