package els

import (
	_ "embed"
	"github.com/boggydigital/compton"
	"github.com/boggydigital/compton/consts/attr"
	"golang.org/x/net/html/atom"
)

var (
	//go:embed "markup/datalist.html"
	markupDatalist []byte
)

func Datalist(id string) compton.Element {
	dataList := compton.NewElement(atom.Datalist, markupDatalist)
	dataList.SetAttribute(attr.Id, id)
	return dataList
}
