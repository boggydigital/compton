package compton

import (
	"github.com/boggydigital/compton/consts/align"
	"github.com/boggydigital/compton/consts/color"
	"github.com/boggydigital/compton/consts/compton_atoms"
	"github.com/boggydigital/compton/consts/direction"
	"github.com/boggydigital/compton/consts/font_weight"
	"github.com/boggydigital/compton/consts/size"
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

func (f *FrowElement) PropVal(p, v string) *FrowElement {
	if fi := f.GetFirstElementByTagName(compton_atoms.FlexItems); fi != nil {
		fi.Append(Fspan(f.r, p+":").ForegroundColor(color.Gray))
		fi.Append(Fspan(f.r, v))
	}
	return f
}

func (f *FrowElement) PropIcons(p string, symbols ...Symbol) *FrowElement {
	if fi := f.GetFirstElementByTagName(compton_atoms.FlexItems); fi != nil {
		fi.Append(Fspan(f.r, p+":").ForegroundColor(color.Gray))
		for _, symbol := range symbols {
			fi.Append(SvgUse(f.r, symbol))
		}
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

func (f *FrowElement) CircleIcon(class string) *FrowElement {
	if fi := f.GetFirstElementByTagName(compton_atoms.FlexItems); fi != nil {
		icon := SvgUse(f.r, Circle)
		if class != "" {
			icon.AddClass(class)
		}
		fi.Append(icon)
	}
	return f
}

func (f *FrowElement) CircleIconColor(c color.Color) *FrowElement {
	if fi := f.GetFirstElementByTagName(compton_atoms.FlexItems); fi != nil {
		fi.Append(SvgUse(f.r, Circle).ForegroundColor(c))
	}
	return f
}

func (f *FrowElement) LinkExternal(title, href string) *FrowElement {
	if fi := f.GetFirstElementByTagName(compton_atoms.FlexItems); fi != nil {
		linkDecoration := Fspan(f.r, "").
			FontWeight(font_weight.Bolder).
			ForegroundColor(color.Cyan)
		link := AText(title, href)
		link.SetAttribute("target", "_top")
		linkDecoration.Append(link)
		fi.Append(linkDecoration)
	}
	return f
}

func Frow(r Registrar) *FrowElement {
	frow := &FrowElement{
		BaseElement: NewElement(compton_atoms.Frow, nil),
		r:           r,
	}

	fi := FlexItems(r, direction.Row).
		ColumnGap(size.XSmall).
		RowGap(size.Unset).
		AlignItems(align.Center).
		FontSize(size.Small)
	fi.AddClass("frow")

	frow.Append(fi)

	return frow
}
