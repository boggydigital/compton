package textline

import (
	_ "embed"
	"github.com/boggydigital/compton"
	"github.com/boggydigital/compton/consts/align"
	"github.com/boggydigital/compton/consts/class"
	"github.com/boggydigital/compton/consts/color"
	"github.com/boggydigital/compton/consts/compton_atoms"
	"github.com/boggydigital/compton/consts/size"
	"github.com/boggydigital/compton/elements/els"
	"io"
)

const (
	registrationName      = "textline"
	styleRegistrationName = "style-" + registrationName
)

var (
	//go:embed "markup/textline.html"
	markupTextLine []byte
	//go:embed "style/textline.css"
	styleTextLine []byte
)

type TextLineElement struct {
	compton.BaseElement
	r compton.Registrar
}

func (tle *TextLineElement) WriteStyles(w io.Writer) error {
	if tle.r.RequiresRegistration(styleRegistrationName) {
		if err := els.Style(styleTextLine).WriteContent(w); err != nil {
			return err
		}
	}
	return tle.BaseElement.WriteStyles(w)
}

func (tle *TextLineElement) ForegroundColor(c color.Color) *TextLineElement {
	tle.AddClass(class.ForegroundColor(c))
	return tle
}

func (tle *TextLineElement) BackgroundColor(c color.Color) *TextLineElement {
	tle.AddClass(class.BackgroundColor(c))
	return tle
}

func (tle *TextLineElement) FontSize(s size.Size) *TextLineElement {
	tle.AddClass(class.FontSize(s))
	return tle
}

func (tle *TextLineElement) AlignContent(a align.Align) *TextLineElement {
	tle.AddClass(class.AlignContent(a))
	return tle
}

func (tle *TextLineElement) AlignItems(a align.Align) *TextLineElement {
	tle.AddClass(class.AlignItems(a))
	return tle
}

func (tle *TextLineElement) JustifyContent(a align.Align) *TextLineElement {
	tle.AddClass(class.JustifyContent(a))
	return tle
}

func (tle *TextLineElement) JustifyItems(a align.Align) *TextLineElement {
	tle.AddClass(class.JustifyItems(a))
	return tle
}

func Text(r compton.Registrar, t string) *TextLineElement {
	tle := &TextLineElement{
		BaseElement: compton.BaseElement{
			TagName: compton_atoms.TextLine,
			Markup:  markupTextLine,
		},
		r: r,
	}
	tle.Append(els.SpanText(t))
	return tle
}
