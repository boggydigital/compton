package details_summary

import (
	_ "embed"
	"github.com/boggydigital/compton"
	"github.com/boggydigital/compton/consts/align"
	"github.com/boggydigital/compton/consts/class"
	"github.com/boggydigital/compton/consts/color"
	"github.com/boggydigital/compton/consts/compton_atoms"
	"github.com/boggydigital/compton/consts/direction"
	"github.com/boggydigital/compton/consts/size"
	"github.com/boggydigital/compton/elements/els"
	"github.com/boggydigital/compton/elements/flex_items"
	"github.com/boggydigital/compton/elements/svg_use"
	"golang.org/x/net/html/atom"
	"io"
)

const (
	registrationName      = "details-sum"
	styleRegistrationName = "style-" + registrationName
)

var (
	//go:embed "style/details-summary.css"
	styleDetailsSummary []byte
)

type DetailsSummaryElement struct {
	compton.BaseElement
	//r       compton.Registrar
	details compton.Element
}

func (dse *DetailsSummaryElement) Append(children ...compton.Element) {
	dse.details.Append(children...)
}

func (dse *DetailsSummaryElement) AppendSummary(children ...compton.Element) {
	if summary := dse.getSummary(); summary != nil {
		summary.Append(children...)
	}
}

//func (dse *DetailsSummaryElement) WriteStyles(w io.Writer) error {
//	if dse.r.RequiresRegistration(styleRegistrationName) {
//		if err := els.Style(styleDetailsSummary, styleRegistrationName).Write(w); err != nil {
//			return err
//		}
//	}
//	return dse.details.WriteStyles(w)
//}

//func (dse *DetailsSummaryElement) WriteRequirements(w io.Writer) error {
//	return dse.details.WriteRequirements(w)
//}
//
//func (dse *DetailsSummaryElement) WriteDeferrals(w io.Writer) error {
//	return dse.details.WriteDeferrals(w)
//}

func (dse *DetailsSummaryElement) SummaryMarginBlockEnd(s size.Size) *DetailsSummaryElement {
	if summaries := dse.details.GetElementsByTagName(atom.Summary); len(summaries) > 0 {
		summaries[0].AddClass(class.MarginBlockEnd(s))
	}
	return dse
}

func (dse *DetailsSummaryElement) DetailsMarginBlockEnd(s size.Size) *DetailsSummaryElement {
	dse.details.AddClass(class.MarginBlockEnd(s))
	return dse
}

func (dse *DetailsSummaryElement) SummaryRowGap(s size.Size) *DetailsSummaryElement {
	if summary := dse.getSummary(); summary != nil {
		summary.AddClass(class.RowGap(s))
	}
	return dse
}

func (dse *DetailsSummaryElement) BackgroundColor(c color.Color) *DetailsSummaryElement {
	if summary := dse.getSummary(); summary != nil {
		summary.AddClass(class.BackgroundColor(c))
	}
	return dse
}

func (dse *DetailsSummaryElement) ForegroundColor(c color.Color) *DetailsSummaryElement {
	if summary := dse.getSummary(); summary != nil {
		summary.AddClass(class.ForegroundColor(c))
	}
	return dse
}

func (dse *DetailsSummaryElement) MarkerColor(c color.Color) *DetailsSummaryElement {
	if summary := dse.getSummary(); summary != nil {
		summary.AddClass(class.MarkerColor(c))
	}
	return dse
}

func (dse *DetailsSummaryElement) getSummary() compton.Element {
	if summaries := dse.details.GetElementsByTagName(atom.Summary); len(summaries) > 0 {
		return summaries[0]
	}
	return nil
}

func (dse *DetailsSummaryElement) Write(w io.Writer) error {
	return dse.details.Write(w)
}

func (dse *DetailsSummaryElement) SetId(id string) {
	dse.details.SetId(id)
}

func create(r compton.Registrar, summary compton.Element, open bool) *DetailsSummaryElement {
	dse := &DetailsSummaryElement{
		BaseElement: compton.BaseElement{
			TagName: compton_atoms.DetailsSummary,
		},
		details: els.Details(),
		//r:       r,
	}

	if open {
		dse.details.SetAttribute("open", "")
	}

	svgPlus := svg_use.SvgUse(r, svg_use.Plus)
	svgPlus.AddClass("details-summary-marker")

	summaryElement := els.Summary()
	summaryTitleRow := flex_items.FlexItems(r, direction.Row).
		ColumnGap(size.Small).
		AlignItems(align.Center)
	summaryTitleRow.Append(svgPlus, summary)
	summaryElement.Append(summaryTitleRow)
	dse.details.Append(summaryElement)

	r.RegisterStyle(styleRegistrationName, styleDetailsSummary)

	return dse
}

func Larger(r compton.Registrar, summary compton.Element, open bool) *DetailsSummaryElement {
	dse := create(r, summary, open)
	dse.details.AddClass("larger")
	return dse
}

func Smaller(r compton.Registrar, summary compton.Element, open bool) *DetailsSummaryElement {
	dse := create(r, summary, open)
	dse.details.AddClass("smaller")
	return dse
}
