package els

import (
	_ "embed"
	"github.com/boggydigital/compton"
	"golang.org/x/net/html/atom"
)

var (
	//go:embed "markup/li.html"
	markupListItem []byte
)

func NewListItem() compton.Element {
	return compton.NewElement(atom.Li, markupListItem)
}

func NewListItemText(txt string) compton.Element {
	listItem := NewListItem()
	listItem.Append(NewText(txt))
	return listItem
}
