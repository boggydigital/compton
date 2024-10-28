package compton

import (
	"github.com/boggydigital/compton/consts/align"
	"github.com/boggydigital/compton/consts/color"
	"github.com/boggydigital/compton/consts/direction"
	"github.com/boggydigital/compton/consts/font_weight"
	"github.com/boggydigital/compton/consts/size"
	"golang.org/x/exp/maps"
	"golang.org/x/exp/slices"
	"strings"
)

func Query(r Registrar, query map[string][]string, titles map[string]string, clearHref, clearTitle string) Element {
	if len(query) == 0 {
		return nil
	}

	sqStack := FlexItems(r, direction.Row).
		RowGap(size.Small).
		JustifyContent(align.Center).
		FontSize(size.Small)

	sortedProperties := maps.Keys(query)
	slices.Sort(sortedProperties)

	for _, property := range sortedProperties {
		values := query[property]
		span := Span()
		propertyTitleLink := A("#" + titles[property])
		propertyTitleText := Fspan(r, titles[property]+": ").
			ForegroundColor(color.Gray)
		propertyTitleLink.Append(propertyTitleText)
		fmtValues := make([]string, 0, len(values))
		for _, value := range values {
			fmtVal := value
			if pt, ok := titles[value]; ok {
				fmtVal = pt
			}
			fmtValues = append(fmtValues, fmtVal)
		}
		propertyValue := Fspan(r, strings.Join(fmtValues, ", ")).
			FontWeight(font_weight.Bolder)
		span.Append(propertyTitleLink, propertyValue)
		sqStack.Append(span)
	}

	clearLink := A(clearHref)
	clearText := Fspan(r, clearTitle).
		ForegroundColor(color.Blue).FontWeight(font_weight.Bolder)
	clearLink.Append(clearText)
	sqStack.Append(clearLink)

	return sqStack
}
