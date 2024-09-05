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

var Li = ListItem

func ListItem() compton.Element {
	return compton.NewElement(atom.Li, markupListItem)
}

func ListItemText(txt string) compton.Element {
	listItem := ListItem()
	listItem.Append(Text(txt))
	return listItem
}
