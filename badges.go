package compton

import (
	"github.com/boggydigital/compton/consts/color"
	"github.com/boggydigital/compton/consts/size"
)

type FormattedBadge struct {
	Title string
	Icon  Symbol
}

func BadgeIcon(r Registrar, icon Symbol, c color.Color) Element {
	return SvgUse(r, icon).ForegroundColor(c)
}

func BadgeText(r Registrar, text string, c color.Color) *FspanElement {
	return Fspan(r, text).
		Width(size.Unset).
		BackgroundColor(color.Transparent).
		ForegroundColor(c)
}
