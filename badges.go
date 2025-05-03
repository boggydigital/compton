package compton

import (
	"github.com/boggydigital/compton/consts/color"
	"github.com/boggydigital/compton/consts/font_weight"
	"github.com/boggydigital/compton/consts/size"
)

type FormattedBadge struct {
	Title string
	Class string
	Color color.Color
}

func Badge(r Registrar, text string, bgColor, fgColor color.Color) *FspanElement {
	return Fspan(r, text).
		Width(size.Unset).
		FontSize(size.XXSmall).
		FontWeight(font_weight.Normal).
		PaddingInline(size.Small).
		PaddingBlock(size.XXSmall).
		BorderRadius(size.XXSmall).
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
