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

func (f *FrowElement) PropVal(p string, vals ...string) Element {
	if fi := f.GetFirstElementByTagName(compton_atoms.FlexItems); fi != nil {
		row := FlexItems(f.r, direction.Row).ColumnGap(size.XSmall)
		row.Append(Fspan(f.r, p).ForegroundColor(color.RepGray))
		row.Append(Fspan(f.r, strings.Join(vals, ", ")).TextAlign(align.Center))
		fi.Append(row)
		return row
	}
	return nil
}

func (f *FrowElement) PropLinkColor(p string, c color.Color, title, href string) Element {
	if fi := f.GetFirstElementByTagName(compton_atoms.FlexItems); fi != nil {
		row := FlexItems(f.r, direction.Row).ColumnGap(size.XSmall)
		row.Append(Fspan(f.r, p).ForegroundColor(color.RepGray))
		linkDecoration := Fspan(f.r, "").
			FontWeight(font_weight.Bolder).
			ForegroundColor(c)
		link := AText(title, href)
		link.SetAttribute("target", "_top")
		linkDecoration.Append(link)
		row.Append(linkDecoration)
		fi.Append(row)
		return row
	}
	return nil
}

func (f *FrowElement) PropIcons(p string, symbols ...Symbol) Element {
	if fi := f.GetFirstElementByTagName(compton_atoms.FlexItems); fi != nil {
		row := FlexItems(f.r, direction.Row).ColumnGap(size.Small)
		row.Append(Fspan(f.r, p).ForegroundColor(color.RepGray))
		for _, symbol := range symbols {
			row.Append(SvgUse(f.r, symbol))
		}
		fi.Append(row)
		return row
	}
	return nil
}

func (f *FrowElement) Heading(title string) Element {
	if fi := f.GetFirstElementByTagName(compton_atoms.FlexItems); fi != nil {
		heading := Fspan(f.r, title).FontWeight(font_weight.Bolder)
		fi.Append(heading)
		return heading

	}
	return nil
}

func (f *FrowElement) Highlight(title string) Element {
	if fi := f.GetFirstElementByTagName(compton_atoms.FlexItems); fi != nil {
		highlight := Fspan(f.r, title).FontWeight(font_weight.Bolder).ForegroundColor(color.Orange)
		fi.Append(highlight)
		return highlight
	}
	return nil
}

func (f *FrowElement) IconColor(symbol Symbol, c color.Color) Element {
	if fi := f.GetFirstElementByTagName(compton_atoms.FlexItems); fi != nil {
		icon := SvgUse(f.r, symbol).ForegroundColor(c)
		fi.Append(icon)
		return icon
	}
	return nil
}

func (f *FrowElement) LinkColor(title, href string, c color.Color) Element {
	if fi := f.GetFirstElementByTagName(compton_atoms.FlexItems); fi != nil {
		linkDecoration := Fspan(f.r, "").
			FontWeight(font_weight.Bolder).
			ForegroundColor(c)
		link := AText(title, href)
		link.SetAttribute("target", "_top")
		linkDecoration.Append(link)
		fi.Append(linkDecoration)
		return linkDecoration
	}
	return nil
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
		RowGap(size.Small).
		AlignItems(align.Center)
	fi.AddClass("frow")

	frow.Append(fi)

	return frow
}
