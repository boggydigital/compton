package c_details

import (
	"bytes"
	_ "embed"
	"github.com/boggydigital/compton"
	"github.com/boggydigital/compton/custom_elements"
	"github.com/boggydigital/compton/measures"
	"golang.org/x/net/html/atom"
	"io"
)

const (
	// Atom for c-details is the second value created,
	// using max value + 1 and leaving 254 more possible atoms
	Atom atom.Atom = 0xffffff01
)

const (
	elementName = "c-details"
)

var (
	//go:embed "markup/template.html"
	markupTemplate []byte
	//go:embed "markup/c-details.html"
	markupDetails []byte
)

type Details struct {
	compton.BaseElement
	summary               string
	open                  bool
	summaryMarginBlockEnd measures.Unit
	wcr                   compton.Registrar
}

func (d *Details) Register(w io.Writer) error {
	if d.wcr.RequiresRegistration(elementName) {
		if err := custom_elements.Define(w, custom_elements.Defaults(elementName)); err != nil {
			return err
		}
		if err := compton.WriteContents(bytes.NewReader(markupTemplate), w, d.templateFragmentWriter); err != nil {
			return err
		}
	}
	return d.Parent.Register(w)
}

func (d *Details) templateFragmentWriter(t string, w io.Writer) error {
	switch t {
	case ".Open":
		if d.open {
			if _, err := io.WriteString(w, "open"); err != nil {
				return err
			}
		}
	case ".Summary":
		if _, err := io.WriteString(w, d.summary); err != nil {
			return err
		}
	case ".SummaryMarginBlockEnd":
		if _, err := io.WriteString(w, d.summaryMarginBlockEnd.String()); err != nil {
			return err
		}
	default:
		return compton.ErrUnknownToken(t)
	}
	return nil
}

//func (d *Details) Write(w io.Writer) error {
//	return d.BaseElement.Write(w)
//}

func (d *Details) Open() *Details {
	d.open = true
	return d
}

func (d *Details) SetSummaryMarginBlockEnd(amount measures.Unit) *Details {
	d.summaryMarginBlockEnd = amount
	return d
}

func New(wcr compton.Registrar, summary string) *Details {
	return &Details{
		BaseElement: compton.BaseElement{
			Markup:  markupDetails,
			TagName: Atom,
		},
		wcr:                   wcr,
		summary:               summary,
		summaryMarginBlockEnd: measures.Normal,
	}
}
