package table

import (
	_ "embed"
	"github.com/boggydigital/compton"
	"github.com/boggydigital/compton/elements"
	"golang.org/x/net/html/atom"
)

var (
	//go:embed "markup/table.html"
	markupTable []byte
)

type Table struct {
	compton.BaseElement
}

func (t *Table) AppendHead(columns ...string) *Table {

	var thead compton.Element
	if theads := t.GetElementsByTagName(atom.Thead); len(theads) > 0 {
		thead = theads[0]
	}

	if thead == nil {
		thead = NewHead()
		t.Append(thead)
	}

	for _, col := range columns {
		th := NewTh()
		th.Append(elements.NewText(col))
		thead.Append(th)
	}

	return t
}

func (t *Table) AppendRow(data ...string) *Table {

	var tbody compton.Element
	if tbodies := t.GetElementsByTagName(atom.Tbody); len(tbodies) > 0 {
		tbody = tbodies[0]
	}

	if tbody == nil {
		tbody = NewBody()
		t.Append(tbody)
	}

	tr := NewTr()
	for _, col := range data {
		td := NewTd()
		td.Append(elements.NewText(col))
		tr.Append(td)
	}
	tbody.Append(tr)

	return t
}

func (t *Table) AppendFoot(columns ...string) *Table {
	var tfoot compton.Element
	if tfeet := t.GetElementsByTagName(atom.Tfoot); len(tfeet) > 0 {
		tfoot = tfeet[0]
	}

	if tfoot == nil {
		tfoot = NewFoot()
		t.Append(tfoot)
	}

	for _, col := range columns {
		td := NewTd()
		td.Append(elements.NewText(col))
		tfoot.Append(td)
	}

	return t
}

func New() *Table {
	return &Table{
		compton.BaseElement{
			Markup:  markupTable,
			TagName: atom.Table,
		},
	}
}
