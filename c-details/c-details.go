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
	elementName        = "c-details-"
	marginBlockEndAttr = "data-margin-block-end"
	openAttr           = "data-open"
)

var (
	//go:embed "markup/template.html"
	markupTemplate []byte
	//go:embed "markup/c-details.html"
	markupDetails []byte
)

type Details struct {
	compton.BaseElement
	open    bool
	summary string
	wcr     compton.Registrar
}

func openClosed(o bool) string {
	switch o {
	case true:
		return "open"
	case false:
		return "closed"
	}
	return ""
}

func (d *Details) Register(w io.Writer) error {
	openClosedName := elementName + openClosed(d.open)
	if d.wcr.RequiresRegistration(openClosedName) {
		if err := custom_elements.Define(w, custom_elements.Defaults(openClosedName)); err != nil {
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
	case ".OpenClosed":
		if _, err := io.WriteString(w, openClosed(d.open)); err != nil {
			return err
		}
	}
	return nil
}

func (d *Details) Open() *Details {
	d.open = true
	d.SetAttr(openAttr, compton.TrueVal)
	return d
}

func (d *Details) SetSummaryMarginBlockEnd(amount measures.Unit) *Details {
	d.SetAttr(marginBlockEndAttr, amount.String())
	return d
}

func (d *Details) Write(w io.Writer) error {
	return compton.WriteContents(bytes.NewReader(markupDetails), w, d.elementFragmentWriter)
}

func (d *Details) elementFragmentWriter(t string, w io.Writer) error {
	switch t {
	case ".Summary":
		if _, err := io.WriteString(w, d.summary); err != nil {
			return err
		}
	case ".OpenClosed":
		if _, err := io.WriteString(w, openClosed(d.open)); err != nil {
			return err
		}
	case compton.ContentToken:
		fallthrough
	case compton.AttributesToken:
		if err := d.BaseElement.WriteFragment(t, w); err != nil {
			return err
		}
	}
	return nil
}

func New(wcr compton.Registrar, summary string) *Details {
	return &Details{
		BaseElement: compton.BaseElement{
			Markup:  markupDetails,
			TagName: Atom,
		},
		wcr:     wcr,
		summary: summary,
	}
}
