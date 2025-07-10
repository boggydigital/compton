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

func badgeContainer(r Registrar, bgColor, fgColor color.Color) *FspanElement {
	return Fspan(r, "").
		Width(size.Unset).
		FontSize(size.XXSmall).
		FontWeight(font_weight.Normal).
		PaddingInline(size.Small).
		PaddingBlock(size.XXSmall).
		BorderRadius(size.XSmall).
		BackgroundColor(bgColor).
		ForegroundColor(fgColor)
}

func badgeTextContainer(r Registrar, fgColor color.Color) *FspanElement {
	return Fspan(r, "").
		Width(size.Unset).
		FontSize(size.XXSmall).
		FontWeight(font_weight.Normal).
		PaddingBlock(size.XXSmall).
		BorderRadius(size.XSmall).
		ForegroundColor(fgColor).
		BackgroundColor(color.Transparent)
}

func smallBadgeContainer(r Registrar, bgColor, fgColor color.Color) *FspanElement {
	return Fspan(r, "").
		Width(size.Unset).
		FontSize(size.XXXSmall).
		FontWeight(font_weight.Normal).
		PaddingInline(size.XSmall).
		PaddingBlock(size.XXSmall).
		BorderRadius(size.XXSmall).
		BackgroundColor(bgColor).
		ForegroundColor(fgColor)
}

func smallBadgeTextContainer(r Registrar, fgColor color.Color) *FspanElement {
	return Fspan(r, "").
		Width(size.Unset).
		FontSize(size.XXXSmall).
		FontWeight(font_weight.Normal).
		PaddingBlock(size.XXSmall).
		BorderRadius(size.XXSmall).
		ForegroundColor(fgColor)
}

func Badge(r Registrar, text string, bgColor, fgColor color.Color) *FspanElement {
	bc := badgeContainer(r, bgColor, fgColor)
	bc.Append(Text(text))
	return bc
}

func SmallBadge(r Registrar, text string, bgColor, fgColor color.Color) *FspanElement {
	sbc := smallBadgeContainer(r, bgColor, fgColor)
	sbc.Append(Text(text))
	return sbc
}

func BadgeIcon(r Registrar, icon Symbol, text string, bgColor, fgColor color.Color) *FspanElement {
	bc := badgeContainer(r, bgColor, fgColor)
	bc.Append(SvgUse(r, icon))
	if text != "" {
		bc.Append(Text("&nbsp;"))
		bc.Append(Text(text))
	}
	return bc
}

func SmallBadgeIcon(r Registrar, icon Symbol, text string, bgColor, fgColor color.Color) *FspanElement {
	sbc := smallBadgeContainer(r, bgColor, fgColor)
	sbc.Append(SvgUse(r, icon))
	if text != "" {
		sbc.Append(Text("&nbsp;"))
		sbc.Append(Text(text))
	}
	return sbc
}

func BadgeText(r Registrar, text string, fgColor color.Color) *FspanElement {
	btc := badgeTextContainer(r, fgColor)
	btc.Append(Text(text))
	return btc
}
