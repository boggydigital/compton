package details_toggle

import (
	"bytes"
	_ "embed"
	"github.com/boggydigital/compton"
	"github.com/boggydigital/compton/colors"
	"github.com/boggydigital/compton/compton_atoms"
	"github.com/boggydigital/compton/custom_elements"
	"github.com/boggydigital/compton/measures"
	"github.com/boggydigital/compton/shared"
	"io"
)

const (
	elementNameTemplate = "details-"
	marginAttr          = "data-margin"
	backgroundColorAttr = "data-background-color"
	foregroundColorAttr = "data-foreground-color"
	openAttr            = "data-open"
)

var (
	//go:embed "markup/template.html"
	markupTemplate []byte
	//go:embed "markup/details-toggle.html"
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
	openClosedName := elementNameTemplate + openClosed(d.open)
	if d.wcr.RequiresRegistration(openClosedName) {
		if err := custom_elements.Define(w, custom_elements.Defaults(openClosedName)); err != nil {
			return err
		}
		if err := compton.WriteContents(bytes.NewReader(markupTemplate), w, d.templateFragmentWriter); err != nil {
			return err
		}
	}
	return d.BaseElement.Register(w)
}

func (d *Details) templateFragmentWriter(t string, w io.Writer) error {
	switch t {
	case ".OpenClosed":
		if _, err := io.WriteString(w, openClosed(d.open)); err != nil {
			return err
		}
	case ".HostBackgroundColors":
		if _, err := io.Copy(w, bytes.NewReader(shared.StyleHostBackgroundColors)); err != nil {
			return err
		}
	case ".HostForegroundColors":
		if _, err := io.Copy(w, bytes.NewReader(shared.StyleHostForegroundColors)); err != nil {
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

func (d *Details) SetSummaryMargin(amount measures.Unit) *Details {
	d.SetAttr(marginAttr, amount.String())
	return d
}

func (d *Details) SetBackgroundColor(color colors.Color) *Details {
	d.SetAttr(backgroundColorAttr, color.String())
	return d
}

func (d *Details) SetForegroundColor(color colors.Color) *Details {
	d.SetAttr(foregroundColorAttr, color.String())
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
			TagName: compton_atoms.DetailsToggle,
		},
		wcr:     wcr,
		summary: summary,
	}
}
