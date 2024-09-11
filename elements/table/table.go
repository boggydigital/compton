package table

import (
	_ "embed"
	"github.com/boggydigital/compton"
	"github.com/boggydigital/compton/elements/els"
	"golang.org/x/net/html/atom"
	"io"
)

const (
	registrationName      = "table"
	styleRegistrationName = "style" + registrationName
)

var (
	//go:embed "markup/table.html"
	markupTable []byte
	//go:embed "style/table.css"
	styleTable []byte
)

type TableElement struct {
	compton.BaseElement
	r compton.Registrar
}

func (te *TableElement) WriteStyles(w io.Writer) error {
	if te.r.RequiresRegistration(styleRegistrationName) {
		if _, err := w.Write(styleTable); err != nil {
			return err
		}
	}
	return te.BaseElement.WriteStyles(w)
}

func (te *TableElement) AppendHead(columns ...string) *TableElement {

	var thead compton.Element
	if theads := te.GetElementsByTagName(atom.Thead); len(theads) > 0 {
		thead = theads[0]
	}

	if thead == nil {
		thead = Thead()
		te.Append(thead)
	}

	for _, col := range columns {
		th := Th()
		th.Append(els.Text(col))
		thead.Append(th)
	}

	return te
}

func (te *TableElement) AppendRow(data ...string) *TableElement {

	var tbody compton.Element
	if tbodies := te.GetElementsByTagName(atom.Tbody); len(tbodies) > 0 {
		tbody = tbodies[0]
	}

	if tbody == nil {
		tbody = Tbody()
		te.Append(tbody)
	}

	tr := Tr()
	for _, col := range data {
		td := Td()
		td.Append(els.Text(col))
		tr.Append(td)
	}
	tbody.Append(tr)

	return te
}

func (te *TableElement) AppendFoot(columns ...string) *TableElement {
	var tfoot compton.Element
	if tfeet := te.GetElementsByTagName(atom.Tfoot); len(tfeet) > 0 {
		tfoot = tfeet[0]
	}

	if tfoot == nil {
		tfoot = Tfoot()
		te.Append(tfoot)
	}

	for _, col := range columns {
		td := Td()
		td.Append(els.Text(col))
		tfoot.Append(td)
	}

	return te
}

func Table(r compton.Registrar) *TableElement {
	return &TableElement{
		BaseElement: compton.BaseElement{
			Markup:  markupTable,
			TagName: atom.Table,
		},
		r: r,
	}
}
