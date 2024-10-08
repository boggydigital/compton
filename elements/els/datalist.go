package els

import (
	_ "embed"
	"github.com/boggydigital/compton"
	"golang.org/x/net/html/atom"
)

var (
	//go:embed "markup/datalist.html"
	markupDatalist []byte
)

func Datalist(id string) compton.Element {
	dataList := compton.NewElement(atom.Datalist, markupDatalist)
	dataList.SetAttribute(compton.IdAttr, id)
	return dataList
}
