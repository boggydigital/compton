package compton

import (
	"github.com/boggydigital/compton/consts/align"
	"github.com/boggydigital/compton/consts/class"
	"github.com/boggydigital/compton/consts/color"
	"github.com/boggydigital/compton/consts/compton_atoms"
	"github.com/boggydigital/compton/consts/direction"
	"github.com/boggydigital/compton/consts/size"
	"golang.org/x/net/html/atom"
	"io"
	"strconv"
)

type DetailsSummaryElement struct {
	BaseElement
	details Element
}

func (dse *DetailsSummaryElement) Append(children ...Element) {
	dse.details.Append(children...)
}

func (dse *DetailsSummaryElement) getSummaryBadges() Element {
	if summary := dse.getSummary(); summary != nil {
		if labels := summary.GetElementsByClassName("badges"); len(labels) > 0 {
			return labels[0]
		}
	}
	return nil
}

func (dse *DetailsSummaryElement) AppendBadges(badges ...Element) {
	if summaryBadges := dse.getSummaryBadges(); summaryBadges != nil {
		summaryBadges.Append(badges...)
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

func (dse *DetailsSummaryElement) SummaryOutlineColor(c color.Color) *DetailsSummaryElement {
	if summary := dse.getSummary(); summary != nil {
		summary.AddClass(class.OutlineColor(c))
	}
	return dse
}

func (dse *DetailsSummaryElement) SummaryOutlineColorHex(hex string) *DetailsSummaryElement {
	if summary := dse.getSummary(); summary != nil {
		summary.AddClass(class.OutlineColorHex(hex))
	}
	return dse
}

func (dse *DetailsSummaryElement) SummaryAlignSelf(a align.Align) *DetailsSummaryElement {
	if summary := dse.getSummary(); summary != nil {
		summary.AddClass(class.AlignSelf(a))
	}
	return dse
}

func (dse *DetailsSummaryElement) SummaryJustifySelf(a align.Align) *DetailsSummaryElement {
	if summary := dse.getSummary(); summary != nil {
		summary.AddClass(class.JustifySelf(a))
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

func (dse *DetailsSummaryElement) SetTabIndex(index int) {
	dse.details.SetAttribute("tabindex", strconv.Itoa(index))
}

func create(r Registrar, title string, small, open bool) *DetailsSummaryElement {
	dse := &DetailsSummaryElement{
		BaseElement: BaseElement{
			TagName: compton_atoms.DetailsSummary,
		},
		details: Details(),
	}

	if open {
		dse.details.SetAttribute("open", "")
	}

	openMarker := SvgUse(r, Circle)
	openMarker.AddClass("open-marker")

	closedMarker := SvgUse(r, Circle)
	closedMarker.AddClass("closed-marker")

	summaryTitle := Fspan(r, title).
		FontSize(size.XSmall).
		Width(size.Unset).
		BackgroundColor(color.Transparent)

	summaryHeading := FlexItems(r, direction.Row).
		ColumnGap(size.Small).
		AlignItems(align.Center).
		ColumnWidthRule(size.Unset).
		Width(size.Unset)

	summaryHeading.AddClass("summary-heading")

	if small {
		summaryHeading.Append(summaryTitle, openMarker, closedMarker)
	} else {
		summaryHeading.Append(openMarker, closedMarker, summaryTitle)
	}

	summaryElement := Summary()
	summaryElement.Append(summaryHeading)

	if !small {
		summaryBadges := Content()
		summaryBadges.AddClass("badges")
		summaryElement.Append(summaryBadges)
	}

	dse.details.Append(summaryElement)

	r.RegisterStyles(DefaultStyle,
		compton_atoms.StyleName(compton_atoms.DetailsSummary))

	return dse
}

func DSLarge(r Registrar, title string, open bool) *DetailsSummaryElement {
	dse := create(r, title, false, open)
	dse.details.AddClass("larger")
	return dse
}

func DSSmall(r Registrar, title string, open bool) *DetailsSummaryElement {
	dse := create(r, title, true, open)
	dse.details.AddClass("smaller")
	return dse
}
