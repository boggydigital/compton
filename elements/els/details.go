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

type Details struct {
	compton.BaseElement
}

func (d *Details) AppendSummary(children ...compton.Element) *Details {
	var summary compton.Element
	if summaries := d.GetElementsByTagName(atom.Summary); len(summaries) > 0 {
		summary = summaries[0]
	}

	if summary == nil {
		summary = NewSummary()
		d.Append(summary)
	}

	for _, child := range children {
		summary.Append(child)
	}

	return d
}

func (d *Details) Open() *Details {
	d.SetAttr("open", "")
	return d
}

func NewDetails() *Details {
	return &Details{
		compton.BaseElement{
			Markup:  markupDetails,
			TagName: atom.Details,
		},
	}
}
