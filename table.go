package compton

import (
	_ "embed"
	"github.com/boggydigital/compton/consts/compton_atoms"
	"golang.org/x/net/html/atom"
)

type TableElement struct {
	BaseElement
	r Registrar
}

func (te *TableElement) AppendHead(columns ...string) *TableElement {

	var thead Element
	if theads := te.GetElementsByTagName(atom.Thead); len(theads) > 0 {
		thead = theads[0]
	}

	if thead == nil {
		thead = Thead()
		te.Append(thead)
	}

	for _, col := range columns {
		th := Th()
		th.Append(Text(col))
		thead.Append(th)
	}

	return te
}

func (te *TableElement) AppendRow(data ...string) *TableElement {

	var tbody Element
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
		td.Append(Text(col))
		tr.Append(td)
	}
	tbody.Append(tr)

	return te
}

func (te *TableElement) AppendFoot(columns ...string) *TableElement {
	var tfoot Element
	if tfeet := te.GetElementsByTagName(atom.Tfoot); len(tfeet) > 0 {
		tfoot = tfeet[0]
	}

	if tfoot == nil {
		tfoot = Tfoot()
		te.Append(tfoot)
	}

	for _, col := range columns {
		td := Td()
		td.Append(Text(col))
		tfoot.Append(td)
	}

	return te
}

func Table(r Registrar) *TableElement {
	table := &TableElement{
		BaseElement: BaseElement{
			TagName:  atom.Table,
			Markup:   markup,
			Filename: atomMarkupFilename(atom.Table),
		},
		r: r,
	}

	r.RegisterStyle(compton_atoms.StyleName(atom.Table), style)

	return table
}
