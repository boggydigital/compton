package labels

import (
	_ "embed"
	"github.com/boggydigital/compton"
	"github.com/boggydigital/compton/consts/class"
	"github.com/boggydigital/compton/consts/compton_atoms"
	"github.com/boggydigital/compton/consts/size"
	"github.com/boggydigital/compton/elements/els"
	"golang.org/x/net/html/atom"
	"io"
)

const (
	registrationName      = "label"
	styleRegistrationName = "style-" + registrationName
)

var (
	//go:embed "style/labels.css"
	styleLabels []byte
)

type FormattedLabel struct {
	Property string
	Title    string
	Class    string
}

type LabelsElement struct {
	compton.BaseElement
	//r         compton.Registrar
	container compton.Element
}

func createLabelElement(fmtLabel FormattedLabel) compton.Element {
	label := els.ListItemText(fmtLabel.Title)
	cs := []string{"label", fmtLabel.Property, fmtLabel.Title, fmtLabel.Class}
	label.AddClass(cs...)
	return label
}

func (lse *LabelsElement) unorderedList() compton.Element {
	if uls := lse.container.GetElementsByTagName(atom.Ul); len(uls) > 0 {
		return uls[0]
	} else {
		panic("labels missing ul element")
	}
}

func (lse *LabelsElement) Write(w io.Writer) error {
	return lse.container.Write(w)
}

func (lse *LabelsElement) FontSize(s size.Size) *LabelsElement {
	lse.unorderedList().AddClass(class.FontSize(s))
	return lse
}

func (lse *LabelsElement) RowGap(s size.Size) *LabelsElement {
	lse.unorderedList().AddClass(class.RowGap(s))
	return lse
}

func (lse *LabelsElement) ColumnGap(s size.Size) *LabelsElement {
	lse.unorderedList().AddClass(class.ColumnGap(s))
	return lse
}

func Labels(r compton.Registrar, fmtLabels ...FormattedLabel) *LabelsElement {
	lse := &LabelsElement{
		BaseElement: compton.BaseElement{
			TagName: compton_atoms.Labels,
		},
		//r: r,
	}

	lse.container = els.Span()
	lse.container.AddClass("labels")

	ul := els.Ul()
	for _, fl := range fmtLabels {
		if fl.Title != "" {
			ul.Append(createLabelElement(fl))
		}
	}
	lse.container.Append(ul)

	r.RegisterStyle(styleRegistrationName, styleLabels)

	return lse
}
