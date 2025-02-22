package compton

import (
	"github.com/boggydigital/compton/consts/align"
	"github.com/boggydigital/compton/consts/color"
	"github.com/boggydigital/compton/consts/compton_atoms"
	"github.com/boggydigital/compton/consts/direction"
	"github.com/boggydigital/compton/consts/font_weight"
	"github.com/boggydigital/compton/consts/size"
	"strings"
)

type FrowElement struct {
	*BaseElement
	r Registrar
}

func (f *FrowElement) Elements(elements ...Element) *FrowElement {
	if fi := f.GetFirstElementByTagName(compton_atoms.FlexItems); fi != nil {
		fi.Append(elements...)
	}
	return f
}

func (f *FrowElement) PropVal(p string, vals ...string) *FrowElement {
	if fi := f.GetFirstElementByTagName(compton_atoms.FlexItems); fi != nil {
		row := FlexItems(f.r, direction.Row).ColumnGap(size.XSmall)
		row.Append(Fspan(f.r, p).ForegroundColor(color.Gray))
		row.Append(Fspan(f.r, strings.Join(vals, ", ")).TextAlign(align.Center))
		fi.Append(row)
	}
	return f
}

func (f *FrowElement) PropIcons(p string, symbols ...Symbol) *FrowElement {
	if fi := f.GetFirstElementByTagName(compton_atoms.FlexItems); fi != nil {
		row := FlexItems(f.r, direction.Row).ColumnGap(size.Small)
		row.Append(Fspan(f.r, p).ForegroundColor(color.Gray))
		for _, symbol := range symbols {
			row.Append(SvgUse(f.r, symbol))
		}
		fi.Append(row)
	}
	return f
}

func (f *FrowElement) Heading(title string) *FrowElement {
	if fi := f.GetFirstElementByTagName(compton_atoms.FlexItems); fi != nil {
		fi.Append(Fspan(f.r, title).FontWeight(font_weight.Bolder))
	}
	return f
}

func (f *FrowElement) Highlight(title string) *FrowElement {
	if fi := f.GetFirstElementByTagName(compton_atoms.FlexItems); fi != nil {
		fi.Append(Fspan(f.r, title).FontWeight(font_weight.Bolder).ForegroundColor(color.Orange))
	}
	return f
}

func (f *FrowElement) IconColor(symbol Symbol, c color.Color) *FrowElement {
	if fi := f.GetFirstElementByTagName(compton_atoms.FlexItems); fi != nil {
		fi.Append(SvgUse(f.r, symbol).ForegroundColor(c))
	}
	return f
}

func (f *FrowElement) LinkColor(title, href string, c color.Color) *FrowElement {
	if fi := f.GetFirstElementByTagName(compton_atoms.FlexItems); fi != nil {
		linkDecoration := Fspan(f.r, "").
			FontWeight(font_weight.Bolder).
			ForegroundColor(c)
		link := AText(title, href)
		link.SetAttribute("target", "_top")
		linkDecoration.Append(link)
		fi.Append(linkDecoration)
	}
	return f
}

func (f *FrowElement) FontSize(s size.Size) *FrowElement {
	if fi := f.GetFirstElementByTagName(compton_atoms.FlexItems); fi != nil {
		if flexItems, ok := fi.(*FlexItemsElement); ok {
			flexItems.FontSize(s)
		}
	}
	return f
}

func Frow(r Registrar) *FrowElement {
	frow := &FrowElement{
		BaseElement: NewElement(compton_atoms.Frow, nil),
		r:           r,
	}

	fi := FlexItems(r, direction.Row).
		ColumnGap(size.Small).
		RowGap(size.XSmall).
		AlignItems(align.Center)
	fi.AddClass("frow")

	frow.Append(fi)

	return frow
}
