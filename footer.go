package compton

import (
	"github.com/boggydigital/compton/consts/font_weight"
	"github.com/boggydigital/compton/consts/size"
)

func Footer(r Registrar, title, href, from string) Element {

	link := A(href)
	link.Append(Fspan(r, title).FontWeight(font_weight.Bolder))

	row := FICenter(r, Fspan(r, "👋"), Fspan(r, "from"), link, Fspan(r, from)).
		ColumnGap(size.XSmall).
		FontSize(size.Small)

	return row
}
