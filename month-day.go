package compton

import (
	"strconv"
	"time"

	"github.com/boggydigital/compton/consts/align"
	"github.com/boggydigital/compton/consts/compton_atoms"
	"github.com/boggydigital/compton/consts/direction"
	"github.com/boggydigital/compton/consts/size"
	"github.com/boggydigital/compton/consts/wrap"
)

func MonthDay(r Registrar, month time.Month, day int) Element {

	r.RegisterStyles(DefaultStyle,
		compton_atoms.StyleName(compton_atoms.MonthDay))

	dateTimeContainer := FlexItems(r, direction.Column).
		FlexWrap(wrap.NoWrap).
		AlignItems(align.Center).
		Width(size.Normal).
		Height(size.Normal).
		RowGap(size.Zero).
		ColumnGap(size.Zero)
	dateTimeContainer.AddClass("month-day")

	if day < 1 || day > 31 {
		day = 1
	}

	monthStr := month.String()
	if len(monthStr) > 3 {
		monthStr = monthStr[:3]
	}

	monthElement := DivText(monthStr)
	monthElement.AddClass("month")
	dayElement := DivText(strconv.FormatInt(int64(day), 10))
	dayElement.AddClass("day")

	dateTimeContainer.Append(monthElement, dayElement)

	return dateTimeContainer
}
