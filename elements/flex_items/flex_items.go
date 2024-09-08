package flex_items

import (
	_ "embed"
	"github.com/boggydigital/compton"
	"github.com/boggydigital/compton/consts/align"
	"github.com/boggydigital/compton/consts/class"
	"github.com/boggydigital/compton/consts/compton_atoms"
	"github.com/boggydigital/compton/consts/direction"
	"github.com/boggydigital/compton/consts/size"
	"github.com/boggydigital/compton/elements/els"
	"io"
)

const (
	registrationName      = "flex-items"
	styleRegistrationName = "style-" + registrationName
)

var (
	//go:embed "markup/flex-items.html"
	markupFlexItems []byte
	//go:embed "style/flex-items.css"
	styleFlexItems []byte
)

type FlexItemsElement struct {
	compton.BaseElement
	r compton.Registrar
}

func (fie *FlexItemsElement) WriteStyles(w io.Writer) error {
	if fie.r.RequiresRegistration(styleRegistrationName) {
		if err := els.Style(styleFlexItems).WriteContent(w); err != nil {
			return err
		}
	}
	return fie.BaseElement.WriteStyles(w)
}

//func (f *FlexItemsElement) WriteRequirements(w io.Writer) error {
//	if f.wcr.RequiresRegistration(registrationName) {
//		if err := custom_elements.Define(w, custom_elements.Defaults(registrationName)); err != nil {
//			return err
//		}
//		if err := compton.WriteContents(bytes.NewReader(markupTemplate), w, f.templateFragmentWriter); err != nil {
//			return err
//		}
//	}
//	return f.BaseElement.WriteRequirements(w)
//}

//func (f *FlexItemsElement) templateFragmentWriter(t string, w io.Writer) error {
//	switch t {
//	case ".Direction":
//		if _, err := io.WriteString(w, f.dir.String()); err != nil {
//			return err
//		}
//	case ".HostColumnGap":
//		if _, err := io.Copy(w, bytes.NewReader(shared.StyleHostColumnGap)); err != nil {
//			return err
//		}
//	case ".HostRowGap":
//		if _, err := io.Copy(w, bytes.NewReader(shared.StyleHostRowGap)); err != nil {
//			return err
//		}
//	case ".HostAlignContent":
//		if _, err := io.Copy(w, bytes.NewReader(shared.StyleHostAlignContent)); err != nil {
//			return err
//		}
//	case ".HostJustifyContent":
//		if _, err := io.Copy(w, bytes.NewReader(shared.StyleHostJustifyContent)); err != nil {
//			return err
//		}
//	case ".HostFlexDirection":
//		if _, err := io.Copy(w, bytes.NewReader(shared.StyleHostFlexDirection)); err != nil {
//			return err
//		}
//	}
//	return nil
//}

func (fie *FlexItemsElement) RowGap(sz size.Size) *FlexItemsElement {
	fie.AddClass(class.RowGap(sz))
	return fie
}

func (fie *FlexItemsElement) ColumnGap(sz size.Size) *FlexItemsElement {
	fie.AddClass(class.ColumnGap(sz))
	return fie
}

func (fie *FlexItemsElement) Gap(sz size.Size) *FlexItemsElement {
	fie.ColumnGap(sz)
	fie.RowGap(sz)
	return fie
}

func (fie *FlexItemsElement) AlignContent(a align.Align) *FlexItemsElement {
	fie.AddClass(class.AlignContent(a))
	return fie
}

func (fie *FlexItemsElement) AlignItems(a align.Align) *FlexItemsElement {
	fie.AddClass(class.AlignItems(a))
	return fie
}

func (fie *FlexItemsElement) JustifyContent(a align.Align) *FlexItemsElement {
	fie.AddClass(class.JustifyContent(a))
	return fie
}

func (fie *FlexItemsElement) JustifyItems(a align.Align) *FlexItemsElement {
	fie.AddClass(class.JustifyItems(a))
	return fie
}

func (fie *FlexItemsElement) FlexDirection(d direction.Direction) *FlexItemsElement {
	fie.AddClass(class.FlexDirection(d))
	return fie
}

func FlexItems(r compton.Registrar, dir direction.Direction) *FlexItemsElement {
	fie := &FlexItemsElement{
		BaseElement: compton.BaseElement{
			Markup:  markupFlexItems,
			TagName: compton_atoms.FlexItems,
		},
		r: r,
	}
	fie.AddClass(class.FlexDirection(dir))
	return fie
}
