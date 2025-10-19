package compton

import (
	"github.com/boggydigital/compton/consts/color"
	"github.com/boggydigital/compton/consts/font_weight"
	"github.com/boggydigital/compton/consts/size"
)

type FormattedBadge struct {
	Title string
	Icon  Symbol
}

func badgeContainer(r Registrar, c color.Color) *FspanElement {
	return Fspan(r, "").
		Width(size.Unset).
		FontSize(size.XXSmall).
		FontWeight(font_weight.Normal).
		BorderRadius(size.XSmall).
		BackgroundColor(color.Transparent).
		ForegroundColor(c)
}

func BadgeIcon(r Registrar, icon Symbol, c color.Color) *FspanElement {
	bc := badgeContainer(r, c)
	bc.Append(SvgUse(r, icon))
	return bc
}

func BadgeText(r Registrar, text string, c color.Color) *FspanElement {
	btc := badgeContainer(r, c)
	btc.Append(Text(text))
	return btc
}
