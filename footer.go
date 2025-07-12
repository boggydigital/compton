package compton

import (
	"github.com/boggydigital/compton/consts/font_weight"
	"github.com/boggydigital/compton/consts/size"
)

func Footer(r Registrar, greeting, href string) Element {

	link := A(href)
	link.Append(Fspan(r, greeting).FontWeight(font_weight.Bolder))

	row := FICenter(r, link).
		ColumnGap(size.XSmall).
		FontSize(size.XSmall)

	return row
}
