package els

import (
	_ "embed"
	"github.com/boggydigital/compton"
	"golang.org/x/net/html/atom"
)

var (
	//go:embed "markup/datalist.html"
	markupDataList []byte
)

func NewDataList(id string) compton.Element {
	dataList := compton.NewElement(atom.Datalist, markupDataList)
	dataList.SetAttr(compton.IdAttr, id)
	return dataList
}
