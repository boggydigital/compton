package compton

import (
	"github.com/boggydigital/compton/consts/color"
	"github.com/boggydigital/compton/consts/font_weight"
	"github.com/boggydigital/compton/consts/size"
	"strconv"
	"strings"
)

type CountFormatter struct {
	single          string
	manyInSingleSet string
	manyInManySets  string
}

func (cf *CountFormatter) Title(from, to, total int) string {
	title := ""
	switch total {
	case 1:
		title = cf.single
	case to:
		if from > 0 {
			title = cf.formatManyInManySets(from, to, total)
		} else {
			title = cf.formatManyInSingleSet(total)
		}
	default:
		title = cf.formatManyInManySets(from, to, total)
	}
	return title
}

func (cf *CountFormatter) TitleElement(r Registrar, from, to, total int) Element {
	return Fspan(r, cf.Title(from, to, total)).
		ForegroundColor(color.Gray).
		FontSize(size.XSmall).
		FontWeight(font_weight.Normal)
}

func (cf *CountFormatter) formatManyInManySets(from, to, total int) string {
	ftt := strings.Replace(cf.manyInManySets, "{from}", strconv.Itoa(from+1), 1)
	ftt = strings.Replace(ftt, "{to}", strconv.Itoa(to), 1)
	ftt = strings.Replace(ftt, "{total}", strconv.Itoa(total), 1)
	return ftt
}

func (cf *CountFormatter) formatManyInSingleSet(total int) string {
	return strings.Replace(cf.manyInSingleSet, "{total}", strconv.Itoa(total), 1)
}

func NewCountFormatter(single, manyInSingleSet, manyInManySets string) *CountFormatter {
	return &CountFormatter{
		single:          single,
		manyInSingleSet: manyInSingleSet,
		manyInManySets:  manyInManySets,
	}
}
