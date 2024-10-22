package compton

import (
	_ "embed"
	"github.com/boggydigital/compton/consts/class"
	"github.com/boggydigital/compton/consts/compton_atoms"
	"github.com/boggydigital/compton/consts/size"
	"golang.org/x/net/html/atom"
	"io"
)

type FormattedLabel struct {
	Property string
	Title    string
	Class    string
}

type LabelsElement struct {
	BaseElement
	//r         compton.Registrar
	container Element
}

func createLabelElement(fmtLabel FormattedLabel) Element {
	label := ListItemText(fmtLabel.Title)
	cs := []string{"label", fmtLabel.Property, fmtLabel.Title, fmtLabel.Class}
	label.AddClass(cs...)
	return label
}

func (lse *LabelsElement) unorderedList() Element {
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

func Labels(r Registrar, fmtLabels ...FormattedLabel) *LabelsElement {
	lse := &LabelsElement{
		BaseElement: BaseElement{
			TagName: compton_atoms.Labels,
		},
		//r: r,
	}

	lse.container = Span()
	lse.container.AddClass("labels")

	ul := Ul()
	for _, fl := range fmtLabels {
		if fl.Title != "" {
			ul.Append(createLabelElement(fl))
		}
	}
	lse.container.Append(ul)

	r.RegisterStyle(compton_atoms.StyleName(compton_atoms.Labels), style)

	return lse
}
