package compton

import (
	_ "embed"
	"github.com/boggydigital/compton/consts/align"
	"github.com/boggydigital/compton/consts/class"
	"github.com/boggydigital/compton/consts/color"
	"github.com/boggydigital/compton/consts/compton_atoms"
	"github.com/boggydigital/compton/consts/direction"
	"github.com/boggydigital/compton/consts/font_weight"
	"github.com/boggydigital/compton/consts/size"
	"golang.org/x/net/html/atom"
	"io"
)

type DetailsSummaryElement struct {
	BaseElement
	details Element
}

func (dse *DetailsSummaryElement) Append(children ...Element) {
	dse.details.Append(children...)
}

func (dse *DetailsSummaryElement) AppendSummary(children ...Element) {
	if summary := dse.getSummary(); summary != nil {
		summary.Append(children...)
	}
}

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

func (dse *DetailsSummaryElement) getSummary() Element {
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

func create(r Registrar, summary Element, open bool) *DetailsSummaryElement {
	dse := &DetailsSummaryElement{
		BaseElement: BaseElement{
			TagName: compton_atoms.DetailsSummary,
		},
		details: Details(),
	}

	if open {
		dse.details.SetAttribute("open", "")
	}

	svgPlus := SvgUse(r, Plus)
	svgPlus.AddClass("details-summary-marker")

	summaryElement := Summary()
	summaryTitleRow := FlexItems(r, direction.Row).
		ColumnGap(size.Small).
		AlignItems(align.Center)
	summaryTitleRow.Append(svgPlus, summary)
	summaryElement.Append(summaryTitleRow)
	dse.details.Append(summaryElement)

	r.RegisterStyles(DefaultStyle,
		compton_atoms.StyleName(compton_atoms.DetailsSummary))

	return dse
}

func DSLarge(r Registrar, summary Element, open bool) *DetailsSummaryElement {
	dse := create(r, summary, open)
	dse.details.AddClass("larger")
	return dse
}

func DSSmall(r Registrar, summary Element, open bool) *DetailsSummaryElement {
	dse := create(r, summary, open)
	dse.details.AddClass("smaller")
	return dse
}

func DSTitle(r Registrar, title string) Element {
	fs := Fspan(r, title).
		FontWeight(font_weight.Bolder).
		FontSize(size.Large)
	return fs
}
