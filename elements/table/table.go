package table

import (
	_ "embed"
	"github.com/boggydigital/compton"
	"github.com/boggydigital/compton/elements/els"
	"golang.org/x/net/html/atom"
)

var (
	//go:embed "markup/table.html"
	markupTable []byte
)

type TableElement struct {
	compton.BaseElement
}

func (t *TableElement) AppendHead(columns ...string) *TableElement {

	var thead compton.Element
	if theads := t.GetElementsByTagName(atom.Thead); len(theads) > 0 {
		thead = theads[0]
	}

	if thead == nil {
		thead = Thead()
		t.Append(thead)
	}

	for _, col := range columns {
		th := Th()
		th.Append(els.Text(col))
		thead.Append(th)
	}

	return t
}

func (t *TableElement) AppendRow(data ...string) *TableElement {

	var tbody compton.Element
	if tbodies := t.GetElementsByTagName(atom.Tbody); len(tbodies) > 0 {
		tbody = tbodies[0]
	}

	if tbody == nil {
		tbody = Tbody()
		t.Append(tbody)
	}

	tr := Tr()
	for _, col := range data {
		td := Td()
		td.Append(els.Text(col))
		tr.Append(td)
	}
	tbody.Append(tr)

	return t
}

func (t *TableElement) AppendFoot(columns ...string) *TableElement {
	var tfoot compton.Element
	if tfeet := t.GetElementsByTagName(atom.Tfoot); len(tfeet) > 0 {
		tfoot = tfeet[0]
	}

	if tfoot == nil {
		tfoot = Tfoot()
		t.Append(tfoot)
	}

	for _, col := range columns {
		td := Td()
		td.Append(els.Text(col))
		tfoot.Append(td)
	}

	return t
}

func Table() *TableElement {
	return &TableElement{
		compton.BaseElement{
			Markup:  markupTable,
			TagName: atom.Table,
		},
	}
}
