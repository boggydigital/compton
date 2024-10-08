package els

import (
	_ "embed"
	"github.com/boggydigital/compton"
	"golang.org/x/net/html/atom"
)

var (
	//go:embed "markup/details.html"
	markupDetails []byte
)

type DetailsElement struct {
	compton.BaseElement
}

func (d *DetailsElement) AppendSummary(children ...compton.Element) *DetailsElement {
	var summary compton.Element
	if summaries := d.GetElementsByTagName(atom.Summary); len(summaries) > 0 {
		summary = summaries[0]
	}

	if summary == nil {
		summary = Summary()
		d.Append(summary)
	}

	for _, child := range children {
		summary.Append(child)
	}

	return d
}

func (d *DetailsElement) Open() *DetailsElement {
	d.SetAttribute("open", "")
	return d
}

func Details() *DetailsElement {
	return &DetailsElement{
		compton.BaseElement{
			Markup:  markupDetails,
			TagName: atom.Details,
		},
	}
}
